package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"

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

	err := func() error {
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
	}()

	if err != nil {
		ReportError("Push Error", err.Error())
	}

	return err
}

type Sources []*Source

type Source struct {
	File          string      `yaml:"file,omitempty"`
	ProjectID     string      `yaml:"project_id,omitempty"`
	AccessToken   string      `yaml:"access_token,omitempty"`
	FileFormat    string      `yaml:"file_format,omitempty"`
	Params        *phraseapp.UploadParams `yaml:"params"`

	RemoteLocales []*phraseapp.Locale
	Extension     string
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
	if strings.Contains(source.File, "**") {
		return source.recurse()
	}

	return source.glob()
}

func (source *Source) glob() ([]string, error) {
	pattern := placeholderRegexp.ReplaceAllString(source.File, "*")
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
		if !f.Mode().IsDir() && strings.HasSuffix(f.Name(), source.Extension) {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

func (source *Source) root() string {
	parts := splitPathToTokens(source.File)
	rootParts := TakeWhile(parts, func(x string) bool {
		return x != "**"
	})
	root := strings.Join(rootParts, separator)
	if root == "" {
		root = "."
	}
	return root
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
			fmt.Println(fmt.Sprintf(
				"RFC:'%s', Name:'%s', Tag;'%s', Pattern:'%s'",
				localeFile.RFC, localeFile.Name, localeFile.Tag,
			))
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

func splitPathToTokens(s string) []string {
	tokens := []string{}
	for _, token := range strings.Split(s, separator) {
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

// Configuration
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
