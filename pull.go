package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/phrase/phraseapp-client/internal/print"
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
		if !ok || len(val) == 0 {
			return fmt.Errorf("Could not find any locales for project %q", target.ProjectID)
		}
		target.RemoteLocales = val
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

	for _, localeFile := range localeFiles {
		err := createFile(localeFile.Path)
		if err != nil {
			return err
		}

		err = target.DownloadAndWriteToFile(client, localeFile)
		if err != nil {
			return fmt.Errorf("%s for %s", err, localeFile.Path)
		} else {
			print.Success("Downloaded %s to %s", localeFile.Message(), localeFile.RelPath())
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
	files := []*LocaleFile{}

	if target.GetLocaleID() != "" {
		// a specific locale was requested
		remoteLocale, err := target.localeForRemote()
		if err != nil {
			return nil, err
		}

		localeFile, err := createLocaleFile(target, remoteLocale)
		if err != nil {
			return nil, err
		}

		files = append(files, localeFile)

	} else if containsLocalePlaceholder(target.File) {
		// multiple locales were requested
		for _, remoteLocale := range target.RemoteLocales {
			localeFile, err := createLocaleFile(target, remoteLocale)
			if err != nil {
				return nil, err
			}

			files = append(files, localeFile)
		}
	} else {
		// no local files match remote locale
		return nil, fmt.Errorf("Could not find any files on your system that matches the locales for porject %q.", target.ProjectID)
	}

	return files, nil
}

func createLocaleFile(target *Target, remoteLocale *phraseapp.Locale) (*LocaleFile, error) {
	localeFile := &LocaleFile{
		Name:       remoteLocale.Name,
		ID:         remoteLocale.ID,
		Code:       remoteLocale.Code,
		Tag:        target.GetTag(),
		FileFormat: target.GetFormat(),
		Path:       target.File,
	}

	absPath, err := resolvedPath(localeFile)
	if err != nil {
		return nil, err
	}

	localeFile.Path = absPath
	return localeFile, nil
}

func resolvedPath(localeFile *LocaleFile) (string, error) {
	absPath, err := filepath.Abs(localeFile.Path)
	if err != nil {
		return "", err
	}

	path := strings.Replace(absPath, "<locale_name>", localeFile.Name, -1)
	path = strings.Replace(path, "<locale_code>", localeFile.Code, -1)
	path = strings.Replace(path, "<tag>", localeFile.Tag, -1)

	return path, nil
}
