package paclient

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/phrase/phraseapp-client/spinner"
	"github.com/phrase/phraseapp-go/phraseapp"
)

type Sources []*Source

type Source struct {
	File        string
	ProjectID   string
	AccessToken string
	FileFormat  string
	Params      *phraseapp.UploadParams

	RemoteLocales []*phraseapp.Locale
	Format        *phraseapp.Format
}

func (sources Sources) Validate() error {
	for _, source := range sources {
		if err := source.CheckPreconditions(); err != nil {
			return err
		}
	}
	return nil
}

func (source *Source) GetLocaleID() string {
	if source.Params != nil && source.Params.LocaleID != nil {
		return *source.Params.LocaleID
	}
	return ""
}

func (source *Source) GetFileFormat() string {
	if source.Params != nil && source.Params.FileFormat != nil {
		return *source.Params.FileFormat
	}
	if source.FileFormat != "" {
		return source.FileFormat
	}
	return ""
}

func (source *Source) CheckPreconditions() error {
	if err := ValidPath(source.File, source.FileFormat, ""); err != nil {
		return err
	}

	duplicatedPlaceholders := []string{}
	for _, name := range []string{"<locale_name>", "<locale_code>", "<tag>"} {
		if strings.Count(source.File, name) > 1 {
			duplicatedPlaceholders = append(duplicatedPlaceholders, name)
		}
	}

	starCount := strings.Count(source.File, "*")
	recCount := strings.Count(source.File, "**")

	// starCount contains the `**` so that must be taken into account.
	if starCount-(recCount*2) > 1 {
		duplicatedPlaceholders = append(duplicatedPlaceholders, "*")
	}

	if recCount > 1 {
		duplicatedPlaceholders = append(duplicatedPlaceholders, "**")
	}

	if len(duplicatedPlaceholders) > 0 {
		dups := strings.Join(duplicatedPlaceholders, ", ")
		return fmt.Errorf(fmt.Sprintf("%s can only occur once in a file pattern!", dups))
	}

	return nil
}

func (src *Source) UnmarshalYAML(unmarshal func(interface{}) error) error {
	m := map[string]interface{}{}
	err := phraseapp.ParseYAMLToMap(unmarshal, map[string]interface{}{
		"file":         &src.File,
		"project_id":   &src.ProjectID,
		"access_token": &src.AccessToken,
		"file_format":  &src.FileFormat,
		"params":       &m,
	})
	if err != nil {
		return err
	}

	src.Params = new(phraseapp.UploadParams)
	return src.Params.ApplyValuesFromMap(m)
}

func (sources Sources) ProjectIds() []string {
	projectIds := []string{}
	for _, source := range sources {
		projectIds = append(projectIds, source.ProjectID)
	}
	return projectIds
}
func (source *Source) uploadFile(client *phraseapp.Client, localeFile *LocaleFile) (*phraseapp.Upload, error) {
	if Debug {
		fmt.Fprintln(os.Stdout, "Source file pattern:", source.File)
		fmt.Fprintln(os.Stdout, "Actual file location:", localeFile.Path)
	}

	params := new(phraseapp.UploadParams)
	*params = *source.Params

	params.File = &localeFile.Path

	if params.LocaleID == nil {
		switch {
		case localeFile.ID != "":
			params.LocaleID = &localeFile.ID
		case localeFile.Code != "":
			params.LocaleID = &localeFile.Code
		}
	}

	if localeFile.Tag != "" {
		var v string
		if params.Tags != nil {
			v = *params.Tags + ","
		}
		v += localeFile.Tag
		params.Tags = &v
	}

	return client.UploadCreate(source.ProjectID, params)
}

func (source *Source) createLocale(client *phraseapp.Client, localeFile *LocaleFile) (*phraseapp.LocaleDetails, error) {
	localeDetails, found, err := source.getLocaleIfExist(client, localeFile)
	if err != nil {
		return nil, err
	} else if found {
		return localeDetails, nil
	}

	localeParams := new(phraseapp.LocaleParams)

	if localeFile.Name != "" {
		localeParams.Name = &localeFile.Name
	} else if localeFile.Code != "" {
		localeParams.Name = &localeFile.Code
	}

	if localeFile.Code == "" {
		localeFile.Code = localeFile.Name
	}

	localeName := source.replacePlaceholderInParams(localeFile)
	if localeName != "" && localeName != localeFile.Code {
		localeParams.Name = &localeName
	}

	if localeFile.Code != "" {
		localeParams.Code = &localeFile.Code
	}

	localeDetails, err = client.LocaleCreate(source.ProjectID, localeParams)
	if err != nil {
		return nil, err
	}
	return localeDetails, nil
}

