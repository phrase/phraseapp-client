package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jpillora/backoff"
	"github.com/phrase/phraseapp-client/internal/paths"
	"github.com/phrase/phraseapp-client/internal/placeholders"
	"github.com/phrase/phraseapp-client/internal/print"
	"github.com/phrase/phraseapp-client/internal/spinner"
	"github.com/phrase/phraseapp-go/phraseapp"
)

type PushCommand struct {
	phraseapp.Config
	Wait bool `cli:"opt --wait desc='Wait for files to be processed'"`
}

func (cmd *PushCommand) Run() error {
	if cmd.Config.Debug {
		// suppresses content output
		cmd.Config.Debug = false
		Debug = true
	}

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	sources, err := SourcesFromConfig(cmd.Config)
	if err != nil {
		return err
	}

	if err := sources.Validate(); err != nil {
		return err
	}

	formatMap, err := formatsByApiName(client)
	if err != nil {
		return fmt.Errorf("Error retrieving format list from PhraseApp: %s", err)
	}

	for _, source := range sources {
		formatName := source.GetFileFormat()
		if val, ok := formatMap[formatName]; ok {
			source.Format = val
		}

		if source.Format == nil {
			return fmt.Errorf("Format %q of source %q is not supported by PhraseApp!", formatName, source.File)
		}
	}

	projectIdToLocales, err := LocalesForProjects(client, sources)
	if err != nil {
		return err
	}
	for _, source := range sources {
		val, ok := projectIdToLocales[source.ProjectID]
		if ok {
			source.RemoteLocales = val
		}
	}

	for _, source := range sources {
		err := source.Push(client, cmd.Wait)
		if err != nil {
			return err
		}
	}
	return nil
}

func (source *Source) Push(client *phraseapp.Client, waitForResults bool) error {
	localeFiles, err := source.LocaleFiles()
	if err != nil {
		return err
	}

	for _, localeFile := range localeFiles {
		fmt.Printf("Uploading %s... ", localeFile.RelPath())

		if localeFile.shouldCreateLocale(source) {
			localeDetails, err := source.createLocale(client, localeFile)
			if err == nil {
				localeFile.ID = localeDetails.ID
				localeFile.Code = localeDetails.Code
				localeFile.Name = localeDetails.Name
			} else {
				fmt.Printf("failed to create locale: %s\n", err)
				continue
			}
		}

		upload, err := source.uploadFile(client, localeFile)
		if err != nil {
			return err
		}

		if waitForResults {
			fmt.Println()

			taskResult := make(chan string, 1)
			taskErr := make(chan error, 1)

			fmt.Printf("Upload ID: %s, filename: %s suceeded. Waiting for your file to be processed... ", upload.ID, upload.Filename)
			spinner.While(func() {
				result, err := getUploadResult(client, source.ProjectID, upload)
				taskResult <- result
				taskErr <- err
			})
			fmt.Println()

			if err := <-taskErr; err != nil {
				return err
			}

			switch <-taskResult {
			case "success":
				print.Success("Successfully uploaded and processed %s.", localeFile.RelPath())
			case "error":
				print.Failure("There was an error processing %s. Your changes were not saved online.", localeFile.RelPath())
			}
		} else {
			fmt.Println("done!")
			fmt.Printf("Check upload ID: %s, filename: %s for information about processing results.\n", upload.ID, upload.Filename)
		}

		if Debug {
			fmt.Fprintln(os.Stderr, strings.Repeat("-", 10))
		}
	}

	return nil
}

func formatsByApiName(client *phraseapp.Client) (map[string]*phraseapp.Format, error) {
	formats, err := client.FormatsList(1, 25)
	if err != nil {
		return nil, err
	}
	formatMap := map[string]*phraseapp.Format{}
	for _, format := range formats {
		formatMap[format.ApiName] = format
	}
	return formatMap, nil
}

