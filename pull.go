package main

import (
	"fmt"

	"github.com/phrase/phraseapp-client/paclient"
	"github.com/phrase/phraseapp-go/phraseapp"
	yaml "gopkg.in/yaml.v2"
)

type PullCommand struct {
	*phraseapp.Config
}

func (cmd *PullCommand) Run() error {
	if cmd.Debug {
		// suppresses content output
		cmd.Debug = false
		Debug = true
	}
	client, err := newClient(cmd.Config.Credentials)
	if err != nil {
		return err
	}

	targets, err := TargetsFromConfig(cmd)
	if err != nil {
		return err
	}

	projectIdToLocales, err := LocalesForProjects(client, targets)
	if err != nil {
		return err
	}
	for _, target := range targets {
		val, ok := projectIdToLocales[target.ProjectID]
		if ok {
			target.RemoteLocales = val
		}
	}

	for _, target := range targets {
		err := target.Pull(client)
		if err != nil {
			return err
		}
	}

	return nil
}

func TargetsFromConfig(cmd *PullCommand) (paclient.Targets, error) {
	if cmd.Config.Targets == nil || len(cmd.Config.Targets) == 0 {
		return nil, fmt.Errorf("no targets for download specified")
	}

	tmp := struct {
		Targets paclient.Targets
	}{}
	err := yaml.Unmarshal(cmd.Config.Targets, &tmp)
	if err != nil {
		return nil, err
	}
	tgts := tmp.Targets

	token := cmd.Credentials.Token
	projectId := cmd.Config.DefaultProjectID
	fileFormat := cmd.Config.DefaultFileFormat

	validTargets := []*paclient.Target{}
	for _, target := range tgts {
		if target == nil {
			continue
		}
		if target.ProjectID == "" {
			target.ProjectID = projectId
		}
		if target.AccessToken == "" {
			target.AccessToken = token
		}
		if target.FileFormat == "" {
			target.FileFormat = fileFormat
		}
		validTargets = append(validTargets, target)
	}

	if len(validTargets) <= 0 {
		return nil, fmt.Errorf("no targets could be identified! Refine the targets list in your config")
	}

	return validTargets, nil
}
