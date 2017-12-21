package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/phrase/phraseapp-client/internal/paths"
	"github.com/phrase/phraseapp-go/phraseapp"
	yaml "gopkg.in/yaml.v2"
)

func SourcesFromConfig(config phraseapp.Config) (Sources, error) {
	if config.Sources == nil || len(config.Sources) == 0 {
		return nil, fmt.Errorf("no sources for upload specified")
	}

	tmp := struct {
		Sources Sources
	}{}
	err := yaml.Unmarshal(config.Sources, &tmp)
	if err != nil {
		return nil, err
	}
	srcs := tmp.Sources

	projectId := config.DefaultProjectID
	fileFormat := config.DefaultFileFormat

	validSources := []*Source{}
	for _, source := range srcs {
		if source == nil {
			continue
		}
		if source.ProjectID == "" {
			source.ProjectID = projectId
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

type Sources []*Source

func (sources Sources) Validate() error {
	for _, source := range sources {
		if err := source.CheckPreconditions(); err != nil {
			return err
		}
	}
	return nil
}

type Source struct {
	File        string
	ProjectID   string
	Branch      string
	AccessToken string
	FileFormat  string
	Params      *phraseapp.UploadParams

	RemoteLocales []*phraseapp.Locale
	Format        *phraseapp.Format
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

func (source *Source) CheckPreconditions() error {
	if err := paths.Validate(source.File, source.FileFormat, ""); err != nil {
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

func (sources Sources) ProjectIds() []string {
	projectIds := []string{}
	for _, source := range sources {
		projectIds = append(projectIds, source.ProjectID)
	}
	return projectIds
}
func (source *Source) uploadFile(client *phraseapp.Client, localeFile *LocaleFile, branch string) (*phraseapp.Upload, error) {
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

	if branch != "" {
		params.Branch = &branch
	}

	return client.UploadCreate(source.ProjectID, params)
}

func (source *Source) createLocale(client *phraseapp.Client, localeFile *LocaleFile, branch string) (*phraseapp.LocaleDetails, error) {
	localeDetails, found, err := source.getLocaleIfExist(client, localeFile)
	if err != nil {
		return nil, err
	} else if found {
		return localeDetails, nil
	}

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
	// localeParams.Branch = branch

	localeDetails, err = client.LocaleCreate(source.ProjectID, localeParams)
	if err != nil {
		return nil, err
	}
	return localeDetails, nil
}

func (source *Source) getLocaleIfExist(client *phraseapp.Client, localeFile *LocaleFile) (*phraseapp.LocaleDetails, bool, error) {
	identifier := localeIdentifier(source, localeFile)
	if identifier == "" {
		return nil, false, nil
	}

	localeDetail, err := client.LocaleShow(source.ProjectID, identifier)
	if phraseapp.IsErrNotFound(err) {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	return localeDetail, true, nil
}

func localeIdentifier(source *Source, localeFile *LocaleFile) string {
	localeName := source.replacePlaceholderInParams(localeFile)
	if localeName != "" && localeName != localeFile.Code {
		return localeName
	}

	if localeFile.Name != "" {
		return localeFile.Name
	}

	if localeFile.Code != "" {
		return localeFile.Code
	}

	return ""
}

func (source *Source) replacePlaceholderInParams(localeFile *LocaleFile) string {
	if localeFile.Code != "" && strings.Contains(source.GetLocaleID(), "<locale_code>") {
		return strings.Replace(source.GetLocaleID(), "<locale_code>", localeFile.Code, 1)
	}
	return ""
}
