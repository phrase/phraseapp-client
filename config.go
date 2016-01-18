package main

import (
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"

	"github.com/phrase/phraseapp-go/phraseapp"
)

const configName = ".phraseapp.yml"

func ReadConfig() (*phraseapp.Config, error) {
	cfg := new(phraseapp.Config)

	content, err := configContent()
	switch {
	case err != nil:
		return nil, err
	case content == nil:
		return cfg, nil
	default:
		return cfg, yaml.Unmarshal(content, cfg)
	}
}

func configContent() ([]byte, error) {
	path, err := configPath()
	switch {
	case err != nil:
		return nil, err
	case path == "":
		return nil, nil
	default:
		return ioutil.ReadFile(path)
	}
}

func configPath() (string, error) {
	if envConfig := os.Getenv("PHRASEAPP_CONFIG"); envConfig != "" {
		possiblePath := path.Join(envConfig)
		if _, err := os.Stat(possiblePath); err == nil {
			return possiblePath, nil
		}
	}

	callerPath, err := os.Getwd()
	if err != nil {
		return "", err
	}

	possiblePath := path.Join(callerPath, configName)
	if _, err := os.Stat(possiblePath); err == nil {
		return possiblePath, nil
	}

	possiblePath = defaultConfigDir()
	if _, err := os.Stat(possiblePath); err != nil && !os.IsNotExist(err) {
		return "", nil
	}

	return possiblePath, nil
}

