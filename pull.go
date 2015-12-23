package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"strings"

	"github.com/phrase/phraseapp-go/phraseapp"
)

type PullCommand struct {
	Credentials
}

func (cmd *PullCommand) Run() error {
	if cmd.Debug {
		// suppresses content output
		cmd.Debug = false
		Debug = true
	}

	client, err := ClientFromCmdCredentials(cmd.Credentials)
	if err != nil {
		return err
	}

	targets, err := TargetsFromConfig(cmd)
	if err != nil {
		return err
	}

	for _, target := range targets {
		err := target.Pull(client)
		if err != nil {
			return err
		}
	}
	return nil
}

type Targets []*Target

type Target struct {
	File          string      `yaml:"file,omitempty"`
	ProjectID     string      `yaml:"project_id,omitempty"`
	AccessToken   string      `yaml:"access_token,omitempty"`
	FileFormat    string      `yaml:"file_format,omitempty"`
	Params        *PullParams `yaml:"params"`
	RemoteLocales []*phraseapp.Locale
}

type PullParams struct {
	phraseapp.LocaleDownloadParams
	LocaleID string
}

func (params *PullParams) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var aux struct {
		LocaleID string `yaml:"locale_id,omitempty"`
	}
	if err := unmarshal(&aux); err != nil {
		return err
	}
	params.LocaleID = aux.LocaleID

	return unmarshal(&params.LocaleDownloadParams)
}


func (target *Target) CheckPreconditions() error {
	if err := ValidPath(target.File, target.FileFormat, ""); err != nil {
		return err
	}

	if strings.Count(target.File, "*") > 0 {
		return fmt.Errorf(
			"File pattern for 'pull' cannot include any 'stars' *. Please specify direct and valid paths with file name!",
			"http://docs.phraseapp.com/developers/cli/configuration/#targets",
		)
	}

	duplicatedPlaceholders := []string{}
	for _, name := range []string{"<locale_name>", "<locale_code>", "<tag>"} {
		if strings.Count(target.File, name) > 1 {
			duplicatedPlaceholders = append(duplicatedPlaceholders, name)
		}
	}

	if len(duplicatedPlaceholders) > 0 {
		dups := strings.Join(duplicatedPlaceholders, ", ")
		return fmt.Errorf(fmt.Sprintf("%s can only occur once in a file pattern!", dups))
	}

	return nil
}

func (target *Target) Pull(client *phraseapp.Client) error {

	if err := target.CheckPreconditions(); err != nil {
		return err
	}

	remoteLocales, err := RemoteLocales(client, target.ProjectID)
	if err != nil {
		return err
	}
	target.RemoteLocales = remoteLocales

	localeFiles, err := target.LocaleFiles()
	if err != nil {
		return err
	}

	localeIdToFileIsDistinct := (target.GetLocaleID() != "" && len(localeFiles) == 1)

	for _, localeFile := range localeFiles {
		err := createFile(localeFile.Path)
		if err != nil {
			return err
		}

		if localeIdToFileIsDistinct {
			if target.GetLocaleID() != "" {
				localeFile.ID = target.GetLocaleID()
			}
		}

		err = target.DownloadAndWriteToFile(client, localeFile)
		if err != nil {
			errmsg := fmt.Sprintf("%s for %s", err, localeFile.Path)
			ReportError("Pull Error", errmsg)
			return fmt.Errorf(errmsg)
		} else {
			sharedMessage("pull", localeFile)
		}
		if Debug {
			fmt.Fprintln(os.Stderr, strings.Repeat("-", 10))
		}
	}

	return nil
}

func (target *Target) DownloadAndWriteToFile(client *phraseapp.Client, localeFile *LocaleFile) error {
	downloadParams := new(phraseapp.LocaleDownloadParams)
	*downloadParams = target.Params.LocaleDownloadParams

	if downloadParams.FileFormat == nil {
		downloadParams.FileFormat = &localeFile.FileFormat
	}

	if Debug {
		fmt.Fprintln(os.Stderr, "Target file pattern:", target.File)
		fmt.Fprintln(os.Stderr, "Actual file path", localeFile.Path)
		fmt.Fprintln(os.Stderr, "LocaleID", localeFile.ID)
		fmt.Fprintln(os.Stderr, "ProjectID", target.ProjectID)
		fmt.Fprintln(os.Stderr, "FileFormat", downloadParams.FileFormat)
		fmt.Fprintln(os.Stderr, "ConvertEmoji", downloadParams.ConvertEmoji)
		fmt.Fprintln(os.Stderr, "IncludeEmptyTranslations", downloadParams.IncludeEmptyTranslations)
		fmt.Fprintln(os.Stderr, "KeepNotranslateTags", downloadParams.KeepNotranslateTags)
		fmt.Fprintln(os.Stderr, "Tag", downloadParams.Tag)
		fmt.Fprintln(os.Stderr, "FormatOptions", downloadParams.FormatOptions)
	}

	res, err := client.LocaleDownload(target.ProjectID, localeFile.ID, downloadParams)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(localeFile.Path, res, 0700)
	if err != nil {
		return err
	}
	return nil
}

