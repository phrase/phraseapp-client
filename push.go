package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

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

	client, err := newClient(cmd.Config.Credentials)
	if err != nil {
		return err
	}

	sources, err := SourcesFromConfig(cmd)
	if err != nil {
		return err
	}

	if err := sources.Validate(); err != nil {
		return err
	}

	formatMap, err := GetFormats(client)
	if err != nil {
		return fmt.Errorf("Error retrieving format list from PhraseApp: %s", err)
	}

	for _, source := range sources {
		formatName := source.GetFileFormat()
		if val, ok := formatMap[formatName]; ok {
			source.Format = val
		}
		if source.Format == nil {
			return fmt.Errorf("Format %q of source %q is not supported by PhraseApp!", source.Format.ApiName, source.File)
		}
	}

	projectIdToLocales, err := LocalesForProjects(client, sources)
	if err != nil {
		return err
	}
	for _, source := range sources {
		val, ok := projectIdToLocales[source.ProjectID]
		if ok {
			source.RemoteLocales = val
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

func (source *Source) Push(client *phraseapp.Client) error {
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
