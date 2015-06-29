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
	FileFormat  string      `yaml:"file_format,omitempty"`
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
	if t.Params != nil && t.Params.FileFormat != "" {
		return t.Params.FileFormat
	}
	if t.FileFormat != "" {
		return t.FileFormat
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

	virtualPaths, err := LocaleFileGlobPull(p, target.GetFormat(), localeToPathMapping, locales)
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

func LocaleFileGlobPull(p *PhrasePath, fileFormat string, paths LocalePaths, locales []*phraseapp.Locale) (LocalePaths, error) {
	switch {
	case p.Mode == "":
		return paths, nil

	case p.Mode == "*":
		if p.LocaleTagInPath {
			return expandSingleDirectory(p, paths, fileFormat)
		} else {
			return buildLocaleNamesForFile(p, fileFormat, paths, locales)
		}

	default:
		return paths, nil
	}
}

func TargetsFromConfig() (Targets, error) {
	content, err := ConfigContent()
	if err != nil {
		return nil, err
	}

	return parsePull(content)
}

func buildLocaleNamesForFile(p *PhrasePath, fileFormat string, paths LocalePaths, locales []*phraseapp.Locale) (LocalePaths, error) {
	newPaths := []*LocalePath{}
	singleDir := paths[0]
	for _, locale := range locales {
		newPath := singleDir.Path + p.Separator + locale.Name + "." + fileFormat
		localeToPathMapping := &LocalePath{Path: newPath, LocaleId: locale.Id, LocaleName: locale.Name, LocaleCode: locale.Code}
		newPaths = append(newPaths, localeToPathMapping)
	}
	return newPaths, nil
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
	downloadParams.FileFormat = target.FileFormat

	params := target.Params

	if target.Params == nil {
		return downloadParams, nil
	}

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
		FileFormat  string `yaml:"file_format,omitempty"`
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
	fileFormat := config.Phraseapp.FileFormat
	targets := config.Phraseapp.Pull.Targets

	for _, target := range targets {
		if target.ProjectId == "" {
			target.ProjectId = projectId
		}
		if target.AccessToken == "" {
			target.AccessToken = token
		}
		if target.FileFormat == "" {
			target.FileFormat = fileFormat
		}
	}

	return targets, nil
}