// Return all locale files from disk that match the source pattern.
func (source *Source) LocaleFiles() (LocaleFiles, error) {
	filePaths, err := paths.Glob(placeholders.ToGlobbingPattern(source.File))
	if err != nil {
		return nil, err
	}

	var localeFiles LocaleFiles
	for _, path := range filePaths {
		if paths.IsPhraseAppYmlConfig(path) {
			continue
		}

		localeFile := new(LocaleFile)
		localeFile.fillFromPath(path, source.File)

		localeFile.Path, err = filepath.Abs(path)
		if err != nil {
			return nil, err
		}

		locale := source.getRemoteLocaleForLocaleFile(localeFile)
		// TODO: sinnvoll?
		if locale != nil {
			localeFile.ExistsRemote = true
			localeFile.Code = locale.Code
			localeFile.Name = locale.Name
			localeFile.ID = locale.ID
		}

		if Debug {
			fmt.Printf(
				"Code:%q, Name:%q, ID:%q, Tag:%q\n",
				localeFile.Code, localeFile.Name, localeFile.ID, localeFile.Tag,
			)
		}

		localeFiles = append(localeFiles, localeFile)
	}

	if len(localeFiles) == 0 {
		abs, err := filepath.Abs(source.File)
		if err != nil {
			abs = source.File
		}
		return nil, fmt.Errorf("Could not find any files on your system that matches: '%s'", abs)
	}

	return localeFiles, nil
}

func (source *Source) getRemoteLocaleForLocaleFile(localeFile *LocaleFile) *phraseapp.Locale {
	candidates := source.RemoteLocales

	filterApplied := false

	filter := func(cands []*phraseapp.Locale, preCond string, pred func(cand *phraseapp.Locale) bool) []*phraseapp.Locale {
		if preCond == "" {
			return cands
		}
		filterApplied = true
		tmpCands := []*phraseapp.Locale{}
		for _, cand := range cands {
			if pred(cand) {
				tmpCands = append(tmpCands, cand)
			}
		}
		return tmpCands
	}

	localeName := source.replacePlaceholderInParams(localeFile)
	if localeName != "" {
		// This means the name can contain the value specified in LocaleID, with
		// `<locale_code>` being substituted by the value of the currently handled
		// localeFile (like push only locales with name `en-US`).
		candidates = filter(candidates, localeName, func(cand *phraseapp.Locale) bool {
			return strings.Contains(cand.Name, localeName)
		})
	} else {
		localeID := source.GetLocaleID()
		candidates = filter(candidates, localeID, func(cand *phraseapp.Locale) bool {
			return cand.Name == localeID || cand.ID == localeID
		})
	}

	candidates = filter(candidates, localeFile.Name, func(cand *phraseapp.Locale) bool {
		return cand.Name == localeFile.Name
	})

	candidates = filter(candidates, localeFile.Code, func(cand *phraseapp.Locale) bool {
		return cand.Code == localeFile.Code
	})

	// If no filter was applied the candidates list still contains all remote
	// locales, while actually nothing matches.
	if !filterApplied {
		return nil
	}

	switch len(candidates) {
	case 0:
		return nil
	case 1:
		return candidates[0]
	default:
		// TODO I guess this should return an error, as this is a problem.
		return candidates[0]
	}
}

func (localeFile *LocaleFile) fillFromPath(path, pattern string) {
	path = filepath.ToSlash(path)
	pathStart, patternStart, pathEnd, patternEnd, err := paths.SplitAtDirGlobOperator(path, pattern)
	if err != nil {
		print.Error(err)
		return
	}

	fillFrom := func(path, pattern string) {
		params, err := placeholders.Resolve(path, pattern)
		if err != nil {
			print.Error(err)
			return
		}

		for placeholder, value := range params {
			switch placeholder {
			case "locale_code":
				localeFile.Code = value
			case "locale_name":
				localeFile.Name = value
			case "tag":
				localeFile.Tag = value
			}
		}
	}

	fillFrom(pathStart, patternStart)
	fillFrom(pathEnd, patternEnd)
}

func (localeFile *LocaleFile) shouldCreateLocale(source *Source) bool {
	if localeFile.ExistsRemote {
		return false
	}

	if source.Format.IncludesLocaleInformation {
		return false
	}

	// we could not find an existing locale in PhraseApp
	// if a locale_name or locale_code was provided by the placeholder logic
	// we assume that it should be created
	// every other source should be uploaded and validated in uploads#create
	return (localeFile.Name != "" || localeFile.Code != "")
}

func getUploadResult(client *phraseapp.Client, projectID string, upload *phraseapp.Upload) (result string, err error) {
	b := &backoff.Backoff{
		Min:    500 * time.Millisecond,
		Max:    10 * time.Second,
		Factor: 2,
		Jitter: true,
	}

	for ; result != "success" && result != "error"; result = upload.State {
		time.Sleep(b.Duration())
		upload, err = client.UploadShow(projectID, upload.ID)
		if err != nil {
			break
		}
	}

	return
}
