package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/daviddengcn/go-colortext"
	"github.com/phrase/phraseapp-api-client/Godeps/_workspace/src/gopkg.in/yaml.v2"
	"github.com/phrase/phraseapp-go/phraseapp"
)

type PushCommand struct {
	phraseapp.AuthCredentials
	DebugPush bool `cli:"opt --debug desc='Debug output (only push+pull)'"`
}

func (cmd *PushCommand) Run() error {
	Authenticate(&cmd.AuthCredentials)

	if cmd.DebugPush {
		Debug = true
	}

	sources, err := SourcesFromConfig(cmd)
	if err != nil {
		return err
	}

	for _, source := range sources {
		err := source.Push()
		if err != nil {
			return err
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

func (source *Source) GetLocaleId() string {
	if source.Params != nil {
		return source.Params.LocaleId
	}
	return ""
}

func (source *Source) Push() error {
	if strings.TrimSpace(source.File) == "" {
		return fmt.Errorf("file of source may not be empty")
	}

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
			return err
		} else {
			sharedMessage("push", localeFile)
		}

		if Debug {
			fmt.Println(strings.Repeat("-", 10))
		}

	}

	return nil
}

func (source *Source) createLocale(localeFile *LocaleFile) (*phraseapp.LocaleDetails, error) {
	localeParams := new(phraseapp.LocaleParams)

	if localeFile.Name != "" {
		localeParams.Name = localeFile.Name
	} else {
		localeParams.Name = localeFile.RFC
	}

	localeName := source.replacePlaceholderInParams(localeFile)
	if localeName != localeFile.RFC {
		localeParams.Name = localeName
	}

	if localeFile.RFC != "" {
		localeParams.Code = localeFile.RFC
	}

	localeDetails, err := phraseapp.LocaleCreate(source.ProjectId, localeParams)
	if err != nil {
		return nil, err
	}
	return localeDetails, nil
}

func (source *Source) replacePlaceholderInParams(localeFile *LocaleFile) string {
	if localeFile.RFC != "" && strings.Contains(source.GetLocaleId(), "<locale_code>") {
		return strings.Replace(source.GetLocaleId(), "<locale_code>", localeFile.RFC, 1)
	}
	return ""
}

func (source *Source) uploadFile(localeFile *LocaleFile) error {
	uploadParams, err := source.setUploadParams(localeFile)
	if err != nil {
		return err
	}

	if Debug {
		fmt.Println("Source file pattern:", source.File)
		fmt.Println("Actual File Location", uploadParams.File)
		if uploadParams.LocaleId != nil {
			fmt.Println("LocaleID/Name", *uploadParams.LocaleId)
		} else {
			fmt.Println("LocaleID/Name", nil)
		}

		if uploadParams.FileFormat != nil {
			fmt.Println("Format", *uploadParams.FileFormat)
		} else {
			fmt.Println("Format", nil)
		}
		fmt.Println("Tags", uploadParams.Tags)
		fmt.Println("Emoji", uploadParams.ConvertEmoji)
		if uploadParams.UpdateTranslations != nil {
			fmt.Println("UpdateTranslations", *uploadParams.UpdateTranslations)
		} else {
			fmt.Println("UpdateTranslations", nil)
		}
		fmt.Println("SkipUnverification", uploadParams.SkipUnverification)
		fmt.Println("FormatOpts", uploadParams.FormatOptions)
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

	taggedMatches := source.FindTaggedMatches(path)

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

func (source *Source) getRemoteLocaleForLocaleFile(localeFile *LocaleFile) *phraseapp.Locale {
	for _, remote := range source.RemoteLocales {
		if remote.Id == source.GetLocaleId() || remote.Name == source.GetLocaleId() {
			return remote
		}

		localeName := source.replacePlaceholderInParams(localeFile)
		if strings.Contains(remote.Name, localeName) {
			return remote
		}

		if remote.Name == localeFile.Name {
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
	separator := string(os.PathSeparator)
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
	projectId := config.Phraseapp.ProjectId
	fileFormat := config.Phraseapp.FileFormat

	if &config.Phraseapp.Push == nil || config.Phraseapp.Push.Sources == nil {
		return nil, fmt.Errorf("no sources for upload specified")
	}

	sources := *config.Phraseapp.Push.Sources

	validSources := []*Source{}
	for _, source := range sources {
		if source == nil {
			continue
		}
		if source.ProjectId == "" {
			source.ProjectId = projectId
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

func (source *Source) setUploadParams(localeFile *LocaleFile) (*phraseapp.LocaleFileImportParams, error) {
	uploadParams := new(phraseapp.LocaleFileImportParams)
	uploadParams.File = localeFile.Path
	uploadParams.FileFormat = &source.FileFormat

	if localeFile.Id != "" {
		uploadParams.LocaleId = &(localeFile.Id)
	} else if localeFile.RFC != "" {
		uploadParams.LocaleId = &(localeFile.RFC)
	}

	if localeFile.Tag != "" {
		uploadParams.Tags = []string{localeFile.Tag}
	}

	if source.Params == nil {
		return uploadParams, nil
	}

	params := source.Params

	localeId := params.LocaleId
	if localeId != "" && !strings.Contains(localeId, "<locale_code>") {
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
	if tags != nil && uploadParams.Tags == nil {
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
		printMessage("Locales created: ", fmt.Sprintf("%d", summary.LocalesCreated))
		printMessage("- Keys created: ", fmt.Sprintf("%d", summary.TranslationKeysCreated))
		printMessage("- Translations created: ", fmt.Sprintf("%d", summary.TranslationsCreated))
		printMessage("- Translations updated: ", fmt.Sprintf("%d", summary.TranslationsUpdated))
		fmt.Print("\n")
	}
}

func printMessage(msg, stat string) {
	fmt.Print(msg)
	ct.Foreground(ct.Green, true)
	fmt.Print(stat)
	ct.ResetColor()
}

// Hard parser algorithm... get a coffee..
type Parser struct {
	Path          string
	SourceFile    string
	Extension     string
	Tokens        []string
	LookBack      map[string]string
	LookAhead     map[string]string
	BackBuffer    []string
	ForwardBuffer []string
}

var separator = string(os.PathSeparator)
var tagRegex = regexp.MustCompile("<(locale_name|tag|locale_code)>")

func (source *Source) FindTaggedMatches(path string) map[string]string {
	parser := Parser{
		Path:       path,
		SourceFile: source.File,
		Extension:  source.Extension,
	}
	parser.Initialize()

	tagged := map[string]string{}
	for _, token := range parser.Tokens {

		parser.BackBuffer = parser.Search(
			[]string{},
			parser.LookBack,
			token,
		)

		parser.ForwardBuffer = parser.Search(
			[]string{},
			parser.LookAhead,
			token,
		)

		if len(parser.ForwardBuffer) > 0 || len(parser.BackBuffer) > 0 {
			matches := parser.Eval(token)
			tagged = updateMatches(tagged, matches)
		}
	}

	return tagged
}

func (parser *Parser) Initialize() {
	lookAhead := map[string]string{}
	lookBack := map[string]string{}
	parts := strings.Split(parser.SourceFile, separator)
	tokens := []string{}
	for idx, token := range parts {
		if token == "." || token == "" {
			continue
		}

		if tagRegex.MatchString(token) {
			lookAhead[token] = ""
			if len(parts) > idx+1 {
				lookAhead[token] = parts[idx+1]
			}
			lookBack[token] = ""
			if idx > 0 {
				if parts[idx-1] != "." {
					lookBack[token] = parts[idx-1]
				}
			}
		}

		tokens = append(tokens, token)
	}

	parser.Tokens = tokens
	parser.LookBack = lookBack
	parser.LookAhead = lookAhead
}

func (parser *Parser) Search(buffer []string, lookUp map[string]string, token string) []string {
	next := lookUp[token]
	nextRegexp := parser.SearchRegexp(next)
	if nextRegexp != "" {
		buffer = append(buffer, nextRegexp)
		return parser.Search(buffer, lookUp, next)
	} else {
		if strings.TrimSpace(next) == parser.Extension {
			buffer = append(buffer, ".*"+next)
		} else {
			buffer = append(buffer, next)
		}
	}
	return buffer
}

func (parser *Parser) Eval(token string) map[string]string {
	sanitizedBack := SanitizeBuffer(Reverse(parser.BackBuffer))
	backString := strings.Join(sanitizedBack, separator)

	sanitizedForward := SanitizeBuffer(parser.ForwardBuffer)
	forwardString := strings.Join(sanitizedForward, separator)

	token = Sanitize(token)
	tokenRegexp := parser.SearchRegexp(token)
	if tokenRegexp != "" {
		token = tokenRegexp
	}

	matcherString := strings.Trim(backString+separator+token+separator+forwardString, separator)

	return parser.TagMatches(matcherString)
}

func (parser *Parser) SearchRegexp(part string) string {
	group := tagRegex.FindString(part)
	if group == "" {
		return ""
	}
	replacer := fmt.Sprintf("(?P%s.+)", group)
	return strings.Replace(part, group, replacer, 1)
}

func (parser *Parser) TagMatches(matcherString string) map[string]string {
	tagged := map[string]string{}
	reMatcher := regexp.MustCompile(matcherString)
	namedMatches := reMatcher.SubexpNames()
	subMatches := reMatcher.FindStringSubmatch(parser.Path)
	for i, subMatch := range subMatches {
		if subMatch != "" {
			tagged[namedMatches[i]] = subMatch
		}
	}
	return tagged
}

func SanitizeBuffer(buffer []string) []string {
	newBuffer := []string{}
	for _, token := range buffer {
		newBuffer = append(newBuffer, Sanitize(token))
	}
	return newBuffer
}

func Sanitize(token string) string {
	token = strings.Replace(token, "**", ".*", -1)
	return token
}

func updateMatches(original, updater map[string]string) map[string]string {
	localeCode := updater["locale_code"]
	localeName := updater["locale_name"]
	tag := updater["tag"]

	if original["locale_code"] == "" {
		original["locale_code"] = strings.Trim(localeCode, separator)
	}

	if original["locale_name"] == "" {
		original["locale_name"] = strings.Trim(localeName, separator)
	}

	if original["tag"] == "" {
		original["tag"] = strings.Trim(tag, separator)
	}
	return original
}

func Reverse(seq []string) []string {
	for i, j := 0, len(seq)-1; i < j; i, j = i+1, j-1 {
		seq[i], seq[j] = seq[j], seq[i]
	}
	return seq
}
