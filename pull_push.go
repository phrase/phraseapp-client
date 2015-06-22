package main

import (
	"fmt"
	"github.com/phrase/phraseapp-api-client/Godeps/_workspace/src/gopkg.in/yaml.v2"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func ConfigPushPull() (*PushPullConfig, error) {
	content, err := ConfigContent()
	if err != nil {
		return nil, err
	}

	return parsePushPullArgs(content)
}

func FileStrategy(file string, fileFormat string) ([]string, error) {

	absPath, err := filepath.Abs(file)

	if err != nil {
		return nil, err
	}

	recursive := path.Join("**", "*")

	switch {
	case strings.HasSuffix(absPath, recursive):
		root := trimSuffix(absPath, recursive)
		return recursiveStrategy(root, fileFormat)

	case strings.HasSuffix(absPath, "*"):
		return singleDirectoryStrategy(absPath, fileFormat)

	}

	err = fileExists(absPath)
	if err != nil {
		return nil, err
	}

	return []string{absPath}, nil
}

// File handling
func recursiveStrategy(root, fileFormat string) ([]string, error) {
	err := fileExists(root)
	if err != nil {
		return nil, err
	}

	fileList := []string{}
	err = filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if isLocaleFile(f.Name(), fileFormat) {
			fileList = append(fileList, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return fileList, nil
}

func singleDirectoryStrategy(file, fileFormat string) ([]string, error) {
	files, err := filepath.Glob(file)
	if err != nil {
		return nil, err
	}
	localeFiles := []string{}
	for _, f := range files {
		if isLocaleFile(f, fileFormat) {
			localeFiles = append(localeFiles, f)
		} else {
			localeFiles = append(localeFiles, f)
		}
	}
	return localeFiles, nil
}

func isLocaleFile(file, extension string) bool {
	fileExtension := fmt.Sprintf(".%s", extension)
	return strings.HasSuffix(file, fileExtension)
}

func trimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

func fileExists(absPath string) error {
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return fmt.Errorf("no such file or directory:", absPath)
	}
	return nil
}

// Parsing
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
func PrettyPrint(projectId, accessToken string, sources []Params, targets []Params) {
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
