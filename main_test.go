package main

import (
	"testing"

	"gopkg.in/yaml.v1"
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
	if sourceParams.FileFormat != "strings" {
		t.Errorf("Expected FileFormat of first target to be %s and not %s", "strings", sourceParams.FileFormat)
	}
}
