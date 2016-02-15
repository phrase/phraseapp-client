package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"

	"unicode/utf8"

	"github.com/phrase/phraseapp-go/phraseapp"
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

	err := func() error {
		client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	}()

	if err != nil {
		ReportError("Push Error", err.Error())
	}

	return err
}

type Sources []*Source

type Source struct {
	File        string
	ProjectID   string
	AccessToken string
	FileFormat  string
	Params      *phraseapp.UploadParams

	RemoteLocales []*phraseapp.Locale
	Extension     string
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
		}

		sharedMessage("push", localeFile)

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
		case localeFile.RFC != "":
			params.LocaleID = &localeFile.RFC
		}
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
	for len(pre) > 0 && os.IsPathSeparator(pre[len(pre) - 1]) {
		pre = pre[0:len(pre) - 1]
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
		expT, gotT := tokens[len(tokens) - i], candTokens[len(candTokens) - i]
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
		if len(pathTokens) < len(tokens) {
			continue
		}
		localeFile := extractParamsFromPathTokens(tokens, pathTokens)

		absolutePath, err := filepath.Abs(path)
		if err != nil {
			return nil, err
		}
		localeFile.Path = absolutePath

		locale := source.getRemoteLocaleForLocaleFile(localeFile)
		if locale != nil {
			localeFile.ExistsRemote = true
			localeFile.RFC = locale.Code
			localeFile.Name = locale.Name
			localeFile.ID = locale.ID
		}

		if Debug {
			fmt.Printf(
				"RFC:%q, Name:%q, ID:%q, Tag:%q\n",
				localeFile.RFC, localeFile.Name, localeFile.ID, localeFile.Tag,
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

		if remote.Name == localeFile.RFC {
			return remote
		}
	}
	return nil
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
			localeFile.RFC = value
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
		errmsg := "no sources for upload specified"
		ReportError("Push Error", errmsg)
		return nil, fmt.Errorf(errmsg)
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
