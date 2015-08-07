package main

import (
	"fmt"
	"github.com/phrase/phraseapp-go/phraseapp"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/user"
	"path"
)

const configName = ".phraseapp.yml"
const defaultDir = "./"

func ConfigDefaultCredentials() (*phraseapp.AuthCredentials, error) {
	content, err := ConfigContent()
	if err != nil {
		content = "{}"
	}

	return parseCredentials(content)
}

func ConfigDefaultParams() (phraseapp.DefaultParams, error) {
	content, err := ConfigContent()
	if err != nil {
		content = "{}"
	}

	return parseDefaults(content)
}

func ConfigCallArgs() (map[string]string, error) {
	content, err := ConfigContent()
	if err != nil {
		content = "{}"
	}

	return parseCallArgs(content)
}

// Paths and content
func ConfigContent() (string, error) {
	path, err := phraseConfigPath()
	if err != nil {
		return "", err
	}

	bytes, err := readFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func phraseConfigPath() (string, error) {
	callerPath, err := os.Getwd()
	if err != nil {
		return "", err
	}

	possiblePath := path.Join(callerPath, configName)
	if _, err := os.Stat(possiblePath); err == nil {
		return possiblePath, nil
	}

	return defaultConfigDir()
}

func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return ioutil.ReadAll(file)
}

func defaultConfigDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", nil
	}
	return path.Join(usr.HomeDir, defaultDir, configName), nil
}

// Parsing
type credentialConf struct {
	Phraseapp struct {
		AccessToken string `yaml:"access_token"`
		Host        string `yaml:"host"`
		Debug       bool   `yaml:"verbose"`
		Username    string
		TFA         bool
	}
}

func parseCredentials(yml string) (*phraseapp.AuthCredentials, error) {
	var conf *credentialConf

	if err := yaml.Unmarshal([]byte(yml), &conf); err != nil {
		fmt.Println("Could not parse .phraseapp.yml")
		return nil, err
	}

	phrase := conf.Phraseapp

	credentials := &phraseapp.AuthCredentials{Token: phrase.AccessToken, Username: phrase.Username, TFA: phrase.TFA, Host: phrase.Host, Debug: phrase.Debug}

	return credentials, nil
}

type defaultsConf struct {
	Phraseapp struct {
		Defaults phraseapp.DefaultParams
	}
}

func parseDefaults(yml string) (phraseapp.DefaultParams, error) {
	var conf *defaultsConf

	err := yaml.Unmarshal([]byte(yml), &conf)
	if err != nil {
		fmt.Println("Could not parse .phraseapp.yml")
		return nil, err
	}

	return conf.Phraseapp.Defaults, nil
}

type CallArgs struct {
	Phraseapp struct {
		AccessToken string `yaml:"access_token"`
		ProjectId   string `yaml:"project_id"`
		Page        int
		PerPage     int
	}
}

func parseCallArgs(yml string) (map[string]string, error) {
	var callArgs *CallArgs

	err := yaml.Unmarshal([]byte(yml), &callArgs)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)

	if callArgs != nil {
		m["ProjectId"] = callArgs.Phraseapp.ProjectId
		m["AccessToken"] = callArgs.Phraseapp.AccessToken
	}

	return m, nil
}
