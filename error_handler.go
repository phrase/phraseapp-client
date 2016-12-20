package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"runtime"

	"error-proxy/errors"

	bserrors "github.com/bugsnag/bugsnag-go/errors"
	"github.com/phrase/phraseapp-go/phraseapp"
)

const ErrorReportingEndpoint = "http://localhost:8080/error"

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

	response, err := http.Post(ErrorReportingEndpoint, "application/json", bytes.NewBuffer(jsonErr))
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
