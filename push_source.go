package main

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"

	"github.com/phrase/phraseapp-client/paclient"
	"github.com/phrase/phraseapp-go/phraseapp"
)

func SourcesFromConfig(cmd *PushCommand) (paclient.Sources, error) {
	if cmd.Config.Sources == nil || len(cmd.Config.Sources) == 0 {
		return nil, fmt.Errorf("no sources for upload specified")
	}

	tmp := struct {
		Sources paclient.Sources
	}{}
	err := yaml.Unmarshal(cmd.Config.Sources, &tmp)
	if err != nil {
		return nil, err
	}
	srcs := tmp.Sources

	token := cmd.Credentials.Token
	projectId := cmd.Config.DefaultProjectID
	fileFormat := cmd.Config.DefaultFileFormat

	validSources := []*paclient.Source{}
	for _, source := range srcs {
		if source == nil {
			continue
		}
		if source.ProjectID == "" {
			source.ProjectID = projectId
		}
		if source.AccessToken == "" {
			source.AccessToken = token
		}
		if source.Params == nil {
			source.Params = new(phraseapp.UploadParams)
		}

		if source.Params.FileFormat == nil {
			switch {
			case source.FileFormat != "":
				source.Params.FileFormat = &source.FileFormat
			case fileFormat != "":
				source.Params.FileFormat = &fileFormat
			}
		}
		validSources = append(validSources, source)
	}

	if len(validSources) <= 0 {
		return nil, fmt.Errorf("no sources could be identified! Refine the sources list in your config")
	}

	return validSources, nil
}
