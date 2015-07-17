package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/phrase/phraseapp-api-client/Godeps/_workspace/src/gopkg.in/yaml.v2"
	"github.com/phrase/phraseapp-go/phraseapp"
)

type Targets []*Target

type Target struct {
	File          string      `yaml:"file,omitempty"`
	ProjectId     string      `yaml:"project_id,omitempty"`
	AccessToken   string      `yaml:"access_token,omitempty"`
	FileFormat    string      `yaml:"file_format,omitempty"`
	Params        *PullParams `yaml:"params,omitempty"`
	RemoteLocales []*phraseapp.Locale
}

type PullParams struct {
	FileFormat               string                  `yaml:"file_format,omitempty"`
	LocaleId                 string                  `yaml:"locale_id,omitempty"`
	ConvertEmoji             *bool                   `yaml:"convert_emoji,omitempty"`
	FormatOptions            *map[string]interface{} `yaml:"format_options,omitempty"`
	IncludeEmptyTranslations *bool                   `yaml:"include_empty_translations,omitempty"`
	KeepNotranslateTags      *bool                   `yaml:"keep_notranslate_tags,omitempty"`
	Tag                      string                  `yaml:"tag,omitempty"`
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

func (t *Target) GetTag() string {
	if t.Params != nil {
		return t.Params.Tag
	}
	return ""
}

func pullCommand() error {
	targets, err := TargetsFromConfig()
	if err != nil {
		return err
	}

	pulledFiles := []string{}
	for _, target := range targets {
		pulledFilesAfterTargetPull, err := target.Pull(pulledFiles)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		}
		pulledFiles = pulledFilesAfterTargetPull
	}
	return nil
}

func (target *Target) Pull(pulledFiles []string) ([]string, error) {
	Authenticate()

	pathComponents, err := ExtractPathComponents(target.File)
	if err != nil {
		return nil, err
	}

	target.RemoteLocales, err = RemoteLocales(target.ProjectId)
	if err != nil {
		return nil, err
	}

	locale := &Locale{Id: target.GetLocaleId(), Tag: target.GetTag()}
	localeToPathMapping, err := pathComponents.ExpandPathsWithLocale(target.RemoteLocales, locale)
	if err != nil {
		return nil, err
	}

	for _, localeToPath := range localeToPathMapping {

		if contains(pulledFiles, localeToPath.Path) {
			continue
		}
		pulledFiles = append(pulledFiles, localeToPath.Path)

		err := createFile(localeToPath.Path)
		if err != nil {
			return nil, err
		}

		sharedMessage("pull", localeToPath)

		err = target.DownloadAndWriteToFile(localeToPath)
		if err != nil {
			printErr(err, fmt.Sprint("for %s", localeToPath.Path))
		}
	}

	return pulledFiles, nil
}

func TargetsFromConfig() (Targets, error) {
	content, err := ConfigContent()
	if err != nil {
		return nil, err
	}

	return parsePull(content)
}

func (target *Target) DownloadAndWriteToFile(locale *Locale) error {
	downloadParams := target.setDownloadParams(locale)

	params := target.Params
	localeId := ""
	if params != nil && params.LocaleId != "" {
		localeId = params.LocaleId
	} else {
		localeId = locale.Id
	}

	res, err := phraseapp.LocaleDownload(target.ProjectId, localeId, downloadParams)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(locale.Path, res, 0700)
	if err != nil {
		return err
	}
	return nil
}

func (target *Target) setDownloadParams(locale *Locale) *phraseapp.LocaleDownloadParams {
	downloadParams := new(phraseapp.LocaleDownloadParams)
	downloadParams.FileFormat = target.FileFormat

	params := target.Params

	if target.Params == nil {
		return downloadParams
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

	tag := params.Tag
	if tag != "" {
		downloadParams.Tag = &tag
	}

	return downloadParams
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
func createFile(path string) error {
	err := exists(path)
	if err != nil {
		absDir := filepath.Dir(path)
		err := exists(absDir)
		if err != nil {
			os.MkdirAll(absDir, 0700)
		}

		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
	}
	return nil
}

func exists(absPath string) error {
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return fmt.Errorf("no such file or directory:", absPath)
	}
	return nil
}
