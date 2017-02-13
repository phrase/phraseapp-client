package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	ct "github.com/daviddengcn/go-colortext"
	"github.com/phrase/phraseapp-go/phraseapp"
)

type PullCommand struct {
	phraseapp.Config
}

func (cmd *PullCommand) Run() error {
	if cmd.Config.Debug {
		// suppresses content output
		cmd.Config.Debug = false
		Debug = true
	}
	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	targets, err := TargetsFromConfig(cmd.Config)
	if err != nil {
		return err
	}

	projectIdToLocales, err := LocalesForProjects(client, targets)
	if err != nil {
		return err
	}
	for _, target := range targets {
		val, ok := projectIdToLocales[target.ProjectID]
		if ok {
			target.RemoteLocales = val
		}
	}

	for _, target := range targets {
		err := target.Pull(client)
		if err != nil {
			return err
		}
	}

	return nil
}

type PullParams struct {
	phraseapp.LocaleDownloadParams
	LocaleID string
}

func (target *Target) Pull(client *phraseapp.Client) error {
	if err := target.CheckPreconditions(); err != nil {
		return err
	}

	localeFiles, err := target.LocaleFiles()
	if err != nil {
		return err
	}

	localeIdToFileIsDistinct := (target.GetLocaleID() != "" && len(localeFiles) == 1)

	for _, localeFile := range localeFiles {
		err := createFile(localeFile.Path)
		if err != nil {
			return err
		}

		if localeIdToFileIsDistinct {
			if target.GetLocaleID() != "" {
				localeFile.ID = target.GetLocaleID()
			}
		}

		err = target.DownloadAndWriteToFile(client, localeFile)
		if err != nil {
			return fmt.Errorf("%s for %s", err, localeFile.Path)
		} else {
			pullMessage(localeFile)
		}
		if Debug {
			fmt.Fprintln(os.Stderr, strings.Repeat("-", 10))
		}
	}

	return nil
}

func (target *Target) DownloadAndWriteToFile(client *phraseapp.Client, localeFile *LocaleFile) error {
	downloadParams := new(phraseapp.LocaleDownloadParams)
	if target.Params != nil {
		*downloadParams = target.Params.LocaleDownloadParams
	}

	if downloadParams.FileFormat == nil {
		downloadParams.FileFormat = &localeFile.FileFormat
	}

	if Debug {
		fmt.Fprintln(os.Stderr, "Target file pattern:", target.File)
		fmt.Fprintln(os.Stderr, "Actual file path", localeFile.Path)
		fmt.Fprintln(os.Stderr, "LocaleID", localeFile.ID)
		fmt.Fprintln(os.Stderr, "ProjectID", target.ProjectID)
		fmt.Fprintln(os.Stderr, "FileFormat", downloadParams.FileFormat)
		fmt.Fprintln(os.Stderr, "ConvertEmoji", downloadParams.ConvertEmoji)
		fmt.Fprintln(os.Stderr, "IncludeEmptyTranslations", downloadParams.IncludeEmptyTranslations)
		fmt.Fprintln(os.Stderr, "KeepNotranslateTags", downloadParams.KeepNotranslateTags)
		fmt.Fprintln(os.Stderr, "Tag", downloadParams.Tag)
		fmt.Fprintln(os.Stderr, "FormatOptions", downloadParams.FormatOptions)
	}

	res, err := client.LocaleDownload(target.ProjectID, localeFile.ID, downloadParams)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(localeFile.Path, res, 0700)
	if err != nil {
		return err
	}
	return nil
}

func (target *Target) LocaleFiles() (LocaleFiles, error) {
	localeID := target.GetLocaleID()

	files := []*LocaleFile{}
	for _, remoteLocale := range target.RemoteLocales {
		if localeID != "" && !(remoteLocale.ID == localeID || remoteLocale.Name == localeID) {
			continue
		}
		err := target.IsValidLocale(remoteLocale, target.File)
		if err != nil {
			return nil, err
		}

		localeFile := &LocaleFile{
			Name:       remoteLocale.Name,
			ID:         remoteLocale.ID,
			Code:       remoteLocale.Code,
			Tag:        target.GetTag(),
			FileFormat: target.GetFormat(),
			Path:       target.File,
		}

		absPath, err := target.ReplacePlaceholders(localeFile)
		if err != nil {
			return nil, err
		}
		localeFile.Path = absPath

		files = append(files, localeFile)
	}

	return files, nil
}

func (target *Target) IsValidLocale(locale *phraseapp.Locale, localPath string) error {
	if locale == nil {
		return fmt.Errorf("Remote locale could not be downloaded correctly!")
	}

	if strings.Contains(localPath, "<locale_code>") && locale.Code == "" {
		return fmt.Errorf("Locale code is not set for Locale with ID: %s but locale_code is used in file name", locale.ID)
	}
	return nil
}

func (target *Target) ReplacePlaceholders(localeFile *LocaleFile) (string, error) {
	absPath, err := filepath.Abs(target.File)
	if err != nil {
		return "", err
	}

	path := strings.Replace(absPath, "<locale_name>", localeFile.Name, -1)
	path = strings.Replace(path, "<locale_code>", localeFile.Code, -1)
	path = strings.Replace(path, "<tag>", localeFile.Tag, -1)

	return path, nil
}

func pullMessage(localeFile *LocaleFile) {
	local := localeFile.RelPath()
	remote := localeFile.Message()
	fmt.Print("Downloaded ")
	ct.Foreground(ct.Green, true)
	fmt.Print(remote)
	ct.ResetColor()
	fmt.Print(" to ")
	ct.Foreground(ct.Green, true)
	fmt.Print(local, "\n")
	ct.ResetColor()
}
