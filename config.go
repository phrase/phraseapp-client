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
	content, err := ConfigContent()
	if err != nil {
		content = []byte("{}")
	}

	cfg := new(phraseapp.Config)
	return cfg, yaml.Unmarshal(content, cfg)
}


func ConfigContent() ([]byte, error) {
	path, err := configPath()
	if err != nil {
		return []byte(""), err
	}

	return readFile(path)
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

	return defaultConfigDir(), nil
}

func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return ioutil.ReadAll(file)
}
