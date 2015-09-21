package main

import (
	"fmt"
	"github.com/phrase/phraseapp-go/phraseapp"
	"path/filepath"
	"sort"
	"testing"
)

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
func getBaseSource() *Source {
	source := &Source{
		File:        "./tests/<locale_code>.yml",
		ProjectID:   "project-id",
		AccessToken: "access-token",
		FileFormat:  "yml",
		Extension:   "",
		Params: &PushParams{
			FileFormat:         "",
			LocaleID:           "",
			ConvertEmoji:       nil,
			SkipUnverification: nil,
			SkipUploadTags:     nil,
			Tags:               nil,
			UpdateTranslations: nil,
		},
		RemoteLocales: getBaseLocales(),
	}
	source.Extension = filepath.Ext(source.File)
	return source
}

func TestSourceFields(t *testing.T) {
	source := getBaseSource()

	if source.File != "./tests/<locale_code>.yml" {
		t.Errorf("Expected File to be %s and not %s", "./tests/<locale_code>.yml", source.File)
	}

	if source.AccessToken != "access-token" {
		t.Errorf("Expected AccesToken to be %s and not %s", "access-token", source.AccessToken)
	}

	if source.ProjectID != "project-id" {
		t.Errorf("Expected ProjectID to be %s and not %s", "project-id", source.ProjectID)
	}

	if source.FileFormat != "yml" {
		t.Errorf("Expected FileFormat to be %s and not %s", "yml", source.FileFormat)
	}

}

func TestSourceCheckPreconditions(t *testing.T) {
	source := getBaseSource()

	for _, file := range []string{
		"./<locale_code>/<locale_code>.yml",
		"./**/**/en.yml",
		"./**/*/*/en.yml",
	} {
		source.File = file
		if err := source.CheckPreconditions(); err == nil {
			t.Errorf("CheckPrecondition did not fail!")
		}
	}

	for _, file := range []string{
		"./<tag>/<locale_code>.yml",
		"./**/*/en.yml",
		"./**/*/<locale_name>/<locale_code>/<tag>.yml",
	} {
		source.File = file
		if err := source.CheckPreconditions(); err != nil {
			t.Errorf("CheckPrecondition should not fail with: %s", err.Error())
		}
	}
}

func TestSourceLocaleFilesOne(t *testing.T) {
	source := getBaseSource()
	localeFiles, err := source.LocaleFiles()

	if err != nil {
		t.Errorf("Should not fail with: %s", err.Error())
	}

	absPath, _ := filepath.Abs("./tests/en.yml")
	expectedFiles := []*LocaleFile{
		&LocaleFile{
			Name: "",
			RFC:  "en",
			ID:   "",
			Path: absPath,
		},
	}

	if len(localeFiles) == len(expectedFiles) {
		if err = compareLocaleFiles(localeFiles, expectedFiles); err != nil {
			t.Errorf(err.Error())
		}
	} else {
		t.Errorf(".LocaleFiles should contain %s and not %s", expectedFiles, localeFiles)
	}
}

func TestSourceLocaleFilesTwo(t *testing.T) {
	source := getBaseSource()
	source.File = "./**/<locale_name>.yml"

	localeFiles, err := source.LocaleFiles()

	if err != nil {
		t.Errorf("Should not fail with: %s", err.Error())
	}

	absPath, _ := filepath.Abs("./tests/en.yml")
	expectedFiles := []*LocaleFile{
		&LocaleFile{
			Name: "en",
			RFC:  "",
			ID:   "",
			Path: absPath,
		},
	}

	if len(localeFiles) == len(expectedFiles) {
		if err = compareLocaleFiles(localeFiles, expectedFiles); err != nil {
			t.Errorf(err.Error())
		}
	} else {
		t.Errorf("LocaleFiles should contain %s and not %s", expectedFiles, localeFiles)
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
