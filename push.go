package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/mgutz/ansi"
	"github.com/phrase/phraseapp-api-client/Godeps/_workspace/src/gopkg.in/yaml.v2"
	"github.com/phrase/phraseapp-go/phraseapp"
)

type PushCommand struct {
	Verbose bool `cli:"opt --verbose default=false"`
}

func (cmd *PushCommand) Run() error {
	if cmd.Verbose {
		Debug = true
	}

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

type Sources []*Source

type Source struct {
	File           string      `yaml:"file,omitempty"`
	ProjectId      string      `yaml:"project_id,omitempty"`
	AccessToken    string      `yaml:"access_token,omitempty"`
	FileFormat     string      `yaml:"file_format,omitempty"`
	Params         *PushParams `yaml:"params,omitempty"`
	RemoteLocales  []*phraseapp.Locale
	PathComponents *PathComponents
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

func (source *Source) Push() error {
	Authenticate()

	remoteLocales, err := RemoteLocales(source.ProjectId)
	if err != nil {
		return err
	}
	source.RemoteLocales = remoteLocales

	localeFiles, err := source.LocaleFiles()
	if err != nil {
		return err
	}

	for _, localeFile := range localeFiles {
		fmt.Println("Uploading", localeFile.RelPath())
		err = source.uploadFile(localeFile)
		if err != nil {
			printErr(err, "")
		}
		sharedMessage("push", localeFile)
	}

	return nil
}

func (source *Source) uploadFile(localeFile *LocaleFile) error {
	uploadParams, err := source.setUploadParams(localeFile)
	if err != nil {
		return err
	}

	aUpload, err := phraseapp.UploadCreate(source.ProjectId, uploadParams)
	if err != nil {
		return err
	}

	printSummary(&aUpload.Summary)

	return nil
}

func (source *Source) IsDir() bool {
	return source.PathComponents.IsDir
}

func (source *Source) GlobPattern() string {
	return source.PathComponents.GlobPattern
}

func (source *Source) LocaleFiles() (LocaleFiles, error) {
	source.Extension = filepath.Ext(source.File)

	filePaths, err := source.glob()
	if err != nil {
		return nil, err
	}
	var localeFiles LocaleFiles
	for _, path := range filePaths {
		localeFile := source.generateLocaleForFile(path)
		localeFiles = append(localeFiles, localeFile)
	}
	return localeFiles, nil
}

func (source *Source) generateLocaleForFile(path string) *LocaleFile {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		printErr(err, "")
		return nil
	}

	return &LocaleFile{Path: absolutePath}
}

func (source *Source) FileWithoutPlaceholder() string {
	re := regexp.MustCompile("<(locale_name|tag|locale_code)>")
	return strings.TrimSuffix(re.ReplaceAllString(source.File, "*"), source.Extension)
}

func (source *Source) glob() ([]string, error) {
	pattern := source.FileWithoutPlaceholder() + "*" + source.Extension
	files, err := filepath.Glob(pattern)

	if Debug {
		fmt.Println("Found", len(files), "files matching the source pattern", pattern)
	}
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

func (source *Source) setUploadParams(localeFile *LocaleFile) (*phraseapp.LocaleFileImportParams, error) {
	uploadParams := new(phraseapp.LocaleFileImportParams)
	uploadParams.File = localeFile.Path
	uploadParams.FileFormat = &source.FileFormat

	if localeFile.Id != "" {
		uploadParams.LocaleId = &(localeFile.Id)
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
	newItems := []int64{summary.LocalesCreated, summary.TranslationsUpdated, summary.TranslationKeysCreated, summary.TranslationsCreated}
	var changed bool
	for _, item := range newItems {
		if item > 0 {
			changed = true
		}
	}
	if changed || Debug {
		localesCreated := joinMessage("Locales created: ", fmt.Sprintf("%d", summary.LocalesCreated))
		keysCreated := joinMessage("Keys created: ", fmt.Sprintf("%d", summary.TranslationKeysCreated))
		translationsCreated := joinMessage("Translations created: ", fmt.Sprintf("%d", summary.TranslationsCreated))
		translationsUpdated := joinMessage("Translations updated: ", fmt.Sprintf("%d", summary.TranslationsUpdated))
		formatted := fmt.Sprintf("%s - %s - %s - %s", localesCreated, keysCreated, translationsCreated, translationsUpdated)
		fmt.Println(formatted)
	}
}

func joinMessage(msg, stat string) string {
	green := ansi.ColorCode("green+b:black")
	reset := ansi.ColorCode("reset")
	return strings.Join([]string{msg, green, stat, reset}, "")
}
