package main

/*
import (
	"path/filepath"

	"github.com/phrase/phraseapp-api-client/Godeps/_workspace/src/gopkg.in/yaml.v2"
	"github.com/phrase/phraseapp-go/phraseapp"
)

func PushSourcesFromConfig() (Sources, error) {
	content, err := ConfigContent()
	if err != nil {
		return nil, err
	}

	return parsePush(content)
}

func Push(p *PhrasePath, source Source) ([]string, error) {
	authenticate()

	paths, err := createFilesForPush(p, source)
	if err != nil {
		return nil, err
	}

	return paths, nil
}

type Source *PushArgs
type Sources []Source

type PushArgs struct {
	File        string      `yaml:"file,omitempty"`
	ProjectId   string      `yaml:"project_id,omitempty"`
	AccessToken string      `yaml:"access_token,omitempty"`
	Params      *PushParams `yaml:"params,omitempty"`
}

type PushParams struct {
	FileFormat string `yaml:"file_format,omitempty"`
	LocaleId   string `yaml:"locale_id,omitempty"`
}

func createFilesForPush(p *PhrasePath, source Source) ([]string, error) {
	files := []string{}

	switch {
	case p.LocaleTagInPath:
		locales, err := phraseapp.LocalesList(source.ProjectId, 1, 25)
		if err != nil {
			return nil, err
		}
		newFiles, err := CreateFilesWithLocales(p, locales)
		if err != nil {
			return nil, err
		}
		files = append(files, newFiles...)
	default:
		absPath, err := filepath.Abs(p.UserPath)
		if err != nil {
			return nil, err
		}
		err = CreateFile(absPath)
		if err != nil {
			return nil, err
		}
		files = append(files, absPath)
	}

	return pushFileStrategy(p, files)
}

func pushFileStrategy(p *PhrasePath, paths []string) ([]string, error) {
	switch {
	case p.Mode == "":
		return paths, nil

	case p.Mode == "*":
		extendedPaths := []string{}
		for _, absPath := range paths {
			pathsPerDirectory, err := SingleDirectoryStrategy(absPath, "")
			if err != nil {
				return nil, err
			}
			extendedPaths = append(extendedPaths, pathsPerDirectory...)
		}
		return extendedPaths, nil
	case p.Mode == "**\/*":
		return paths, nil

	default:
		return paths, nil
	}
}

// Parsing
type PushConfig struct {
	Phraseapp struct {
		AccessToken string `yaml:"access_token"`
		ProjectId   string `yaml:"project_id"`
		Push        struct {
			Sources Sources
		}
	}
}

func parsePush(yml string) (Sources, error) {
	var config *PushConfig

	err := yaml.Unmarshal([]byte(yml), &config)
	if err != nil {
		return nil, err
	}

	token := config.Phraseapp.AccessToken
	projectId := config.Phraseapp.ProjectId
	sources := config.Phraseapp.Push.Sources

	for _, source := range sources {
		if source.ProjectId == "" {
			source.ProjectId = projectId
		}
		if source.AccessToken == "" {
			source.AccessToken = token
		}
	}

	return sources, nil
}
*/
