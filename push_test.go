package main

import (
	"github.com/phrase/phraseapp-go/phraseapp"
	"path/filepath"
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
			t.Errorf("source.CheckPrecondition did no fail!")
		}
	}

	for _, file := range []string{
		"./<tag>/<locale_code>.yml",
		"./**/*/en.yml",
		"./**/*/<locale_name>/<locale_code>/<tag>.yml",
	} {
		source.File = file
		if err := source.CheckPreconditions(); err != nil {
			t.Errorf("source.CheckPrecondition should not fail with: %s", err.Error())
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
		compareLocaleFiles(t, localeFiles, expectedFiles)
	} else {
		t.Errorf("source.LocaleFiles should contain %s and not %s", expectedFiles, localeFiles)
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
		compareLocaleFiles(t, localeFiles, expectedFiles)
	} else {
		t.Errorf("source.LocaleFiles should contain %s and not %s", expectedFiles, localeFiles)
	}
}

func compareLocaleFiles(t *testing.T, actualFiles LocaleFiles, expectedFiles LocaleFiles) {
	for idx, localeFile := range actualFiles {
		expected := expectedFiles[idx]
		actual := localeFile
		if expected.Path != localeFile.Path {
			t.Errorf("Expected Path %s should eql %s", expected.Path, actual.Path)
		}
		if expected.Name != localeFile.Name {
			t.Errorf("Expected Name %s should eql %s", expected.Name, actual.Name)
		}
		if expected.RFC != localeFile.RFC {
			t.Errorf("Expected RFC %s should eql %s", expected.RFC, actual.RFC)
		}
		if expected.ID != localeFile.ID {
			t.Errorf("Expected ID %s should eql %s", expected.ID, actual.ID)
		}
	}
}
