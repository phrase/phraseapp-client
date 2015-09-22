package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/daviddengcn/go-colortext"
	"github.com/phrase/phraseapp-go/phraseapp"
)

type PushCommand struct {
	Credentials
}

func (cmd *PushCommand) Run() error {
	if cmd.Debug {
		// suppresses content output
		cmd.Debug = false
		Debug = true
	}
	client, err := ClientFromCmdCredentials(cmd.Credentials)
	if err != nil {
		return err
	}

	sources, err := SourcesFromConfig(cmd)
	if err != nil {
		return err
	}

	for _, source := range sources {
		err := source.Push(client)
		if err != nil {
			return err
		}
	}
	return nil
}

type Sources []*Source

type Source struct {
	File          string      `yaml:"file,omitempty"`
	ProjectID     string      `yaml:"project_id,omitempty"`
	AccessToken   string      `yaml:"access_token,omitempty"`
	FileFormat    string      `yaml:"file_format,omitempty"`
	Params        *PushParams `yaml:"params"`
	RemoteLocales []*phraseapp.Locale
	Extension     string
}

type PushParams struct {
	FileFormat   string `yaml:"file_format,omitempty"`
	LocaleID     string `yaml:"locale_id,omitempty"`
	ConvertEmoji *bool  `yaml:"convert_emoji,omitempty"`
	//FormatOptions      *map[string]interface{} `yaml:"format_options,omitempty"`
	SkipUnverification *bool   `yaml:"skip_unverification,omitempty"`
	SkipUploadTags     *bool   `yaml:"skip_upload_tags,omitempty"`
	Tags               *string `yaml:"tags,omitempty"`
	UpdateTranslations *bool   `yaml:"update_translations,omitempty"`
}

var separator = string(os.PathSeparator)

func (source *Source) Push(client *phraseapp.Client) error {
	if err := CheckPreconditions(source.File); err != nil {
		return err
	}

	source.Extension = filepath.Ext(source.File)

	remoteLocales, err := RemoteLocales(client, source.ProjectID)
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
			localeDetails, err := source.createLocale(client, localeFile)
			if err == nil {
				localeFile.ID = localeDetails.ID
				localeFile.RFC = localeDetails.Code
				localeFile.Name = localeDetails.Name
			}
		}

		err = source.uploadFile(client, localeFile)
		if err != nil {
			return err
		} else {
			sharedMessage("push", localeFile)
		}

		if Debug {
			fmt.Fprintln(os.Stderr, strings.Repeat("-", 10))
		}

	}

	return nil
}

func (source *Source) createLocale(client *phraseapp.Client, localeFile *LocaleFile) (*phraseapp.LocaleDetails, error) {
	localeParams := new(phraseapp.LocaleParams)

	if localeFile.Name != "" {
		localeParams.Name = &localeFile.Name
	} else if localeFile.RFC != "" {
		localeParams.Name = &localeFile.RFC
	}

	localeName := source.replacePlaceholderInParams(localeFile)
	if localeName != localeFile.RFC {
		localeParams.Name = &localeName
	}

	if localeFile.RFC != "" {
		localeParams.Code = &localeFile.RFC
	}

	localeDetails, err := client.LocaleCreate(source.ProjectID, localeParams)
	if err != nil {
		return nil, err
	}
	return localeDetails, nil
}

func (source *Source) replacePlaceholderInParams(localeFile *LocaleFile) string {
	if localeFile.RFC != "" && strings.Contains(source.GetLocaleID(), "<locale_code>") {
		return strings.Replace(source.GetLocaleID(), "<locale_code>", localeFile.RFC, 1)
	}
	return ""
}

func (source *Source) uploadFile(client *phraseapp.Client, localeFile *LocaleFile) error {
	uploadParams, err := source.setUploadParams(localeFile)
	if err != nil {
		return err
	}

	if Debug {
		fmt.Fprintln(os.Stdout, "Source file pattern:", source.File)
		fmt.Fprintln(os.Stdout, "Actual file location:", localeFile.Path)
	}

	aUpload, err := client.UploadCreate(source.ProjectID, uploadParams)
	if err != nil {
		return err
	}

	printSummary(&aUpload.Summary)

	return nil
}

func (source *Source) LocaleFiles() (LocaleFiles, error) {
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

	parser := Parser{
		SourceFile: source.File,
		Extension:  source.Extension,
	}
	parser.Initialize()

	if err := parser.Search(); err != nil {
		return nil, err
	}

	var localeFiles LocaleFiles
	for _, path := range filePaths {

		if !parser.MatchesPath(path) {
			continue
		}

		temporaryLocaleFile, err := parser.Eval(path)
		if err != nil {
			return nil, err
		}

		localeFile, err := source.generateLocaleForFile(temporaryLocaleFile, path)
		if err != nil {
			return nil, err
		}

		if Debug {
			fmt.Println(fmt.Sprintf(
				"RFC:'%s', Name:'%s', Tag;'%s', Pattern:'%s'",
				localeFile.RFC, localeFile.Name, localeFile.Tag, parser.Matcher.String(),
			))
		}

		localeFiles = append(localeFiles, localeFile)
	}

	if len(localeFiles) <= 0 {
		return nil, fmt.Errorf("file pattern did not identify any files on your system!")
	}
	return localeFiles, nil
}

