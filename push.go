package main

import (
	"fmt"

	"github.com/phrase/phraseapp-go/phraseapp"
)

type PushCommand struct {
	*phraseapp.Config
}

func (cmd *PushCommand) Run() error {
	if cmd.Debug {
		// suppresses content output
		cmd.Debug = false
		Debug = true
	}

	client, err := newClient(cmd.Config.Credentials)
	if err != nil {
		return err
	}

	sources, err := SourcesFromConfig(cmd)
	if err != nil {
		return err
	}

	if err := sources.Validate(); err != nil {
		return err
	}

	formatMap, err := GetFormats(client)
	if err != nil {
		return fmt.Errorf("Error retrieving format list from PhraseApp: %s", err)
	}

	for _, source := range sources {
		formatName := source.GetFileFormat()
		if val, ok := formatMap[formatName]; ok {
			source.Format = val
		}

		if source.Format == nil {
			return fmt.Errorf("Format %q of source %q is not supported by PhraseApp!", formatName, source.File)
		}
	}

	projectIdToLocales, err := LocalesForProjects(client, sources)
	if err != nil {
		return err
	}
	for _, source := range sources {
		val, ok := projectIdToLocales[source.ProjectID]
		if ok {
			source.RemoteLocales = val
		}
	}

	for _, source := range sources {
		err := source.Push(client)
		if err != nil {
			return err
		}
	}
	return nil
}
