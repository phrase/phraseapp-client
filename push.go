package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mgutz/ansi"
	"github.com/phrase/phraseapp-api-client/Godeps/_workspace/src/gopkg.in/yaml.v2"
	"github.com/phrase/phraseapp-go/phraseapp"
)

type Sources []*Source

type Source struct {
	File           string      `yaml:"file,omitempty"`
	ProjectId      string      `yaml:"project_id,omitempty"`
	AccessToken    string      `yaml:"access_token,omitempty"`
	FileFormat     string      `yaml:"file_format,omitempty"`
	Params         *PushParams `yaml:"params,omitempty"`
	RemoteLocales  []*phraseapp.Locale
	PathComponents *PathComponents
	Root           string
	Extension      string
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

func pushCommand() error {
	sources, err := SourcesFromConfig()
	if err != nil {
		return err
	}

	for _, source := range sources {
		err := source.Push()
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		}
	}
	return nil
}

func (source *Source) Push() error {
	Authenticate()

	locales, err := source.Locales()
	if err != nil {
		return err
	}

	for _, locale := range locales {
		fmt.Printf("Trying to push %s", locale.Path)
		err = source.uploadFile(locale)
		if err != nil {
			printErr(err, "")
		}
		sharedMessage("push", locale)
	}

	return nil
}

func (source *Source) uploadFile(locale *Locale) error {
	uploadParams, err := source.setUploadParams(locale)
	if err != nil {
		return err
	}

	aUpload, err := phraseapp.UploadCreate(source.ProjectId, uploadParams)
	if err != nil {
		return err
	}

	printSummary(&aUpload.Summary)

	fmt.Printf("%%", aUpload)

	return nil
}

func (source *Source) IsDir() bool {
	return source.PathComponents.IsDir
}

func (source *Source) GlobPattern() string {
	return source.PathComponents.GlobPattern
}

func (source *Source) Locales() (Locales, error) {
	filePaths, err := source.glob()
	if err != nil {
		return nil, err
	}
	var locales Locales
	for _, path := range filePaths {
		locales = append(locales, &Locale{Path: path})
	}
	return locales, nil
}

func (source *Source) glob() ([]string, error) {
	files, err := filepath.Glob(source.Root + "*." + source.Extension)

	if err != nil {
		return nil, err
	}

	return files, nil
}

func SourcesFromConfig() (Sources, error) {
	content, err := ConfigContent()
	if err != nil {
		return nil, err
	}

	var config *PushConfig

	err = yaml.Unmarshal([]byte(content), &config)
	if err != nil {
		return nil, err
	}

	token := config.Phraseapp.AccessToken
	projectId := config.Phraseapp.ProjectId
	fileFormat := config.Phraseapp.FileFormat
	sources := *config.Phraseapp.Push.Sources

	for _, source := range sources {
		if source.ProjectId == "" {
			source.ProjectId = projectId
		}
		if source.AccessToken == "" {
			source.AccessToken = token
		}
		if source.FileFormat == "" {
			source.FileFormat = fileFormat
		}
	}

	return sources, nil
}

func (source *Source) setUploadParams(locale *Locale) (*phraseapp.LocaleFileImportParams, error) {
	uploadParams := new(phraseapp.LocaleFileImportParams)
	uploadParams.File = locale.Path
	uploadParams.FileFormat = &source.FileFormat

	if locale.Id != "" {
		uploadParams.LocaleId = &(locale.Id)
	}

	if source.Params == nil {
		return uploadParams, nil
	}

	params := source.Params

	localeId := params.LocaleId
	if localeId != "" {
		uploadParams.LocaleId = &localeId
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
		FileFormat  string `yaml:"file_format,omitempty"`
		Push        struct {
			Sources *Sources
		}
	}
}

func printSummary(summary *phraseapp.SummaryType) {
	localesCreated := joinMessage("Locales created: ", fmt.Sprintf("%d", summary.LocalesCreated))
	keysCreated := joinMessage("Keys created: ", fmt.Sprintf("%d", summary.TranslationKeysCreated))
	translationsCreated := joinMessage("Translations created: ", fmt.Sprintf("%d", summary.TranslationsCreated))
	translationsUpdated := joinMessage("Translations updated: ", fmt.Sprintf("%d", summary.TranslationsUpdated))
	formatted := fmt.Sprintf("%s - %s - %s - %s", localesCreated, keysCreated, translationsCreated, translationsUpdated)
	fmt.Println(formatted)
}

func joinMessage(msg, stat string) string {
	green := ansi.ColorCode("green+b:black")
	reset := ansi.ColorCode("reset")
	return strings.Join([]string{msg, green, stat, reset}, "")
}
