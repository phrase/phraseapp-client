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
	FileEncoding   string `yaml:"file_encoding,omitempty"`
	LocaleID     string `yaml:"locale_id,omitempty"`
	ConvertEmoji *bool  `yaml:"convert_emoji,omitempty"`
	//FormatOptions      *map[string]interface{} `yaml:"format_options,omitempty"`
	SkipUnverification *bool   `yaml:"skip_unverification,omitempty"`
	SkipUploadTags     *bool   `yaml:"skip_upload_tags,omitempty"`
	Tags               *string `yaml:"tags,omitempty"`
	UpdateTranslations *bool   `yaml:"update_translations,omitempty"`
}

var separator = string(os.PathSeparator)

func (source *Source) CheckPreconditions() error {
	if err := ValidPath(source.File, source.FileFormat, ""); err != nil {
		return err
	}

	duplicatedPlaceholders := []string{}
	for _, name := range []string{"<locale_name>", "<locale_code>", "<tag>"} {
		if strings.Count(source.File, name) > 1 {
			duplicatedPlaceholders = append(duplicatedPlaceholders, name)
		}
	}

	starCount := strings.Count(source.File, "*")
	recCount := strings.Count(source.File, "**")

	if recCount == 0 && starCount > 1 || starCount-(recCount*2) > 1 {
		duplicatedPlaceholders = append(duplicatedPlaceholders, "*")
	}

	if recCount > 1 {
		duplicatedPlaceholders = append(duplicatedPlaceholders, "**")
	}

	if len(duplicatedPlaceholders) > 0 {
		dups := strings.Join(duplicatedPlaceholders, ", ")
		return fmt.Errorf(fmt.Sprintf("%s can only occur once in a file pattern!", dups))
	}

	return nil
}

func (source *Source) Push(client *phraseapp.Client) error {
	if err := source.CheckPreconditions(); err != nil {
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
			} else {
				fmt.Printf("failed to create locale: %s\n", err)
				continue
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
	if localeFile.RFC == "" {
		return nil, fmt.Errorf("no locale code specified")
	}

	localeParams := new(phraseapp.LocaleParams)

	if localeFile.Name != "" {
		localeParams.Name = &localeFile.Name
	} else if localeFile.RFC != "" {
		localeParams.Name = &localeFile.RFC
	}

	localeName := source.replacePlaceholderInParams(localeFile)
	if localeName != "" && localeName != localeFile.RFC {
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

	reducer := Reducer{
		SourceFile: source.File,
		Extension:  source.Extension,
	}
	reducer.Initialize()

	if err := reducer.Reduce(); err != nil {
		return nil, err
	}

	var localeFiles LocaleFiles
	for _, path := range filePaths {
		if !reducer.MatchesPath(path) {
			continue
		}

		temporaryLocaleFile, err := reducer.Eval(path)
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
				localeFile.RFC, localeFile.Name, localeFile.Tag, reducer.Matcher.String(),
			))
		}

		localeFiles = append(localeFiles, localeFile)
	}

	if len(localeFiles) <= 0 {
		abs, err := filepath.Abs(source.File)
		if err != nil {
			return nil, err
		}
		errmsg := fmt.Sprintf("Could not find any files on your system that matches: '%s'", abs)
		ReportError("Push Error", errmsg)
		return nil, fmt.Errorf(errmsg)
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
			errmsg := fmt.Sprintf("%s for pattern: %s", err, source.File)
			ReportError("Push Error", errmsg)
			return fmt.Errorf(errmsg)
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
		errmsg := "no sources for upload specified"
		ReportError("Push Error", errmsg)
		return nil, fmt.Errorf(errmsg)
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
		errmsg := "no sources could be identified! Refine the sources list in your config"
		ReportError("Push Error", errmsg)
		return nil, fmt.Errorf(errmsg)
	}

	return validSources, nil
}

func (source *Source) GetLocaleID() string {
	if source.Params != nil {
		return source.Params.LocaleID
	}
	return ""
}

