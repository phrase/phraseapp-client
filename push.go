package main

import (
	"fmt"
	"github.com/mgutz/ansi"
	"github.com/phrase/phraseapp-api-client/Godeps/_workspace/src/gopkg.in/yaml.v2"
	"github.com/phrase/phraseapp-go/phraseapp"
	"os"
	"path/filepath"
	"strings"
)

type Sources []*Source

type Source struct {
	File        string      `yaml:"file,omitempty"`
	ProjectId   string      `yaml:"project_id,omitempty"`
	AccessToken string      `yaml:"access_token,omitempty"`
	FileFormat  string      `yaml:"file_format,omitempty"`
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
	if s.Params != nil && s.Params.FileFormat != "" {
		return s.Params.FileFormat
	}
	if s.FileFormat != "" {
		return s.FileFormat
	}
	return ""
}

func (s *Source) GetLocaleId() string {
	if s.Params != nil {
		return s.Params.LocaleId
	}
	return ""
}

func PushAll(sources Sources) error {
	alreadySeen := []string{}
	for _, source := range sources {
		newSeen, err := source.Push(alreadySeen)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		}
		alreadySeen = newSeen
	}
	return nil
}

func (source *Source) Push(alreadySeen []string) ([]string, error) {
	Authenticate()

	p := PathComponents(source.File)

	locales, err := phraseapp.LocalesList(source.ProjectId, 1, 25)
	if err != nil {
		return nil, err
	}

	localeToPathMapping, err := ExpandPathsWithLocale(p, source.GetLocaleId(), locales)
	if err != nil {
		return nil, err
	}

	virtualPaths, err := LocaleFileGlob(p, source.GetFormat(), localeToPathMapping)
	if err != nil {
		return nil, err
	}

	for _, localeToPath := range virtualPaths {

		if wasAlreadySeen(alreadySeen, localeToPath.Path) {
			continue
		}
		alreadySeen = append(alreadySeen, localeToPath.Path)

		sharedMessage("push", localeToPath)

		err = uploadFile(source, localeToPath)
		if err != nil {
			printErr(err, "")
		}
	}

	return alreadySeen, nil
}

func LocaleFileGlob(p *PhrasePath, fileFormat string, paths LocalePaths) (LocalePaths, error) {
	switch {
	case p.Mode == "":
		return paths, nil

	case p.Mode == "*":
		return expandSingleDirectory(p, paths, fileFormat)

	case p.Mode == "**/*":
		return recurseDirectory(fileFormat, paths)

	default:
		return paths, nil
	}
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

	printSummary(&aUpload.Summary)

	return nil
}

func setUploadParams(source *Source, localePath *LocalePath) (*phraseapp.LocaleFileImportParams, error) {
	uploadParams := new(phraseapp.LocaleFileImportParams)
	uploadParams.File = localePath.Path
	uploadParams.FileFormat = &source.FileFormat
	remoteLocaleId := localePath.LocaleId

	if remoteLocaleId != "" {
		uploadParams.LocaleId = &remoteLocaleId
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

// File expansion for * and **/* and ""
func expandSingleDirectory(p *PhrasePath, paths LocalePaths, fileFormat string) (LocalePaths, error) {
	expandedPaths := []*LocalePath{}
	for _, localePath := range paths {

		asDirectory := fmt.Sprintf("%s/", localePath.Path)
		pathsPerDirectory, err := glob(asDirectory, fileFormat)

		if err != nil {
			return nil, err
		}

		for _, absPath := range pathsPerDirectory {
			localeToPathMapping := CopyLocalePath(absPath, localePath)
			expandedPaths = append(expandedPaths, localeToPathMapping)
		}

	}
	return expandedPaths, nil
}

func recurseDirectory(fileFormat string, paths LocalePaths) (LocalePaths, error) {
	expandedPaths := []*LocalePath{}
	for _, localePath := range paths {
		newPaths, err := walk(localePath.Path, fileFormat)
		if err != nil {
			return nil, err
		}
		for _, newPath := range newPaths {
			localeToPathMapping := CopyLocalePath(newPath, localePath)
			expandedPaths = append(expandedPaths, localeToPathMapping)
		}
	}
	return expandedPaths, nil
}

func walk(root, fileFormat string) ([]string, error) {
	fileList := []string{}
	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if isLocaleFile(f.Name(), fileFormat) {
			fileList = append(fileList, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return fileList, nil
}

func glob(root, fileFormat string) ([]string, error) {
	files, err := filepath.Glob(root + "*")

	if err != nil {
		return nil, err
	}
	localeFiles := []string{}
	for _, f := range files {
		if fileFormat != "" {
			if isLocaleFile(f, fileFormat) {
				localeFiles = append(localeFiles, f)
			}
		} else {
			localeFiles = append(localeFiles, f)
		}
	}
	return localeFiles, nil
}

// Parsing
type PushConfig struct {
	Phraseapp struct {
		AccessToken string `yaml:"access_token"`
		ProjectId   string `yaml:"project_id"`
		FileFormat  string `yaml:"file_format,omitempty"`
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
	fileFormat := config.Phraseapp.FileFormat
	sources := config.Phraseapp.Push.Sources

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
