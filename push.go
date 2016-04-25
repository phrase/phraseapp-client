package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/phrase/phraseapp-client/Godeps/_workspace/src/gopkg.in/yaml.v2"

	"unicode/utf8"

	"github.com/phrase/phraseapp-client/Godeps/_workspace/src/github.com/phrase/phraseapp-go/phraseapp"
)

type PushCommand struct {
	*phraseapp.Config
}

func (cmd *PushCommand) Run() error {
	if cmd.Debug {
		// suppresses content output
		cmd.Debug = false
		Debug = true
	}

	client, err := newClient(cmd.Config.Credentials)
	if err != nil {
		return err
	}

	sources, err := SourcesFromConfig(cmd)
	if err != nil {
		return err
	}

	formats, err := client.FormatsList(1, 25)
	if err == nil {
		err = sources.setFormats(formats)
		if err != nil {
		}
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
	File        string
	ProjectID   string
	AccessToken string
	FileFormat  string
	Params      *phraseapp.UploadParams

	RemoteLocales []*phraseapp.Locale
	Format        *phraseapp.Format
}

func (src *Source) UnmarshalYAML(unmarshal func(interface{}) error) error {
	m := map[string]interface{}{}
	err := phraseapp.ParseYAMLToMap(unmarshal, map[string]interface{}{
		"file":         &src.File,
		"project_id":   &src.ProjectID,
		"access_token": &src.AccessToken,
		"file_format":  &src.FileFormat,
		"params":       &m,
	})
	if err != nil {
		return err
	}

	src.Params = new(phraseapp.UploadParams)
	return src.Params.ApplyValuesFromMap(m)
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

	// starCount contains the `**` so that must be taken into account.
	if starCount-(recCount*2) > 1 {
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

		if localeFile.shouldCreateLocale(source) {
			localeDetails, err := source.createLocale(client, localeFile)
			if err == nil {
				localeFile.ID = localeDetails.ID
				localeFile.Code = localeDetails.Code
				localeFile.Name = localeDetails.Name
			} else {
				fmt.Printf("failed to create locale: %s\n", err)
				continue
			}
		}

		err = source.uploadFile(client, localeFile)
		if err != nil {
			return err
		}

		sharedMessage("push", localeFile)

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
	} else if localeFile.Code != "" {
		localeParams.Name = &localeFile.Code
	}

	if localeFile.Code == "" {
		localeFile.Code = localeFile.Name
	}

	localeName := source.replacePlaceholderInParams(localeFile)
	if localeName != "" && localeName != localeFile.Code {
		localeParams.Name = &localeName
	}

	if localeFile.Code != "" {
		localeParams.Code = &localeFile.Code
	}

	localeDetails, err := client.LocaleCreate(source.ProjectID, localeParams)
	if err != nil {
		return nil, err
	}
	return localeDetails, nil
}

func (source *Source) replacePlaceholderInParams(localeFile *LocaleFile) string {
	if localeFile.Code != "" && strings.Contains(source.GetLocaleID(), "<locale_code>") {
		return strings.Replace(source.GetLocaleID(), "<locale_code>", localeFile.Code, 1)
	}
	return ""
}

func (source *Source) uploadFile(client *phraseapp.Client, localeFile *LocaleFile) error {
	if Debug {
		fmt.Fprintln(os.Stdout, "Source file pattern:", source.File)
		fmt.Fprintln(os.Stdout, "Actual file location:", localeFile.Path)
	}

	params := new(phraseapp.UploadParams)
	*params = *source.Params

	params.File = &localeFile.Path

	if params.LocaleID == nil {
		switch {
		case localeFile.ID != "":
			params.LocaleID = &localeFile.ID
		case localeFile.Code != "":
			params.LocaleID = &localeFile.Code
		}
	}

	if localeFile.Tag != "" {
		var v string
		if params.Tags != nil {
			v = *params.Tags + ","
		}
		v += localeFile.Tag
		params.Tags = &v
	}

	_, err := client.UploadCreate(source.ProjectID, params)
	return err
}

func (source *Source) SystemFiles() ([]string, error) {
	pattern := placeholderRegexp.ReplaceAllString(source.File, "*")
	parts := strings.SplitN(pattern, "**", 2)
	var pre, post string

	pre = parts[0]
	// strip trailing path separators
	for len(pre) > 0 && os.IsPathSeparator(pre[len(pre)-1]) {
		pre = pre[0 : len(pre)-1]
	}

	if len(parts) == 2 {
		post = parts[1]
		for len(post) > 0 && os.IsPathSeparator(post[0]) {
			post = post[1:]
		}
	}

	candidates, err := filepath.Glob(pre)
	if err != nil {
		return nil, err
	}

	prefix := "." + string(os.PathSeparator)
	if strings.HasPrefix(pre, prefix) {
		pre = strings.TrimLeft(pre, prefix)
	}

	var matches []string
	if post != "" {
		tokens := splitPathIntoSegments(strings.Replace(post, "*", ".*", -1))
		tokenCountPre := len(splitPathIntoSegments(pre))

		for _, cand := range candidates {
			if !isDir(cand) {
				continue
			}

			cands, err := findFilesInPath(cand)
			if err != nil {
				return nil, err
			}

			for _, c := range cands {
				if validateFileCandidate(tokens, tokenCountPre, c) {
					matches = append(matches, c)
				}
			}
		}
	} else {
		matches = candidates
	}

	return matches, nil
}

func isDir(path string) bool {
	stat, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return stat.IsDir()
}

func validateFileCandidate(tokens []string, ignoreTokenCnt int, cand string) bool {
	candTokens := splitPathIntoSegments(cand)
	candTokenCnt := len(candTokens)

	if candTokenCnt < (len(tokens) + ignoreTokenCnt) {
		return false
	}

	candTokens = candTokens[ignoreTokenCnt:]

	for i := 1; i <= len(tokens); i++ {
		expT, gotT := tokens[len(tokens)-i], candTokens[len(candTokens)-i]
		switch {
		case strings.Contains(expT, "*"):
			matched, err := regexp.MatchString(expT, gotT)
			if err != nil {
				panic(err)
			}
			if !matched {
				return false
			}
		case expT != gotT:
			return false
		}
	}

	return true
}

func splitPathIntoSegments(path string) []string {
	segments := []string{}
	start := 0
	for i := range path {
		if os.IsPathSeparator(path[i]) {
			segments = append(segments, path[start:i])
			start = i + 1
		}
	}
	return append(segments, path[start:])
}

func findFilesInPath(root string) ([]string, error) {
	files := []string{}
	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !f.Mode().IsDir() {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

// Return all locale files from disk that match the source pattern.
func (source *Source) LocaleFiles() (LocaleFiles, error) {

	filePaths, err := source.SystemFiles()
	if err != nil {
		return nil, err
	}

	tokens := splitPathToTokens(source.File)

	var localeFiles LocaleFiles
	for _, path := range filePaths {
		pathTokens := splitPathToTokens(path)
		localeFile := extractParamsFromPathTokens(tokens, pathTokens)

		absolutePath, err := filepath.Abs(path)
		if err != nil {
			return nil, err
		}

		localeFile.Path = absolutePath

		locale := source.getRemoteLocaleForLocaleFile(localeFile)
		if locale != nil {
			localeFile.ExistsRemote = true
			localeFile.Code = locale.Code
			localeFile.Name = locale.Name
			localeFile.ID = locale.ID
		}

		if Debug {
			fmt.Printf(
				"Code:%q, Name:%q, ID:%q, Tag:%q\n",
				localeFile.Code, localeFile.Name, localeFile.ID, localeFile.Tag,
			)
		}

		localeFiles = append(localeFiles, localeFile)
	}

	if len(localeFiles) <= 0 {
		abs, err := filepath.Abs(source.File)
		if err != nil {
			abs = source.File
		}
		return nil, fmt.Errorf("Could not find any files on your system that matches: '%s'", abs)
	}
	return localeFiles, nil
}

func (source *Source) getRemoteLocaleForLocaleFile(localeFile *LocaleFile) *phraseapp.Locale {
	candidates := source.RemoteLocales

	filterApplied := false

	filter := func(cands []*phraseapp.Locale, preCond string, pred func(cand *phraseapp.Locale) bool) []*phraseapp.Locale {
		if preCond == "" {
			return cands
		}
		filterApplied = true
		tmpCands := []*phraseapp.Locale{}
		for _, cand := range cands {
			if pred(cand) {
				tmpCands = append(tmpCands, cand)
			}
		}
		return tmpCands
	}

	localeName := source.replacePlaceholderInParams(localeFile)
	if localeName != "" {
		// This means the name can contain the value specified in LocaleID, with
		// `<locale_code>` being substituted by the value of the currently handled
		// localeFile (like push only locales with name `en-US`).
		candidates = filter(candidates, localeName, func(cand *phraseapp.Locale) bool {
			return strings.Contains(cand.Name, localeName)
		})
	} else {
		localeID := source.GetLocaleID()
		candidates = filter(candidates, localeID, func(cand *phraseapp.Locale) bool {
			return cand.Name == localeID || cand.ID == localeID
		})
	}

	candidates = filter(candidates, localeFile.Name, func(cand *phraseapp.Locale) bool {
		return cand.Name == localeFile.Name
	})

	candidates = filter(candidates, localeFile.Code, func(cand *phraseapp.Locale) bool {
		return cand.Code == localeFile.Code
	})

	// If no filter was applied the candidates list still contains all remote
	// locales, while actually nothing matches.
	if !filterApplied {
		return nil
	}

	switch len(candidates) {
	case 0:
		return nil
	case 1:
		return candidates[0]
	default:
		// TODO I guess this should return an error, as this is a problem.
		return candidates[0]
	}
}

func splitString(s string, set string) []string {
	if len(set) == 1 {
		return strings.Split(s, set)
	}

	slist := []string{}
	charSet := map[rune]bool{}

	for _, r := range set {
		charSet[r] = true
	}

	start := 0
	for i, r := range s {
		if _, found := charSet[r]; found {
			slist = append(slist, s[start:i])
			start = i + utf8.RuneLen(r)
		}
	}
	if start < len(s) {
		slist = append(slist, s[start:])
	}

	return slist
}

func splitPathToTokens(s string) []string {
	tokens := []string{}
	splitSet := separator
	if separator == "\\" {
		splitSet = "\\/"
	}
	for _, token := range splitString(s, splitSet) {
		if token == "." || token == "" {
			continue
		}
		tokens = append(tokens, token)
	}
	return tokens
}

func extractParamsFromPathTokens(srcTokens, pathTokens []string) *LocaleFile {
	localeFile := new(LocaleFile)

	for idx, token := range srcTokens {
		pathToken := pathTokens[idx]
		if token == "*" {
			continue
		}
		if token == "**" {
			break
		}
		extractParamFromPathToken(localeFile, token, pathToken)
	}

	if Contains(srcTokens, "**") {
		offset := 1
		for idx := len(srcTokens) - 1; idx >= 0; idx-- {
			token := srcTokens[idx]
			pathToken := pathTokens[len(pathTokens)-offset]
			offset += 1

			if token == "*" {
				continue
			}
			if token == "**" {
				break
			}

			extractParamFromPathToken(localeFile, token, pathToken)
		}
	}

	return localeFile
}

func extractParamFromPathToken(localeFile *LocaleFile, srcToken, pathToken string) {
	groups := placeholderRegexp.FindAllString(srcToken, -1)
	if len(groups) <= 0 {
		return
	}

	match := strings.Replace(srcToken, ".", "[.]", -1)
	if strings.HasPrefix(match, "*") {
		match = strings.Replace(match, "*", ".*", -1)
	}

	for _, group := range groups {
		replacer := fmt.Sprintf("(?P%s.+)", group)
		match = strings.Replace(match, group, replacer, 1)
	}

	if match == "" {
		return
	}

	tmpRegexp, err := regexp.Compile(match)
	if err != nil {
		return
	}

	namedMatches := tmpRegexp.SubexpNames()
	subMatches := tmpRegexp.FindStringSubmatch(pathToken)
	for i, subMatch := range subMatches {
		value := strings.Trim(subMatch, separator)
		switch namedMatches[i] {
		case "locale_code":
			localeFile.Code = value
		case "locale_name":
			localeFile.Name = value
		case "tag":
			localeFile.Tag = value
		default:
			// ignore
		}
	}
}

func SourcesFromConfig(cmd *PushCommand) (Sources, error) {
	if cmd.Config.Sources == nil || len(cmd.Config.Sources) == 0 {
		return nil, fmt.Errorf("no sources for upload specified")
	}

	tmp := struct {
		Sources Sources
	}{}
	err := yaml.Unmarshal(cmd.Config.Sources, &tmp)
	if err != nil {
		return nil, err
	}
	srcs := tmp.Sources

	token := cmd.Credentials.Token
	projectId := cmd.Config.DefaultProjectID
	fileFormat := cmd.Config.DefaultFileFormat

	validSources := []*Source{}
	for _, source := range srcs {
		if source == nil {
			continue
		}
		if source.ProjectID == "" {
			source.ProjectID = projectId
		}
		if source.AccessToken == "" {
			source.AccessToken = token
		}
		if source.Params == nil {
			source.Params = new(phraseapp.UploadParams)
		}

		if source.Params.FileFormat == nil {
			switch {
			case source.FileFormat != "":
				source.Params.FileFormat = &source.FileFormat
			case fileFormat != "":
				source.Params.FileFormat = &fileFormat
			}
		}
		validSources = append(validSources, source)
	}

	if len(validSources) <= 0 {
		return nil, fmt.Errorf("no sources could be identified! Refine the sources list in your config")
	}

	return validSources, nil
}

func (source *Source) GetLocaleID() string {
	if source.Params != nil && source.Params.LocaleID != nil {
		return *source.Params.LocaleID
	}
	return ""
}

func (source *Source) GetFileFormat() string {
	if source.Params != nil && source.Params.FileFormat != nil {
		return *source.Params.FileFormat
	}
	if source.FileFormat != "" {
		return source.FileFormat
	}
	return ""
}

func (sources Sources) setFormats(formats []*phraseapp.Format) error {
	formatMap := map[string]*phraseapp.Format{}
	for _, format := range formats {
		formatMap[format.ApiName] = format
	}

	for _, source := range sources {
		formatName := source.GetFileFormat()
		if val, ok := formatMap[formatName]; ok {
			source.Format = val
		}
	}

	return nil
}

func (localeFile *LocaleFile) shouldCreateLocale(source *Source) bool {
	if localeFile.ExistsRemote {
		return false
	}

	if source.Format.IncludesLocaleInformation {
		return false
	}

	// we could not find an existing locale in PhraseApp
	// if a locale_name or locale_code was provided by the placeholder logic
	// we assume that it should be created
	// every other source should be uploaded and validated in uploads#create
	return (localeFile.Name != "" || localeFile.Code != "")
}
