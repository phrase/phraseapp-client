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
			printErr(err, "")
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

		if !localeFile.ExistsRemote {
			localeDetails, err := source.createLocale(localeFile)
			if err == nil {
				localeFile.Id = localeDetails.Id
				localeFile.RFC = localeDetails.Code
				localeFile.Name = localeDetails.Name
			}
		}

		err = source.uploadFile(localeFile)
		if err != nil {
			printErr(err, "")
		} else {
			sharedMessage("push", localeFile)
		}

	}

	return nil
}

func (source *Source) createLocale(localeFile *LocaleFile) (*phraseapp.LocaleDetails, error) {
	localeParams := new(phraseapp.LocaleParams)

	if localeFile.RFC != "" {
		localeParams.Code = localeFile.RFC
	}
	if localeFile.Name != "" {
		localeParams.Name = localeFile.Name
	} else {
		localeParams.Name = localeFile.RFC
	}

	localeDetails, err := phraseapp.LocaleCreate(source.ProjectId, localeParams)
	if err != nil {
		return nil, err
	}
	return localeDetails, nil
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

func (source *Source) LocaleFiles() (LocaleFiles, error) {
	source.Extension = filepath.Ext(source.File)

	recursiveFiles := []string{}
	if strings.Contains(source.File, "**") {
		rec, err := source.recurse()
		if err != nil {
			return nil, err
		}
		recursiveFiles = rec
	}

	globFiles, err := source.glob()
	if err != nil {
		return nil, err
	}

	filePaths := []string{}
	for _, f := range globFiles {
		if !Contains(filePaths, f) {
			filePaths = append(filePaths, f)
		}
	}
	for _, f := range recursiveFiles {
		if !Contains(filePaths, f) {
			filePaths = append(filePaths, f)
		}
	}

	var localeFiles LocaleFiles
	for _, path := range filePaths {
		localeFile, err := source.generateLocaleForFile(path)
		if err != nil {
			return nil, err
		}
		localeFiles = append(localeFiles, localeFile)
	}
	return localeFiles, nil
}

func (source *Source) generateLocaleForFile(path string) (*LocaleFile, error) {

	taggedMatches := source.findTaggedMatches(path)
	name := taggedMatches["locale_name"]
	rfc := taggedMatches["locale_code"]
	tag := taggedMatches["tag"]

	lc := &LocaleFile{}
	if name != "" {
		lc.Name = name
	}

	if rfc != "" {
		lc.RFC = rfc
	}

	if tag != "" {
		lc.Tag = tag
	}

	locale := source.getRemoteLocaleForLocaleFile(lc)
	if locale != nil {
		lc.ExistsRemote = true
		lc.RFC = locale.Code
		lc.Name = locale.Name
		lc.Id = locale.Id
	}

	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	lc.Path = absolutePath

	return lc, nil
}

func (source *Source) findTaggedMatches(path string) map[string]string {
	re := regexp.MustCompile("<(locale_name|tag|locale_code)>")

	separator := string(os.PathSeparator)
	taggedMatches := map[string]string{}

	// config/locale/<locale_code>.yml -> ["config", "locale", "<locale_code>.yml"]
	parts := strings.Split(source.File, separator)
	for _, part := range parts {
		if !re.MatchString(part) {
			continue
		}

		// <locale_code>.yml -> (?P<locale_code>.+).yml
		match := part
		group := re.FindString(part)
		if group == "" {
			continue
		}
		replacer := fmt.Sprintf("(?P%s.+)", group)
		match = strings.Replace(match, group, replacer, 1)

		reMatcher := regexp.MustCompile(match)
		namedMatches := reMatcher.SubexpNames()
		subMatches := reMatcher.FindStringSubmatch(path)
		for i, subMatch := range subMatches {
			if subMatch != "" {
				split := strings.Split(subMatch, separator)

				// the match is either from start or end of path
				// res/values-en/Strings.xml -> en/Strings.xml
				newMatch := split[0]
				if strings.HasPrefix(match, replacer) {
					// config/en.lproj -> config/en
					newMatch = split[len(split)-1]
				}

				taggedMatches[namedMatches[i]] = newMatch
			}
		}
	}

	return taggedMatches
}

func (source *Source) getRemoteLocaleForLocaleFile(localeFile *LocaleFile) *phraseapp.Locale {
	for _, remote := range source.RemoteLocales {
		if source.Params != nil {
			localeId := source.Params.LocaleId
			if remote.Id == localeId || remote.Name == localeId {
				return remote
			}
		}
		if remote.Name == localeFile.Name || remote.Code == localeFile.RFC {
			return remote
		}
	}
	return nil
}

func (source *Source) fileWithoutPlaceholder() string {
	re := regexp.MustCompile("<(locale_name|tag|locale_code)>")
	return strings.TrimSuffix(re.ReplaceAllString(source.File, "*"), source.Extension)
}

func (source *Source) extensionWithoutPlaceholder() string {
	re := regexp.MustCompile("<(locale_name|tag|locale_code)>")
	if re.MatchString(source.Extension) {
		return ""
	}
	return "*" + source.Extension
}

func (source *Source) glob() ([]string, error) {
	pattern := source.fileWithoutPlaceholder() + source.extensionWithoutPlaceholder()

	files, err := filepath.Glob(pattern)

	if Debug {
		fmt.Println("Found", len(files), "files matching the source pattern", pattern)
	}
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (source *Source) recurse() ([]string, error) {
	files := []string{}
	err := filepath.Walk(source.root(), func(path string, f os.FileInfo, err error) error {
		if strings.HasSuffix(f.Name(), source.Extension) {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func (source *Source) root() string {
	separator := string(os.PathSeparator)
	parts := strings.Split(source.File, separator)
	root := TakeWhile(parts, func(x string) bool { return x != "**" })
	return strings.Join(root, separator)
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
	} else if localeFile.RFC != "" {
		uploadParams.LocaleId = &(localeFile.RFC)
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
