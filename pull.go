package main

import (
	"path/filepath"

	"fmt"
	"github.com/phrase/phraseapp-api-client/Godeps/_workspace/src/gopkg.in/yaml.v2"
	"github.com/phrase/phraseapp-go/phraseapp"
	"os"
	"path"
)

type Target *PullArgs
type Targets []Target

type PullArgs struct {
	File        string      `yaml:"file,omitempty"`
	ProjectId   string      `yaml:"project_id,omitempty"`
	AccessToken string      `yaml:"access_token,omitempty"`
	Params      *PullParams `yaml:"params,omitempty"`
}

type PullParams struct {
	FileFormat string `yaml:"file_format,omitempty"`
	LocaleId   string `yaml:"locale_id,omitempty"`
}

func Pull(p *PhrasePath, target Target) error {
	authenticate()

	locales, err := phraseapp.LocalesList(target.ProjectId, 1, 25)
	if err != nil {
		return err
	}

	paths, err := expandPathsWithLocale(p, target, locales)
	if err != nil {
		return err
	}

	virtualPaths, err := fileGlobbingPull(p, paths)
	if err != nil {
		return err
	}

	err = createFiles(p, virtualPaths)
	if err != nil {
		return err
	}

	for _, localePath := range virtualPaths {

		err := downloadAndWriteToFile(target, localePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		}
	}

	return nil
}

func PullTargetsFromConfig() (Targets, error) {
	content, err := ConfigContent()
	if err != nil {
		return nil, err
	}

	return parsePull(content)
}

func downloadAndWriteToFile(target Target, localePath *LocalePath) error {
	downloadParams := new(phraseapp.LocaleDownloadParams)
	downloadParams.FileFormat = target.Params.FileFormat

	fmt.Println("Downloading: ", localePath.LocaleId)
	res, err := phraseapp.LocaleDownload(target.ProjectId, localePath.LocaleId, downloadParams)
	if err != nil {
		fmt.Println("1", err)
		return err
	}
	fh, err := os.OpenFile(localePath.Path, os.O_WRONLY, 0700)
	if err != nil {
		return err
	}
	defer fh.Close()

	fmt.Println(string(res))
	_, err = fh.Write(res)
	if err != nil {
		fmt.Println("3", err)
		return err
	}
	return nil
}

// locale File handling
func createFiles(p *PhrasePath, virtualPaths LocalePaths) error {
	for _, localePath := range virtualPaths {
		defaultName := DefaultFileName(p.Mode, localePath.Path)
		if defaultName != "" {
			localePath.Path = path.Join(localePath.Path, p.Separator, defaultName)
		}
		fmt.Println("Creating: ", localePath.Path)
		err := CreateFile(localePath.Path)
		if err != nil {
			return err
		}
	}
	return nil
}

func expandPathsWithLocale(p *PhrasePath, target Target, locales []*phraseapp.Locale) (LocalePaths, error) {
	switch {
	case p.LocaleTagInPath:
		newFiles, err := FilePathsWithLocales(p, locales)
		if err != nil {
			return nil, err
		}
		return newFiles, nil

	default:
		return expandPathWithLocale(p, target, locales)
	}

}

func expandPathWithLocale(p *PhrasePath, target Target, locales []*phraseapp.Locale) (LocalePaths, error) {
	absPath, err := filepath.Abs(p.UserPath)
	if err != nil {
		return nil, err
	}

	if target.Params != nil {
		localeId := localeIdForPath(target.Params.LocaleId, locales)
		if localeId == "" {
			return nil, fmt.Errorf("locale specified in your path did not match any remote locales")
		}

		localePath := []*LocalePath{&LocalePath{Path: absPath, LocaleId: localeId}}
		return localePath, nil
	}

	return nil, fmt.Errorf("no target locale id specified")
}

func localeIdForPath(localeId string, locales []*phraseapp.Locale) string {
	for _, locale := range locales {
		if localeId == locale.Id {
			return locale.Id
		}
	}
	return ""
}

// File handling
func fileGlobbingPull(p *PhrasePath, paths LocalePaths) (LocalePaths, error) {
	switch {
	case p.Mode == "":
		return paths, nil

	case p.Mode == "*":
		if p.LocaleTagInFile() {
			return singleDirectoryPull(p, paths)
		} else {
			return paths, nil
		}

	default:
		return paths, nil
	}
}

func singleDirectoryPull(p *PhrasePath, paths LocalePaths) (LocalePaths, error) {
	extendedPaths := []*LocalePath{}
	for _, localePath := range paths {

		pathsPerDirectory, err := SingleDirectoryStrategy(localePath.Path+"/", "")

		if err != nil {
			return nil, err
		}

		for _, aPath := range pathsPerDirectory {
			localePath := &LocalePath{Path: aPath, LocaleId: localePath.LocaleId}
			extendedPaths = append(extendedPaths, localePath)
		}

	}
	return extendedPaths, nil
}

// Parsing
type PullConfig struct {
	Phraseapp struct {
		AccessToken string `yaml:"access_token"`
		ProjectId   string `yaml:"project_id"`
		Pull        struct {
			Targets Targets
		}
	}
}

func parsePull(yml string) (Targets, error) {
	var config *PullConfig

	err := yaml.Unmarshal([]byte(yml), &config)
	if err != nil {
		return nil, err
	}

	token := config.Phraseapp.AccessToken
	projectId := config.Phraseapp.ProjectId
	targets := config.Phraseapp.Pull.Targets

	for _, target := range targets {
		if target.ProjectId == "" {
			target.ProjectId = projectId
		}
		if target.AccessToken == "" {
			target.AccessToken = token
		}
	}

	return targets, nil
}
