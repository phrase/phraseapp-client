package main

import (
	"strings"
	"testing"

	"github.com/phrase/phraseapp-go/phraseapp"
)

func getBaseTarget() *Target {
	target := &Target{
		File:          "./tests/<locale_code>.yml",
		ProjectID:     "project-id",
		AccessToken:   "access-token",
		FileFormat:    "yml",
		Params:        new(PullParams),
		RemoteLocales: getBaseLocales(),
	}
	return target
}

func TestTargetFields(t *testing.T) {
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

func TestPreconditions(t *testing.T) {
	target := &Target{
		File:        "./tests/en.yml",
		ProjectID:   "project-id",
		AccessToken: "access-token",
		FileFormat:  "yml",
		Params:      new(PullParams),
		RemoteLocales: []*phraseapp.Locale{
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
		},
	}

	// no information given
	expect := "Could not find any locale information."
	err := target.CheckPreconditions()
	if err == nil {
		t.Errorf("Expected to fail for pattern %q. Did not fail.", target.File)
	} else if !strings.Contains(err.Error(), expect) {
		t.Errorf("Expected to fail with %q got %q", expect, err)
	}

	// placeholder used
	target.File = "some/path/<locale_code>.yml"
	err = target.CheckPreconditions()
	if err != nil {
		t.Errorf("Should not have failed for pattern %q. Error was: %q", target.File, err)
	}

	// locale_id set
	target.File = "some/path/en.yml"
	target.Params.LocaleID = "en"
	err = target.CheckPreconditions()
	if err != nil {
		t.Errorf("Should not have failed for pattern %q. Error was: %q", target.File, err)
	}

	// locale_id and placeholder set
	expect = "Found 'locale_id' in params and a (<locale_code|locale_name>) placeholder."
	target.File = "some/path/<locale_code>.yml"
	target.Params.LocaleID = "en"
	err = target.CheckPreconditions()
	if err == nil {
		t.Errorf("Expected to fail for pattern %q. Did not fail.", target.File)
	} else if !strings.Contains(err.Error(), expect) {
		t.Errorf("Expected to fail with %q got %q", expect, err)
	}

	// tag in placeholder but no tag provided
	target.Params.LocaleID = "en"
	target.File = "some/<tag>/en.yml"
	expect = "Using <tag> placeholder but no tags were provided."
	err = target.CheckPreconditions()
	if err == nil {
		t.Errorf("Expected to fail for pattern %q. Did not fail.", target.File)
	} else if !strings.Contains(err.Error(), expect) {
		t.Errorf("Expected to fail with %q got %q", expect, err)
	}
}

func TestPlaceholderPreconditions(t *testing.T) {
	target := getBaseTarget()
	for _, file := range []string{
		"",
		"no_extension",
		"./<locale_code>/<locale_code>.yml",
		"./**/**/en.yml",
		"./**/*/*/en.yml",
		"./**/*/en.yml",
		"./en.yml",
		"./**/*/<locale_name>/<locale_code>/<tag>.yml",
	} {
		target.File = file
		if err := target.CheckPreconditions(); err == nil {
			t.Errorf("CheckPrecondition did not fail for pattern: '%s'", file)
		}
	}

	for _, file := range []string{
		"./<tag>/<locale_code>.yml",
		"./<locale_name>/<locale_code>/<tag>.yml",
	} {
		target.File = file
		target.Params.Tag = sPt("any tag")
		if err := target.CheckPreconditions(); err != nil {
			t.Errorf("CheckPrecondition should not fail with: %s", err.Error())
		}
	}
}

func sPt(s string) *string {
	return &s
}
