package main

import (
	"bytes"
	"fmt"
	"github.com/phrase/phraseapp-api-client/Godeps/_workspace/src/gopkg.in/yaml.v2"
	"github.com/phrase/phraseapp-go/phraseapp"
	"os"
)

type Sources []*Source

type Source struct {
	File        string      `yaml:"file,omitempty"`
	ProjectId   string      `yaml:"project_id,omitempty"`
	AccessToken string      `yaml:"access_token,omitempty"`
	Params      *PushParams `yaml:"params,omitempty"`
}

type PushParams struct {
	FileFormat         string                  `yaml:"file_format,omitempty"`
	LocaleId           string                  `yaml:"locale_id,omitempty"`
	ConvertEmoji       *bool                   `yaml:"convert_emoji,omitempty"`
	FormatOptions      *map[string]interface{} `yaml:"format_options,omitempty"`
	SkipUnverification *bool                   `yaml:"skip_unverification,omitempty"`
	SkipUploadTags     *bool                   `yaml:"skip_upload_tags,omitempty"`
	Tags               []string                `yaml:"tags,omitempty"`
	UpdateTranslations *bool                   `yaml:"update_translations,omitempty"`
}

func (s *Source) GetFormat() string {
	if s.Params != nil {
		return s.Params.FileFormat
	}
	return ""
}

func (s *Source) GetLocaleId() string {
	if s.Params != nil {
		return s.Params.LocaleId
	}
	return ""
}

func (source *Source) Push() error {
	Authenticate()

	p := PathComponents(source.File)

	locales, err := phraseapp.LocalesList(source.ProjectId, 1, 25)
	if err != nil {
		return err
	}

	localeToPathMapping, err := ExpandPathsWithLocale(p, source.GetLocaleId(), locales)
	if err != nil {
		return err
	}

	virtualPaths, err := LocaleFileGlob(p, source.GetFormat(), localeToPathMapping)
	if err != nil {
		return err
	}

	for _, localeToPath := range virtualPaths {

		uploadMessaging(localeToPath)

		err = uploadFile(source, localeToPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		}
	}

	return nil
}

func SourcesFromConfig() (Sources, error) {
	content, err := ConfigContent()
	if err != nil {
		return nil, err
	}

	return parsePush(content)
}

func uploadFile(source *Source, localePath *LocalePath) error {
	uploadParams, err := setUploadParams(source, localePath)
	if err != nil {
		return err
	}

	aUpload, err := phraseapp.UploadCreate(source.ProjectId, uploadParams)
	if err != nil {
		return err
	}

	localesCreated := aUpload.Summary.LocalesCreated
	translationKeysCreated := aUpload.Summary.TranslationKeysCreated
	translationsCreated := aUpload.Summary.TranslationsCreated
	translationsUpdated := aUpload.Summary.TranslationsUpdated

	fmt.Println("Locales Created:", localesCreated,
		"Keys Created:", translationKeysCreated,
		"Translations Created:", translationsCreated,
		"Translations Updated:", translationsUpdated)
	return nil
}

func setUploadParams(source *Source, localePath *LocalePath) (*phraseapp.LocaleFileImportParams, error) {
	uploadParams := new(phraseapp.LocaleFileImportParams)
	uploadParams.File = localePath.Path

	if source.Params == nil {
		return uploadParams, nil
	}

	params := source.Params

	localeId := params.LocaleId
	if localeId != "" {
		uploadParams.LocaleId = &localeId
	} else {
		remoteLocaleId := localePath.LocaleId
		if remoteLocaleId != "" {
			uploadParams.LocaleId = &remoteLocaleId
		}
	}

	format := params.FileFormat
	if format != "" {
		uploadParams.FileFormat = &format
	}

	convertEmoji := params.ConvertEmoji
	if convertEmoji != nil {
		uploadParams.ConvertEmoji = convertEmoji
	}

	formatOptions := params.FormatOptions
	if formatOptions != nil {
		uploadParams.FormatOptions = formatOptions
	}

	skipUnverification := params.SkipUnverification
	if skipUnverification != nil {
		uploadParams.SkipUnverification = skipUnverification
	}

	skipUploadTags := params.SkipUploadTags
	if skipUploadTags != nil {
		uploadParams.SkipUploadTags = skipUploadTags
	}

	tags := params.Tags
	if tags != nil {
		uploadParams.Tags = tags
	}

	updateTranslations := params.UpdateTranslations
	if updateTranslations != nil {
		uploadParams.UpdateTranslations = updateTranslations
	}

	return uploadParams, nil
}

// Parsing
type PushConfig struct {
	Phraseapp struct {
		AccessToken string `yaml:"access_token"`
		ProjectId   string `yaml:"project_id"`
		Push        struct {
			Sources Sources
		}
	}
}

func parsePush(yml string) (Sources, error) {
	var config *PushConfig

	err := yaml.Unmarshal([]byte(yml), &config)
	if err != nil {
		return nil, err
	}

	token := config.Phraseapp.AccessToken
	projectId := config.Phraseapp.ProjectId
	sources := config.Phraseapp.Push.Sources

	for _, source := range sources {
		if source.ProjectId == "" {
			source.ProjectId = projectId
		}
		if source.AccessToken == "" {
			source.AccessToken = token
		}
	}

	return sources, nil
}

func uploadMessaging(localeToPath *LocalePath) {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Pushing: %s ", localeToPath.Path))
	if localeToPath.LocaleName != "" {
		buffer.WriteString(fmt.Sprintf("To: %s", localeToPath.LocaleName))
		if localeToPath.LocaleCode != "" {
			buffer.WriteString(fmt.Sprintf(" (%s)", localeToPath.LocaleCode))
		}
	} else if localeToPath.LocaleId != "" {
		buffer.WriteString(fmt.Sprintf("To: %s", localeToPath.LocaleId))
	}
	fmt.Println(buffer.String())
}
