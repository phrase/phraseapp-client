package main

import "testing"

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
		if err := target.CheckPreconditions(); err != nil {
			t.Errorf("CheckPrecondition should not fail with: %s", err.Error())
		}
	}
}