func (source *Source) getLocaleIfExist(client *phraseapp.Client, localeFile *LocaleFile) (*phraseapp.LocaleDetails, bool, error) {
	identifier := localeIdentifier(source, localeFile)
	if identifier == "" {
		return nil, false, nil
	}

	localeDetail, err := client.LocaleShow(source.ProjectID, identifier)
	if isNotFound(err) {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	return localeDetail, true, nil
}

func (source *Source) replacePlaceholderInParams(localeFile *LocaleFile) string {
	if localeFile.Code != "" && strings.Contains(source.GetLocaleID(), "<locale_code>") {
		return strings.Replace(source.GetLocaleID(), "<locale_code>", localeFile.Code, 1)
	}
	return ""
}

func localeIdentifier(source *Source, localeFile *LocaleFile) string {
	localeName := source.replacePlaceholderInParams(localeFile)
	if localeName != "" && localeName != localeFile.Code {
		return localeName
	}

	if localeFile.Name != "" {
		return localeFile.Name
	}

	if localeFile.Code != "" {
		return localeFile.Code
	}

	return ""
}

func isNotFound(err error) bool {
	return (err != nil && strings.Contains(err.Error(), "404"))
}

// Return all locale files from disk that match the source pattern.
func (source *Source) LocaleFiles() (LocaleFiles, error) {
	filePaths, err := source.SystemFiles()
	if err != nil {
		return nil, err
	}

	patternTokens := splitPathToTokens(source.File)

	var localeFiles LocaleFiles
	for _, path := range filePaths {
		pathTokens := splitPathToTokens(path)
		localeFile := extractParamsFromPathTokens(patternTokens, pathTokens)

		absolutePath, err := filepath.Abs(path)
		if err != nil {
			return nil, err
		}

		localeFile.Path = absolutePath

		locale := source.getRemoteLocaleForLocaleFile(localeFile)
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

	if len(localeFiles) <= 0 {
		abs, err := filepath.Abs(source.File)
		if err != nil {
			abs = source.File
		}
		return nil, fmt.Errorf("Could not find any files on your system that matches: '%s'", abs)
	}
	return localeFiles, nil
}

func (source *Source) Push(client *phraseapp.Client) error {
	localeFiles, err := source.LocaleFiles()
	if err != nil {
		return err
	}

	for _, localeFile := range localeFiles {
		fmt.Println("Uploading", localeFile.RelPath())

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

		taskResult := make(chan string, 1)
		taskErr := make(chan error, 1)

		spinner.Spin("Waiting for your file to be processed... ", func(taskFinished chan<- struct{}) {
			result, err := getUploadResult(client, source.ProjectID, upload)
			taskResult <- result
			taskErr <- err
			taskFinished <- struct{}{}
		})

		if err := <-taskErr; err != nil {
			return err
		}

		switch <-taskResult {
		case "success":
			PrintSuccess("Uploaded " + localeFile.RelPath() + " successfully.")
		case "error":
			PrintFailure("There was an error processing " + localeFile.RelPath() + ". Your changes were not saved online.")
		}

		if Debug {
			fmt.Fprintln(os.Stderr, strings.Repeat("-", 10))
		}
	}
	return nil
}

func (source *Source) SystemFiles() ([]string, error) {
	pattern := PlaceholderRegexp.ReplaceAllString(source.File, "*")
	parts := strings.SplitN(pattern, "**", 2)
	var pre, post string

	pre = parts[0]
	// strip trailing path separators
	for len(pre) > 0 && os.IsPathSeparator(pre[len(pre)-1]) {
		pre = pre[0 : len(pre)-1]
	}

	if len(parts) == 2 {
		post = parts[1]
		for len(post) > 0 && os.IsPathSeparator(post[0]) {
			post = post[1:]
		}
	}

	candidates, err := filepath.Glob(pre)
	if err != nil {
		return nil, err
	}

	prefix := "." + string(os.PathSeparator)
	if strings.HasPrefix(pre, prefix) {
		pre = strings.TrimLeft(pre, prefix)
	}

	var matches []string
	if post != "" {
		tokens := splitPathIntoSegments(strings.Replace(post, "*", ".*", -1))
		tokenCountPre := len(splitPathIntoSegments(pre))

		for _, cand := range candidates {
			if !isDir(cand) {
				continue
			}

			cands, err := findFilesInPath(cand)
			if err != nil {
				return nil, err
			}

			for _, c := range cands {
				if validateFileCandidate(tokens, tokenCountPre, c) {
					matches = append(matches, c)
				}
			}
		}
	} else {
		matches = candidates
	}

	return matches, nil
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

func splitPathToTokens(s string) []string {
	tokens := []string{}
	splitSet := separator
	if separator == "\\" {
		splitSet = "\\/"
	}
	for _, token := range splitString(s, splitSet) {
		if token == "." || token == "" {
			continue
		}
		tokens = append(tokens, token)
	}
	return tokens
}

func extractParamsFromPathTokens(patternTokens, pathTokens []string) *LocaleFile {
	localeFile := new(LocaleFile)

	if Debug {
		fmt.Println("pattern:", patternTokens)
		fmt.Println("path:", pathTokens)
	}

	for idx, patternToken := range patternTokens {
		pathToken := pathTokens[idx]

		if patternToken == "*" {
			continue
		}
		if patternToken == "**" {
			break
		}
		localeFile.extractParamFromPathToken(patternToken, pathToken)
	}

	if Contains(patternTokens, "**") {
		offset := 1
		for idx := len(patternTokens) - 1; idx >= 0; idx-- {
			patternToken := patternTokens[idx]
			pathToken := pathTokens[len(pathTokens)-offset]
			offset += 1

			if patternToken == "*" {
				continue
			} else if patternToken == "**" {
				break
			}

			localeFile.extractParamFromPathToken(patternToken, pathToken)
		}
	}

	return localeFile
}