func (target *Target) LocaleFiles() (LocaleFiles, error) {
	localeID := target.GetLocaleID()
	files := []*LocaleFile{}
	for _, remoteLocale := range target.RemoteLocales {
		if localeID != "" && !(remoteLocale.ID == localeID || remoteLocale.Name == localeID) {
			continue
		}
		err := target.IsValidLocale(remoteLocale, target.File)
		if err != nil {
			return nil, err
		}

		localeFile := &LocaleFile{
			Name:       remoteLocale.Name,
			ID:         remoteLocale.ID,
			RFC:        remoteLocale.Code,
			Tag:        target.GetTag(),
			FileFormat: target.GetFormat(),
			Path:       target.File,
		}

		absPath, err := target.ReplacePlaceholders(localeFile)
		if err != nil {
			return nil, err
		}
		localeFile.Path = absPath

		files = append(files, localeFile)
	}

	return files, nil
}

func (target *Target) IsValidLocale(locale *phraseapp.Locale, localPath string) error {
	if locale == nil {
		errmsg := "Remote locale could not be downloaded correctly!"
		ReportError("Pull Error", errmsg)
		return fmt.Errorf(errmsg)
	}

	if strings.Contains(localPath, "<locale_code>") && locale.Code == "" {
		errmsg := fmt.Sprintf("Locale code is not set for Locale with ID: %s but locale_code is used in file name", locale.ID)
		ReportError("Pull Error", errmsg)
		return fmt.Errorf(errmsg)
	}
	return nil
}

func (target *Target) ReplacePlaceholders(localeFile *LocaleFile) (string, error) {
	absPath, err := filepath.Abs(target.File)
	if err != nil {
		return "", err
	}

	path := strings.Replace(absPath, "<locale_name>", localeFile.Name, -1)
	path = strings.Replace(path, "<locale_code>", localeFile.RFC, -1)
	path = strings.Replace(path, "<tag>", localeFile.Tag, -1)

	return path, nil
}

func (t *Target) GetFormat() string {
	if t.Params != nil && t.Params.FileFormat != nil {
		return *t.Params.FileFormat
	}
	if t.FileFormat != "" {
		return t.FileFormat
	}
	return ""
}

func (t *Target) GetLocaleID() string {
	if t.Params != nil {
		return t.Params.LocaleID
	}
	return ""
}

func (t *Target) GetTag() string {
	if t.Params != nil && t.Params.Tag != nil {
		return *t.Params.Tag
	}
	return ""
}

// Configuration
type PullConfig struct {
	Phraseapp struct {
		AccessToken string `yaml:"access_token"`
		ProjectID   string `yaml:"project_id"`
		FileFormat  string `yaml:"file_format,omitempty"`
		Pull        struct {
			Targets Targets
		}
	}
}

func TargetsFromConfig(cmd *PullCommand) (Targets, error) {
	content, err := ConfigContent()
	if err != nil {
		return nil, err
	}

	var config *PullConfig

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

	if &config.Phraseapp.Pull == nil || config.Phraseapp.Pull.Targets == nil {
		errmsg := "no targets for download specified"
		ReportError("Pull Error", errmsg)
		return nil, fmt.Errorf(errmsg)
	}

	targets := config.Phraseapp.Pull.Targets

	validTargets := []*Target{}
	for _, target := range targets {
		if target == nil {
			continue
		}
		if target.ProjectID == "" {
			target.ProjectID = projectId
		}
		if target.AccessToken == "" {
			target.AccessToken = token
		}
		if target.FileFormat == "" {
			target.FileFormat = fileFormat
		}
		validTargets = append(validTargets, target)
	}

	if len(validTargets) <= 0 {
		errmsg := "no targets could be identified! Refine the targets list in your config"
		ReportError("Pull Error", errmsg)
		return nil, fmt.Errorf(errmsg)
	}

	return validTargets, nil
}

func createFile(path string) error {
	err := Exists(path)
	if err != nil {
		absDir := filepath.Dir(path)
		err := Exists(absDir)
		if err != nil {
			os.MkdirAll(absDir, 0700)
		}

		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
	}
	return nil
}
