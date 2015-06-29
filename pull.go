package main

import (
	"fmt"
	"github.com/phrase/phraseapp-api-client/Godeps/_workspace/src/gopkg.in/yaml.v2"
	"github.com/phrase/phraseapp-go/phraseapp"
	"os"
)

type Targets []*Target

type Target struct {
	File        string      `yaml:"file,omitempty"`
	ProjectId   string      `yaml:"project_id,omitempty"`
	AccessToken string      `yaml:"access_token,omitempty"`
	Params      *PullParams `yaml:"params,omitempty"`
}

type PullParams struct {
	FileFormat               string                  `yaml:"file_format,omitempty"`
	LocaleId                 string                  `yaml:"locale_id,omitempty"`
	ConvertEmoji             *bool                   `yaml:"convert_emoji,omitempty"`
	FormatOptions            *map[string]interface{} `yaml:"format_options,omitempty"`
	IncludeEmptyTranslations *bool                   `yaml:"include_empty_translations,omitempty"`
	KeepNotranslateTags      *bool                   `yaml:"keep_notranslate_tags,omitempty"`
	TagId                    *string                 `yaml:"tag_id,omitempty"`
}

func (t *Target) GetFormat() string {
	if t.Params != nil {
		return t.Params.FileFormat
	}
	return ""
}

func (t *Target) GetLocaleId() string {
	if t.Params != nil {
		return t.Params.LocaleId
	}
	return ""
}

func (target *Target) Pull() error {
	Authenticate()

	p := PathComponents(target.File)

	locales, err := phraseapp.LocalesList(target.ProjectId, 1, 25)
	if err != nil {
		return err
	}

	localeToPathMapping, err := ExpandPathsWithLocale(p, target.GetLocaleId(), locales)
	if err != nil {
		return err
	}

	virtualPaths, err := LocaleFileGlob(p, target.GetFormat(), localeToPathMapping)
	if err != nil {
		return err
	}

	err = CreateFiles(p, virtualPaths, target.GetFormat())
	if err != nil {
		return err
	}

	for _, localeToPath := range virtualPaths {

		sharedMessage("pull", localeToPath)

		err := downloadAndWriteToFile(target, localeToPath)
		if err != nil {
			printError(err, fmt.Sprint(" for %s", localeToPath.Path))
		}
	}

	return nil
}

func pullMsg(a, b string) (string, string) {
	return b, a
}

func TargetsFromConfig() (Targets, error) {
	content, err := ConfigContent()
	if err != nil {
		return nil, err
	}

	return parsePull(content)
}

func downloadAndWriteToFile(target *Target, localePath *LocalePath) error {
	downloadParams, err := setDownloadParams(target, localePath)
	if err != nil {
		return err
	}

	params := target.Params
	localeId := ""
	if params != nil && params.LocaleId != "" {
		localeId = params.LocaleId
	} else {
		localeId = localePath.LocaleId
	}

	res, err := phraseapp.LocaleDownload(target.ProjectId, localeId, downloadParams)
	if err != nil {
		return err
	}
	fh, err := os.OpenFile(localePath.Path, os.O_WRONLY, 0700)
	if err != nil {
		return err
	}
	defer fh.Close()

	_, err = fh.Write(res)
	if err != nil {
		return err
	}
	return nil
}

func setDownloadParams(target *Target, localePath *LocalePath) (*phraseapp.LocaleDownloadParams, error) {
	downloadParams := new(phraseapp.LocaleDownloadParams)

	if target.Params == nil {
		return downloadParams, nil
	}

	params := target.Params

	format := params.FileFormat
	if format != "" {
		downloadParams.FileFormat = format
	}

	convertEmoji := params.ConvertEmoji
	if convertEmoji != nil {
		downloadParams.ConvertEmoji = convertEmoji
	}

	formatOptions := params.FormatOptions
	if formatOptions != nil {
		downloadParams.FormatOptions = formatOptions
	}

	includeEmptyTranslations := params.IncludeEmptyTranslations
	if includeEmptyTranslations != nil {
		downloadParams.IncludeEmptyTranslations = includeEmptyTranslations
	}

	keepNotranslateTags := params.KeepNotranslateTags
	if keepNotranslateTags != nil {
		downloadParams.KeepNotranslateTags = keepNotranslateTags
	}

	tagId := params.TagId
	if tagId != nil {
		downloadParams.TagId = tagId
	}

	return downloadParams, nil
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
