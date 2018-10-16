package main

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestPullLocaleFiles(t *testing.T) {
	target := getBaseTarget()
	localeFiles, err := target.LocaleFiles()

	if err != nil {
		t.Errorf("Should not fail with: %s", err.Error())
	}

	enPath, _ := filepath.Abs("./tests/en.yml")
	dePath, _ := filepath.Abs("./tests/de.yml")
	expectedFiles := []*LocaleFile{
		{
			Name: "english",
			Code: "en",
			ID:   "en-locale-id",
			Path: enPath,
		},
		{
			Name: "german",
			Code: "de",
			ID:   "de-locale-id",
			Path: dePath,
		},
	}

	if len(localeFiles) == len(expectedFiles) {
		if err = compareLocaleFiles(localeFiles, expectedFiles); err != nil {
			t.Errorf(err.Error())
		}
	} else {
		t.Errorf("LocaleFiles should contain %v and not %v", expectedFiles, localeFiles)
	}
}

func TestResolvedPath(t *testing.T) {
	target := getBaseTarget()
	target.File = "./<locale_code>/<tag>/<locale_name>.yml"
	localeFile := &LocaleFile{
		Name: "english",
		Code: "en",
		ID:   "en-locale-id",
		Tag:  "abc",
		Path: target.File,
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

func TestLocaleFilesWithMultipleTags(t *testing.T) {
	target := &Target{
		File:          "./tests/<locale_code>/<tag>.yml",
		ProjectID:     "project-id",
		AccessToken:   "access-token",
		FileFormat:    "yml",
		Params:        new(PullParams),
		RemoteLocales: getBaseLocales(),
	}
	tags := "abc,abc2"
	target.Params.Tags = &tags
	target.Params.LocaleID = "en-locale-id"

	files, err := target.LocaleFiles()
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}

	if len(files) != 2 {
		t.Errorf("Expected 2 files and and not %d", len(files))
	}

	if !strings.HasSuffix(files[0].Path, "/tests/en/abc.yml") {
		t.Errorf("File path is '%s' and should end with '%s'", files[0].Path, "/tests/en/abc.yml")
	}

	if !strings.HasSuffix(files[1].Path, "/tests/en/abc2.yml") {
		t.Errorf("File path is '%s' and should end with '%s'", files[1].Path, "/tests/en/abc2.yml")
	}
}
