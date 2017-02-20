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
		&LocaleFile{
			Name: "english",
			Code: "en",
			ID:   "en-locale-id",
			Path: enPath,
		},
		&LocaleFile{
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
	newPath, err := resolvedPath(localeFile)
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}

	if !strings.HasSuffix(newPath, "/en/abc/english.yml") {
		t.Errorf("Expected the new path to eql '%s' and not %s", "/en/abc/english.yml", newPath)
	}
}
