package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/phrase/phraseapp-client/internal/paths"
	"github.com/phrase/phraseapp-client/internal/placeholders"
	"github.com/phrase/phraseapp-client/internal/shared"
	"github.com/phrase/phraseapp-go/phraseapp"
	"gopkg.in/yaml.v2"
)

type Targets []*Target

func (targets Targets) ProjectIds() []string {
	projectIds := []string{}
	for _, target := range targets {
		projectIds = append(projectIds, target.ProjectID)
	}
	return projectIds
}

type Target struct {
	File          string
	ProjectID     string
	AccessToken   string
	FileFormat    string
	Params        *PullParams
	RemoteLocales []*phraseapp.Locale
}

func (target *Target) CheckPreconditions() error {
	if err := paths.Validate(target.File, target.FileFormat, ""); err != nil {
		return err
	}

	preconditions := []func(*Target) error{
		containsStars,
		containsDuplicatePlaceholders,
		containsAmbiguousLocaleInformation,
		containsInvalidTagInformation,
	}

	for _, precondition := range preconditions {
		if err := precondition(target); err != nil {
			return err
		}
	}

	return nil
}

func containsStars(target *Target) error {
	if strings.Contains(target.File, "*") {
		return fmt.Errorf("File pattern for 'pull' cannot include any 'stars' *. Please specify direct and valid paths with file name!\n %s#targets", shared.DocsConfigUrl)
	}
	return nil
}

func containsDuplicatePlaceholders(target *Target) error {
	duplicatedPlaceholders := []string{}
	for _, name := range []string{"<locale_name>", "<locale_code>", "<tag>"} {
		if strings.Count(target.File, name) > 1 {
			duplicatedPlaceholders = append(duplicatedPlaceholders, name)
		}
	}

	if len(duplicatedPlaceholders) > 0 {
		dups := strings.Join(duplicatedPlaceholders, ", ")
		return fmt.Errorf("The following placeholders occur more than once (no duplicates allowed):\n%s", dups)
	}

	return nil
}

func containsAmbiguousLocaleInformation(target *Target) error {
	if target.GetLocaleID() == "" && !placeholders.ContainsLocalePlaceholder(target.File) {
		// need more locale information
		return fmt.Errorf("Could not find any locale information. Please specify a 'locale_id' in your params or provide a placeholder (<locale_code|locale_name>)")
	} else if target.GetLocaleID() != "" && placeholders.ContainsLocalePlaceholder(target.File) {
		// ambiguous (too many information)
		return fmt.Errorf("Found 'locale_id' in params and a (<locale_code|locale_name>) placeholder. Please only select one per file pattern.")
	}

	return nil
}

func containsInvalidTagInformation(target *Target) error {
	if len(target.GetTags()) == 0 && placeholders.ContainsTagPlaceholder(target.File) {
		// tag provided but no params
		return fmt.Errorf("Using <tag> placeholder but no tags were provided. Please specify 'tags: \"my_tag\"' in the params section.")
	}
	return nil
}

//
func (target *Target) localeForRemote() (*phraseapp.Locale, error) {
	for _, locale := range target.RemoteLocales {
		if locale.ID == target.GetLocaleID() || locale.Name == target.GetLocaleID() {
			return locale, nil
		}
	}
	return nil, fmt.Errorf("Provided locale_id %q but did not match any remote locales in project %q", target.GetLocaleID(), target.ProjectID)
}

func (target *Target) ReplacePlaceholders(localeFile *LocaleFile) (string, error) {
	absPath, err := filepath.Abs(target.File)
	if err != nil {
		return "", err
	}

	path := strings.Replace(absPath, "<locale_name>", localeFile.Name, -1)
	path = strings.Replace(path, "<locale_code>", localeFile.Code, -1)
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

func (t *Target) GetTags() []string {
	tagList := []string{}
	var tagsParam string
	if t.Params != nil && t.Params.Tags != nil {
		tagsParam = *t.Params.Tags
	} else if t.Params != nil && t.Params.Tag != nil {
		tagsParam = *t.Params.Tag
	}

	if tagsParam != "" {
		tagsParam = strings.Replace(tagsParam, " ", "", -1)
		tagList = strings.Split(tagsParam, ",")
	}

	return tagList
}

func TargetsFromConfig(config phraseapp.Config) (Targets, error) {
	if config.Targets == nil || len(config.Targets) == 0 {
		return nil, fmt.Errorf("no targets for download specified")
	}

	tmp := struct {
		Targets Targets
	}{}
	err := yaml.Unmarshal(config.Targets, &tmp)
	if err != nil {
		return nil, err
	}
	tgts := tmp.Targets

	projectId := config.DefaultProjectID
	fileFormat := config.DefaultFileFormat

	validTargets := []*Target{}
	for _, target := range tgts {
		if target == nil {
			continue
		}
		if target.ProjectID == "" {
			target.ProjectID = projectId
		}
		if target.FileFormat == "" {
			target.FileFormat = fileFormat
		}
		validTargets = append(validTargets, target)
	}

	if len(validTargets) <= 0 {
		return nil, fmt.Errorf("no targets could be identified! Refine the targets list in your config")
	}

	return validTargets, nil
}

func (tgt *Target) UnmarshalYAML(unmarshal func(interface{}) error) error {
	m := map[string]interface{}{}
	err := phraseapp.ParseYAMLToMap(unmarshal, map[string]interface{}{
		"file":         &tgt.File,
		"project_id":   &tgt.ProjectID,
		"access_token": &tgt.AccessToken,
		"file_format":  &tgt.FileFormat,
		"params":       &m,
	})
	if err != nil {
		return err
	}

	tgt.Params = new(PullParams)
	if v, found := m["locale_id"]; found {
		if tgt.Params.LocaleID, err = phraseapp.ValidateIsString("params.locale_id", v); err != nil {
			return err
		}
		// Must delete the param from the map as the LocaleDownloadParams type
		// doesn't support this one and the apply method would return an error.
		delete(m, "locale_id")
	}

	return tgt.Params.ApplyValuesFromMap(m)
}