func (source *Source) setUploadParams(localeFile *LocaleFile) (*phraseapp.UploadParams, error) {
	uploadParams := new(phraseapp.UploadParams)
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

	fileEncoding := params.FileEncoding
	if fileEncoding != "" {
		uploadParams.FileEncoding = &fileEncoding
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

// Print out
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

// Reducer Algorithm:
// reducer := &Reducer{
// 		SourceFile: file,
// 		Extension: extension,
// }
// reducer.Initialize()
// reducer.Reduce()
// if reducer.MatchesPath(path) {
// 		reducer.Eval(path)
// }
type Reducer struct {
	SourceFile string
	Extension  string
	Tokens     []string
	Reductions []*Reduction
	Matcher    *regexp.Regexp
}

type Reduction struct {
	Original string
	Matcher  *regexp.Regexp
}

func (reducer *Reducer) Initialize() {
	tokens := []string{}
	for _, token := range strings.Split(reducer.SourceFile, "/") {
		if token == "." || token == "" {
			continue
		}
		tokens = append(tokens, token)
	}
	reducer.Tokens = tokens
}

func (reducer *Reducer) Reduce() error {
	reducer.Reductions = []*Reduction{}
	heads := []string{}
	for _, token := range reducer.Tokens {
		head, matcher, err := reducer.toRegexp(token)
		if err != nil {
			return err
		}
		heads = append(heads, head)
		reducer.Reductions = append(reducer.Reductions, &Reduction{
			Original: token,
			Matcher:  matcher,
		})
	}

	reMatcher, err := reducer.wholeMatcher(heads)
	if err != nil {
		return err
	}
	reducer.Matcher = reMatcher

	return nil
}

func (reducer *Reducer) MatchesPath(path string) bool {
	return reducer.Matcher.MatchString(path)
}

func (reducer *Reducer) Eval(path string) (*LocaleFile, error) {
	tagged := reducer.unify(path)

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

// Private Reducer Methods
func (reducer *Reducer) toRegexp(token string) (string, *regexp.Regexp, error) {
	// <locale_code>.yml => <locale_code>__PHRASEAPP_REGEXP_DOT__yml
	head := escapeRegexp(token)
	// <locale_code>__PHRASEAPP_REGEXP_DOT__yml => (?P<locale_code>.+)__PHRASEAPP_REGEXP_DOT__yml
	nextRegexp := reducer.convertToGroupRegexp(head)
	if nextRegexp != "" {
		head = nextRegexp
	}
	// (?P<locale_code>.+)__PHRASEAPP_REGEXP_DOT__yml => (?P<locale_code>.+)[.]yml
	head = convertRegexp(head)

	// Edge case: [.]strings => .*[.]strings
	if reducer.isLastToken(head) && strings.HasPrefix(head, "[.]") {
		head = ".*" + head
	}
	matcher, err := regexp.Compile(head)
	if err != nil {
		return "", nil, err
	}

	return head, matcher, nil
}

func (reducer *Reducer) isLastToken(head string) bool {
	return strings.Contains(head, strings.TrimLeft(reducer.Extension, "."))
}

func (reducer *Reducer) convertToGroupRegexp(part string) string {
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

func (reducer *Reducer) wholeMatcher(heads []string) (*regexp.Regexp, error) {
	matcherString := strings.Join(heads, "[/\\\\]")

	if strings.HasPrefix(matcherString, "[/\\\\]") {
		matcherString = strings.TrimLeft(matcherString, "[/\\\\]")
	}
	if strings.HasSuffix(matcherString, "[/\\\\]") {
		matcherString = strings.TrimRight(matcherString, "[/\\\\]")
	}

	reMatcher, err := regexp.Compile(matcherString)
	if err != nil {
		return nil, err
	}
	return reMatcher, nil
}

// ** can endlessly pump at any position
// initialization assumption: A Regexp validated the input path, the input is assumed to be correct.
// 		1. start leftmost until a ** occurs
// 			1.1 ./**/<tag>/<locale_name>/<locale_code>.yml -> nothing happens
// 			1.2 ./<locale_name>/<tag>/**/<locale_code>.yml -> unifiy <locale_name> and <tag>, break at **
// 			1.3 ./<locale_name>/**/<tag>/<locale_code>.yml -> unify <locale_name>, break at **
//		=> the left hand side of ** is fully unified
//
// 		2. start rightmost until the ** occurs
// 			2.1 ./<locale_name>/<tag>/**/<locale_code>.yml -> unifiy <locale_code>, break at **
// 			2.2 ./**/<tag>/<locale_name>/*.yml -> unify <locale_name> and <tag>, break at **
// 			2.3 ./<locale_name>/**/<tag>/<locale_code>.yml -> unify <tag> and <locale_code>
func (reducer *Reducer) unify(path string) map[string]string {
	tagged := map[string]string{}

	tokens := strings.Split(path, separator)
	reductions := reducer.Reductions

	if strings.Contains(reducer.SourceFile, "**") || reducer.fileContainsStar() {
		for idx, reduction := range reductions {
			if reduction.Original == "**" {
				break
			}
			if reduction.Original == "*" {
				continue
			}
			partlyTagged := reducer.tagMatches(reduction, tokens[idx])
			tagged = reducer.updateTaggedMatches(tagged, partlyTagged)
		}
	}

	offset := 0
	for i := len(reductions) - 1; i >= 0; i-- {
		reduction := reductions[i]
		if reduction.Original == "**" {
			break
		}
		if reduction.Original == "*" {
			offset += 1
			continue
		}
		idx := len(tokens) - offset - 1
		partlyTagged := reducer.tagMatches(reduction, tokens[idx])
		tagged = reducer.updateTaggedMatches(tagged, partlyTagged)
		offset += 1
	}

	return tagged
}

func (reducer *Reducer) fileContainsStar() bool {
	last := reducer.Reductions[len(reducer.Reductions)-1]
	return strings.Contains(reducer.SourceFile, "*") && !strings.Contains(last.Original, "*")
}

func (reducer *Reducer) tagMatches(reduction *Reduction, token string) map[string]string {
	tagged := map[string]string{}
	namedMatches := reduction.Matcher.SubexpNames()
	subMatches := reduction.Matcher.FindStringSubmatch(token)
	for i, subMatch := range subMatches {
		if subMatch != "" {
			tagged[namedMatches[i]] = subMatch
		}
	}
	return tagged
}

func (reducer *Reducer) updateTaggedMatches(original, updater map[string]string) map[string]string {
	localeCode := updater["locale_code"]
	localeName := updater["locale_name"]
	tag := updater["tag"]

	if original["locale_code"] == "" {
		original["locale_code"] = strings.Trim(localeCode, "/")
	}

	if original["locale_name"] == "" {
		original["locale_name"] = strings.Trim(localeName, "/")
	}

	if original["tag"] == "" {
		original["tag"] = strings.Trim(tag, "/")
	}
	return original
}

// Escaping of Regexp
func escapeRegexp(token string) string {
	for _, e := range regexpMapping {
		token = strings.Replace(token, e.Symbol, e.Escaper(), -1)
	}
	return token
}

func convertRegexp(token string) string {
	for _, e := range regexpMapping {
		token = strings.Replace(token, e.Escaper(), e.SymbolAsRegexp, -1)
	}
	return token
}

type RegexpEscaping struct {
	Symbol         string
	Replacer       string
	SymbolAsRegexp string
}

func (regexpEscaping *RegexpEscaping) Escaper() string {
	return fmt.Sprintf("__PHRASEAPP_REGEXP_%s__", regexpEscaping.Replacer)
}

var regexpMapping = []*RegexpEscaping{
	&RegexpEscaping{
		Symbol:         "**",
		Replacer:       "DOUBLE_STAR",
		SymbolAsRegexp: ".*",
	},
	&RegexpEscaping{
		Symbol:         "*",
		Replacer:       "SINGLE_STAR",
		SymbolAsRegexp: ".*",
	},
	&RegexpEscaping{
		Symbol:         "+",
		Replacer:       "PLUS",
		SymbolAsRegexp: "[+]",
	},
	&RegexpEscaping{
		Symbol:         "?",
		Replacer:       "QUESTION_MARK",
		SymbolAsRegexp: "[?]",
	},
	&RegexpEscaping{
		Symbol:         ".",
		Replacer:       "DOT",
		SymbolAsRegexp: "[.]",
	},
	&RegexpEscaping{
		Symbol:         "^",
		Replacer:       "CARET",
		SymbolAsRegexp: "\\^",
	},
	&RegexpEscaping{
		Symbol:         "(",
		Replacer:       "OPEN_BRACKET",
		SymbolAsRegexp: "[(]",
	},
	&RegexpEscaping{
		Symbol:         ")",
		Replacer:       "CLOSE_BRACKET",
		SymbolAsRegexp: "[)]",
	},
	&RegexpEscaping{
		Symbol:         "[",
		Replacer:       "OPEN_SQUARE_BRACKET",
		SymbolAsRegexp: "\\[",
	},
	&RegexpEscaping{
		Symbol:         "]",
		Replacer:       "CLOSE_SQUARE_BRACKET",
		SymbolAsRegexp: "\\]",
	},
	&RegexpEscaping{
		Symbol:         "{",
		Replacer:       "OPEN_CURLY_BRACKET",
		SymbolAsRegexp: "[{]",
	},
	&RegexpEscaping{
		Symbol:         "}",
		Replacer:       "CLOSE_CURLY_BRACKET",
		SymbolAsRegexp: "[}]",
	},
	&RegexpEscaping{
		Symbol:         "|",
		Replacer:       "PIPE",
		SymbolAsRegexp: "[|]",
	},
}
