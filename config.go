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
	content, err := configContent()
	if err != nil {
		return nil, err
	}

	return parseCredentials(content)
}

func ConfigDefaultParams() (phraseapp.DefaultParams, error) {
	content, err := configContent()
	if err != nil {
		return nil, err
	}

	return parseDefaults(content)
}

func ConfigCallArgs() (map[string]string, error) {
	content, err := configContent()
	if err != nil {
		return nil, err
	}

	return parseCallArgs(content)
}

func ConfigPushPull() (*PushPullConfig, error) {
	content, err := configContent()
	if err != nil {
		return nil, err
	}

	return parsePushPullArgs(content)
}

// Paths and content
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

func configContent() (string, error) {
	path, err := phraseConfigPath()
	if err != nil {
		return "", nil
	}

	b, err := bytesAtPath(path)
	if err != nil {
		return "", nil
	}
	return string(b), nil
}

func bytesAtPath(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ioutil.ReadAll(f)
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
		AccessToken string `yaml:"access_token"`
		Username    string
		TFA         bool
	}
}

func parseCredentials(yml string) (*phraseapp.AuthCredentials, error) {
	var conf *credentialConf

	if err := yaml.Unmarshal([]byte(yml), &conf); err != nil {
		return nil, err
	}

	phrase := conf.Phraseapp
	credentials := &phraseapp.AuthCredentials{Token: phrase.AccessToken, Username: phrase.Username, TFA: phrase.TFA}

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

type Params struct {
	File        string
	AccessToken string `yaml:"access_token"`
	ProjectId   string `yaml:"project_id"`
	Format      string
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
