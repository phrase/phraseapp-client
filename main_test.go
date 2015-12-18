package main

import (
	"testing"

	"fmt"
	"github.com/phrase/phraseapp-go/phraseapp"
	"gopkg.in/yaml.v1"
	"sort"
)

func TestPullConfig(t *testing.T) {
	pullConfig := `
phraseapp:
  pull:
    targets:
      -
        file: ./locales/file.yml
        params:
          file_format: strings
`
	config := &PullConfig{}
	err := yaml.Unmarshal([]byte(pullConfig), &config)
	if err != nil {
		t.Errorf(err.Error())
	}

	targetParams := config.Phraseapp.Pull.Targets[0].Params
	if targetParams.FileFormat != "strings" {
		t.Errorf("Expected FileFormat of first target to be %s and not %s", "strings", targetParams.FileFormat)
	}
}

func TestPushConfig(t *testing.T) {
	pushConfig := `
phraseapp:
  push:
    sources:
      -
        file: ./locales/file.yml
        params:
          file_format: strings
`
	config := &PushConfig{}
	err := yaml.Unmarshal([]byte(pushConfig), &config)
	if err != nil {
		t.Errorf(err.Error())
	}

	sourceParams := config.Phraseapp.Push.Sources[0].Params
	if sourceParams.FileFormat == nil || *sourceParams.FileFormat != "strings" {
		t.Errorf("Expected FileFormat of first target to be %s and not %s", "strings", sourceParams.FileFormat)
	}
}

func getBaseLocales() []*phraseapp.Locale {
	return []*phraseapp.Locale{
		&phraseapp.Locale{
			Code: "en",
			ID:   "en-locale-id",
			Name: "english",
		},
		&phraseapp.Locale{
			Code: "de",
			ID:   "de-locale-id",
			Name: "german",
		},
	}
}

type ByPath []*LocaleFile

func (a ByPath) Len() int           { return len(a) }
func (a ByPath) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPath) Less(i, j int) bool { return a[i].Path < a[j].Path }

func compareLocaleFiles(actualFiles LocaleFiles, expectedFiles LocaleFiles) error {
	sort.Sort(ByPath(actualFiles))
	sort.Sort(ByPath(expectedFiles))
	for idx, actualFile := range actualFiles {
		expected := expectedFiles[idx]
		actual := actualFile
		if expected.Path != actual.Path {
			return fmt.Errorf("Expected Path %s should eql %s", expected.Path, actual.Path)
		}
		if expected.Name != actual.Name {
			return fmt.Errorf("Expected Name %s should eql %s", expected.Name, actual.Name)
		}
		if expected.RFC != actual.RFC {
			return fmt.Errorf("Expected RFC %s should eql %s", expected.RFC, actual.RFC)
		}
		if expected.ID != actual.ID {
			return fmt.Errorf("Expected ID %s should eql %s", expected.ID, actual.ID)
		}
	}
	return nil
}
