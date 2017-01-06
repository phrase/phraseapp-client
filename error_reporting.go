package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"runtime"

	bserrors "github.com/bugsnag/bugsnag-go/errors"
	"github.com/phrase/client-error-proxy/errors"
	"github.com/phrase/phraseapp-go/phraseapp"
)

const DefaultErrorReportingEndpoint = "https://phraseapp-client-errors.herokuapp.com/errors"

func reportError(cliErr *bserrors.Error, cfg *phraseapp.Config) {
	serializableErr := errors.NewFromBugsnagError(cliErr)

	serializableErr.MetaData = errors.MetaData{
		Environment: errors.Environment{
			Arch: runtime.GOARCH,
			OS:   runtime.GOOS,
		},
		Project: errors.Project{
			ID: projectID(cfg),
		},
		User: errors.User{
			AccessToken: abbreviatedToken(cfg),
			Username:    username(cfg),
		},
		Version: errors.Version{
			BuiltAt:                  BUILT_AT,
			Revision:                 REVISION,
			GeneratorRevision:        RevisionGenerator,
			DocRevision:              RevisionDocs,
			LibraryRevision:          LIBRARY_REVISION,
			LibraryGeneratorRevision: phraseapp.RevisionGenerator,
			LibraryDocRevision:       phraseapp.RevisionDocs,
			GoVersion:                runtime.Version(),
		},
	}

	jsonErr, err := json.Marshal(serializableErr)
	if err != nil {
		return
	}

	endpoint := os.Getenv("ERROR_REPORTING_ENDPOINT")
	if endpoint == "" {
		endpoint = DefaultErrorReportingEndpoint
	}

	response, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonErr))
	if err != nil {
		return
	}

	response.Body.Close()
}

func projectID(cfg *phraseapp.Config) string {
	if cfg != nil {
		return cfg.DefaultProjectID
	}

	return ""
}

func abbreviatedToken(cfg *phraseapp.Config) string {
	if cfg != nil && cfg.Credentials != nil {
		if len(cfg.Token) == 64 {
			return cfg.Token[len(cfg.Token)-8:]
		}
	}

	return ""
}

func username(cfg *phraseapp.Config) string {
	if cfg != nil && cfg.Credentials != nil {
		return cfg.Username
	}

	return ""
}
