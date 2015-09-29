package main

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"
)

func getBaseTarget() *Target {
	target := &Target{
		File:        "./tests/<locale_code>.yml",
		ProjectID:   "project-id",
		AccessToken: "access-token",
		FileFormat:  "yml",
		Params: &PullParams{
			FileFormat:                 "",
			LocaleID:                   "",
			Encoding:                   "",
			ConvertEmoji:               false,
			IncludeEmptyTranslations:   false,
			KeepNotranslateTags:        false,
			SkipUnverifiedTranslations: false,
			Tag: "",
		},
		RemoteLocales: getBaseLocales(),
	}
	return target
}

func TestPullPreconditions(t *testing.T) {
	fmt.Println("Pull#Target#CheckPreconditions")
	target := getBaseTarget()
	for _, file := range []string{
		"",
		"no_extension",
		"./<locale_code>/<locale_code>.yml",
		"./**/**/en.yml",
		"./**/*/*/en.yml",
		"./**/*/en.yml",
		"./**/*/<locale_name>/<locale_code>/<tag>.yml",
	} {
		target.File = file
		if err := target.CheckPreconditions(); err == nil {
			t.Errorf("CheckPrecondition did not fail for pattern: '%s'", file)
		}
	}

	for _, file := range []string{
		"./<tag>/<locale_code>.yml",
		"./en.yml",
		"./<locale_name>/<locale_code>/<tag>.yml",
	} {
		target.File = file
		if err := target.CheckPreconditions(); err != nil {
			t.Errorf("CheckPrecondition should not fail with: %s", err.Error())
		}
	}
}

func TestTargetFields(t *testing.T) {
	fmt.Println("Pull#Target#Fields")
	target := getBaseTarget()

	if target.File != "./tests/<locale_code>.yml" {
		t.Errorf("Expected File to be %s and not %s", "./tests/<locale_code>.yml", target.File)
	}

	if target.AccessToken != "access-token" {
		t.Errorf("Expected AccesToken to be %s and not %s", "access-token", target.AccessToken)
	}

	if target.ProjectID != "project-id" {
		t.Errorf("Expected ProjectID to be %s and not %s", "project-id", target.ProjectID)
	}

	if target.FileFormat != "yml" {
		t.Errorf("Expected FileFormat to be %s and not %s", "yml", target.FileFormat)
	}

}

func TestTargetLocaleFiles(t *testing.T) {
	fmt.Println("Pull#Target#LocaleFiles")
	target := getBaseTarget()
	localeFiles, err := target.LocaleFiles()

	if err != nil {
		t.Errorf("Should not fail with: %s", err.Error())
	}

	enPath, _ := filepath.Abs("./tests/en.yml")
	dePath, _ := filepath.Abs("./tests/de.yml")
	expectedFiles := []*LocaleFile{
		&LocaleFile{
			Name: "english",
			RFC:  "en",
			ID:   "en-locale-id",
			Path: enPath,
		},
		&LocaleFile{
			Name: "german",
			RFC:  "de",
			ID:   "de-locale-id",
			Path: dePath,
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

func TestReplacePlaceholders(t *testing.T) {
	fmt.Println("Pull#Target#ReplacePlaceholders")
	target := getBaseTarget()
	target.File = "./<locale_code>/<tag>/<locale_name>.yml"
	localeFile := &LocaleFile{
		Name: "english",
		RFC:  "en",
		ID:   "en-locale-id",
		Tag:  "abc",
		Path: "",
	}
	newPath, err := target.ReplacePlaceholders(localeFile)
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}

	if !strings.HasSuffix(newPath, "/en/abc/english.yml") {
		t.Errorf("Expected the new path to eql '%s' and not %s", "/en/abc/english.yml", newPath)
	}
}