func (source *Source) generateLocaleForFile(localeFile *LocaleFile, path string) (*LocaleFile, error) {
	locale := source.getRemoteLocaleForLocaleFile(localeFile)
	if locale != nil {
		localeFile.ExistsRemote = true
		localeFile.RFC = locale.Code
		localeFile.Name = locale.Name
		localeFile.ID = locale.ID
	}

	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	localeFile.Path = absolutePath

	return localeFile, nil
}

func (source *Source) getRemoteLocaleForLocaleFile(localeFile *LocaleFile) *phraseapp.Locale {
	for _, remote := range source.RemoteLocales {
		if remote.Name == source.GetLocaleID() || remote.ID == source.GetLocaleID() {
			return remote
		}

		localeName := source.replacePlaceholderInParams(localeFile)
		if localeName != "" && strings.Contains(remote.Name, localeName) {
			return remote
		}

		if remote.Name == localeFile.Name {
			return remote
		}
	}
	return nil
}

func (source *Source) fileWithoutPlaceholder() string {
	return strings.TrimSuffix(placeholderRegexp.ReplaceAllString(source.File, "*"), source.Extension)
}

func (source *Source) extensionWithoutPlaceholder() string {
	if placeholderRegexp.MatchString(source.Extension) {
		return ""
	}
	return "*" + source.Extension
}

func (source *Source) glob() ([]string, error) {
	pattern := source.fileWithoutPlaceholder() + source.extensionWithoutPlaceholder()

	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	if Debug {
		fmt.Fprintln(os.Stderr, "Found", len(files), "files matching the source pattern", pattern)
	}

	return files, nil
}

func (source *Source) recurse() ([]string, error) {
	files := []string{}
	err := filepath.Walk(source.root(), func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("%s for pattern: %s", err, source.File)
		}
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
	parts := strings.Split(source.File, separator)
	rootParts := TakeWhile(parts, func(x string) bool {
		return x != "**"
	})
	root := strings.Join(rootParts, separator)
	if root == "" {
		root = "."
	}
	return root
}

// Config, params and printing
type PushConfig struct {
	Phraseapp struct {
		AccessToken string `yaml:"access_token"`
		ProjectID   string `yaml:"project_id"`
		FileFormat  string `yaml:"file_format,omitempty"`
		Push        struct {
			Sources Sources
		}
	}
}

func SourcesFromConfig(cmd *PushCommand) (Sources, error) {
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
	if cmd.Token != "" {
		token = cmd.Token
	}
	projectId := config.Phraseapp.ProjectID
	fileFormat := config.Phraseapp.FileFormat

	if &config.Phraseapp.Push == nil || config.Phraseapp.Push.Sources == nil {
		return nil, fmt.Errorf("no sources for upload specified")
	}

	sources := config.Phraseapp.Push.Sources

	validSources := []*Source{}
	for _, source := range sources {
		if source == nil {
			continue
		}
		if source.ProjectID == "" {
			source.ProjectID = projectId
		}
		if source.AccessToken == "" {
			source.AccessToken = token
		}
		if source.FileFormat == "" {
			source.FileFormat = fileFormat
		}
		validSources = append(validSources, source)
	}

	if len(validSources) <= 0 {
		return nil, fmt.Errorf("no sources could be identified! Refine the sources list in your config")
	}

	return validSources, nil
}

func (source *Source) GetLocaleID() string {
	if source.Params != nil {
		return source.Params.LocaleID
	}
	return ""
}

func (source *Source) setUploadParams(localeFile *LocaleFile) (*phraseapp.LocaleFileImportParams, error) {
	uploadParams := new(phraseapp.LocaleFileImportParams)
	uploadParams.File = &localeFile.Path
	uploadParams.FileFormat = &source.FileFormat

	if localeFile.ID != "" {
		uploadParams.LocaleID = &localeFile.ID
	} else if localeFile.RFC != "" {
		uploadParams.LocaleID = &localeFile.RFC
	}

	if localeFile.Tag != "" {
		uploadParams.Tags = &localeFile.Tag
	}

	if source.Params == nil {
		return uploadParams, nil
	}

	params := source.Params

	localeID := params.LocaleID
	if localeID != "" {
		uploadParams.LocaleID = &localeID
	}

	format := params.FileFormat
	if format != "" {
		uploadParams.FileFormat = &format
	}

	convertEmoji := params.ConvertEmoji
	if convertEmoji != nil {
		uploadParams.ConvertEmoji = convertEmoji
	}

	//	formatOptions := params.FormatOptions
	//	if formatOptions != nil {
	//		uploadParams.FormatOptions = formatOptions
	//	}

	skipUnverification := params.SkipUnverification
	if skipUnverification != nil {
		uploadParams.SkipUnverification = skipUnverification
	}

	skipUploadTags := params.SkipUploadTags
	if skipUploadTags != nil {
		uploadParams.SkipUploadTags = skipUploadTags
	}

	tags := params.Tags
	if tags != nil && uploadParams.Tags == nil {
		uploadParams.Tags = tags
	}

	updateTranslations := params.UpdateTranslations
	if updateTranslations != nil {
		uploadParams.UpdateTranslations = updateTranslations
	}

	return uploadParams, nil
}

