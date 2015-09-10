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
