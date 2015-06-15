package main

import (
	"github.com/phrase/phraseapp-api-client/Godeps/_workspace/src/gopkg.in/yaml.v2"
	"github.com/phrase/phraseapp-go/phraseapp"
	"io/ioutil"
	"os"
	"os/user"
	"path"
)

const configName = ".phraseapp.yml"
const defaultDir = ".config/phraseapp"

func ConfigDefaultCredentials() (*phraseapp.AuthCredentials, error) {
	path, err := phraseConfigPath()
	if err != nil {
		return nil, err
	}

	content, err := contentAtPath(path)
	if err != nil {
		return nil, err
	}

	return parseCredentials(content)
}

func ConfigDefaultParams() (phraseapp.DefaultParams, error) {
	path, err := phraseConfigPath()
	if err != nil {
		return nil, err
	}

	content, err := contentAtPath(path)
	if err != nil {
		return nil, err
	}

	return parseDefaults(content)
}

func ConfigCallArgs() (*CallArgs, error) {
	path, err := phraseConfigPath()
	if err != nil {
		return nil, err
	}

	content, err := contentAtPath(path)
	if err != nil {
		return nil, err
	}

	return parseCallArgs(content)
}

// Path utils
func contentAtPath(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func phraseConfigPath() (string, error) {
	callerPath, err := os.Getwd()
	if err != nil {
		return "", err
	}

	files, _ := ioutil.ReadDir(callerPath)
	for _, f := range files {
		if f.Name() == configName {
			return callerPath + configName, nil
		}
	}

	return defaultConfigDir()
}

func defaultConfigDir() (string, error) {
	usr, e := user.Current()
	if e != nil {
		return "", nil
	}
	return path.Join(usr.HomeDir, defaultDir, configName), nil
}

// Parsing
type credentialConf struct {
	Phraseapp struct {
		Token    string
		Username string
		TFA      bool
	}
}

func parseCredentials(yml string) (*phraseapp.AuthCredentials, error) {
	var conf *credentialConf

	if err := yaml.Unmarshal([]byte(yml), &conf); err != nil {
		return nil, err
	}

	phrase := conf.Phraseapp
	credentials := &phraseapp.AuthCredentials{Token: phrase.Token, Username: phrase.Username, TFA: phrase.TFA}

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
		return nil, err
	}

	return conf.Phraseapp.Defaults, nil
}

type CallArgs struct {
	Phraseapp struct {
		ProjectId string
		Page      int
		PerPage   int
	}
}

func parseCallArgs(yml string) (*CallArgs, error) {
	var callArgs *CallArgs

	err := yaml.Unmarshal([]byte(yml), &callArgs)
	if err != nil {
		return nil, err
	}

	return callArgs, nil
}