// Parser Algorithm, Stateful
// parser := &Parser{ SourceFile: file, Extension: ext }
// parser.Initialize()
// parser.Search()
// parser.Eval( path )
type Parser struct {
	SourceFile string
	Extension  string
	Tokens     []string
	Buffer     []string
	Matcher    *regexp.Regexp
}

func (parser *Parser) Initialize() {
	tokens := []string{}
	for _, token := range strings.Split(parser.SourceFile, separator) {
		if token == "." || token == "" {
			continue
		}
		tokens = append(tokens, token)
	}
	parser.Tokens = tokens
	parser.Buffer = []string{}
}

func (parser *Parser) Search() error {
	for _, token := range parser.Tokens {
		head := strings.Replace(token, ".", "_PHRASEAPP_REGEXP_DOT_", 1)
		// if next is part of the placeholder dsl then it can be expanded and will be
		// converted to a regexp, else we found a path part that is already a static string
		nextRegexp := parser.ConvertToRegexp(head)
		if nextRegexp != "" {
			head = nextRegexp

			// if it is the end of the path and matching the extension
			// e.g. config/.yml we convert it to config/.*.yml
		} else if strings.TrimSpace(token) == parser.Extension {
			head = ".*" + head
		}

		head = parser.SanitizeRegexp(head)
		head = strings.Replace(head, "_PHRASEAPP_REGEXP_DOT_", "[.]", 1)
		parser.Buffer = append(parser.Buffer, head)
	}

	fileAsRegexp := strings.Join(parser.Buffer, separator)
	matcherString := strings.Trim(fileAsRegexp, separator)

	reMatcher, err := regexp.Compile(matcherString)
	if err != nil {
		return err
	}

	// valid regular expression
	parser.Matcher = reMatcher

	return nil
}

func (parser *Parser) ConvertToRegexp(part string) string {
	groups := placeholderRegexp.FindAllString(part, -1)
	if len(groups) <= 0 {
		return ""
	}
	for _, group := range groups {
		replacer := fmt.Sprintf("(?P%s.+)", group)
		part = strings.Replace(part, group, replacer, 1)
	}
	return part
}

func (parser *Parser) SanitizeRegexp(token string) string {
	newToken := strings.Replace(token, "**", "__PHRASE_DOUBLE_STAR__", -1)
	newToken = strings.Replace(newToken, "*", ".*", -1)
	newToken = strings.Replace(newToken, "..", ".", -1)
	newToken = strings.Replace(newToken, "__PHRASE_DOUBLE_STAR__", ".*", -1)
	return newToken
}

func (parser *Parser) MatchesPath(path string) bool {
	return parser.Matcher.MatchString(path)
}

func (parser *Parser) Eval(path string) (*LocaleFile, error) {
	tagged := parser.TagMatches(path)
	name := tagged["locale_name"]
	rfc := tagged["locale_code"]
	tag := tagged["tag"]

	localeFile := &LocaleFile{}
	if name != "" {
		localeFile.Name = name
	}

	if rfc != "" {
		localeFile.RFC = rfc
	}

	if tag != "" {
		localeFile.Tag = tag
	}

	return localeFile, nil
}

// @return named matches for the given path matching the file pattern
// example: {"locale_name" : "English", "locale_code" : "en-Gb", "tag" : "my_tag"}
func (parser *Parser) TagMatches(path string) map[string]string {
	tagged := map[string]string{}
	namedMatches := parser.Matcher.SubexpNames()
	subMatches := parser.Matcher.FindStringSubmatch(path)
	for i, subMatch := range subMatches {
		if subMatch != "" {
			if strings.Contains(subMatch, separator) {
				subSlice := strings.Split(subMatch, separator)
				subMatch = subSlice[len(subSlice)-1]
			}
			tagged[namedMatches[i]] = subMatch
		}
	}
	return tagged
}

// print out
func printSummary(summary *phraseapp.SummaryType) {
	newItems := []int64{
		summary.LocalesCreated,
		summary.TranslationsUpdated,
		summary.TranslationKeysCreated,
		summary.TranslationsCreated,
	}
	var changed bool
	for _, item := range newItems {
		if item > 0 {
			changed = true
		}
	}
	if changed || Debug {
		printMessage("Locales created: ", fmt.Sprintf("%d", summary.LocalesCreated))
		printMessage(" - Keys created: ", fmt.Sprintf("%d", summary.TranslationKeysCreated))
		printMessage(" - Translations created: ", fmt.Sprintf("%d", summary.TranslationsCreated))
		printMessage(" - Translations updated: ", fmt.Sprintf("%d", summary.TranslationsUpdated))
		fmt.Print("\n")
	}
}

func printMessage(msg, stat string) {
	fmt.Print(msg)
	ct.Foreground(ct.Green, true)
	fmt.Print(stat)
	ct.ResetColor()
}
