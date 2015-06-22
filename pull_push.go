package main

import (
	"fmt"
	"github.com/phrase/phraseapp-api-client/Godeps/_workspace/src/gopkg.in/yaml.v2"
	"path"
	"path/filepath"
	"strings"
)

// File Handling
func Normalize() error {
	config, err := configPushPull()
	if err != nil {
		return err
	}
	projectId := config.Phraseapp.ProjectId
	accessToken := config.Phraseapp.AccessToken
	sources := config.Phraseapp.Push.Sources
	targets := config.Phraseapp.Pull.Targets

	prettyPrint(projectId, accessToken, sources, targets)

	for _, source := range sources {

		paths, _ := FileStrategy(source.File)
		fmt.Println(paths)
	}

	return nil
}

func FileStrategy(file string) ([]string, error) {

	localizationDirectory := "*"
	fullRecursive := path.Join("**", "*")

	if strings.HasSuffix(file, fullRecursive) {

		return filepath.Glob(file)

	} else if strings.HasSuffix(file, localizationDirectory) {

		return filepath.Glob(file)

	}

	return []string{file}, nil
}

// Parsing
func configPushPull() (*PushPullConfig, error) {
	content, err := ConfigContent()
	if err != nil {
		return nil, err
	}

	return parsePushPullArgs(content)
}

type Params struct {
	File        string
	AccessToken string `yaml:"access_token"`
	ProjectId   string `yaml:"project_id"`
	Format      string
	LocaleId    string `yaml:"locale_id"`
}

type PushPullConfig struct {
	Phraseapp struct {
		AccessToken string `yaml:"access_token"`
		ProjectId   string `yaml:"project_id"`
		Push        struct {
			Sources []Params
		}
		Pull struct {
			Targets []Params
		}
	}
}

func parsePushPullArgs(yml string) (*PushPullConfig, error) {
	var pushPullConfig *PushPullConfig

	err := yaml.Unmarshal([]byte(yml), &pushPullConfig)
	if err != nil {
		return nil, err
	}

	return pushPullConfig, nil
}

// debugging
func prettyPrint(projectId, accessToken string, sources []Params, targets []Params) {
	fmt.Println("Phraseapp:")
	fmt.Println("  ProjectId:", projectId)
	fmt.Println("  AccessToken:", accessToken)

	fmt.Println("  Push:")
	for i, params := range sources {
		fmt.Println("    Source", i+1)
		fmt.Println("      token:", params.AccessToken)
		fmt.Println("      pid:", params.ProjectId)
		fmt.Println("      file:", params.File)
		fmt.Println("      format:", params.Format)
		fmt.Println("      lid:", params.LocaleId)

	}

	fmt.Println("  Pull:")
	for i, params := range targets {
		fmt.Println("    Targets", i+1)
		fmt.Println("      token:", params.AccessToken)
		fmt.Println("      pid:", params.ProjectId)
		fmt.Println("      file:", params.File)
		fmt.Println("      format:", params.Format)
		fmt.Println("      lid:", params.LocaleId)
	}
}
