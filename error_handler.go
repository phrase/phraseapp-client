package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/phrase/phraseapp-client/Godeps/_workspace/src/github.com/daviddengcn/go-colortext"
	"github.com/phrase/phraseapp-client/Godeps/_workspace/src/github.com/phrase/phraseapp-go/phraseapp"
	"net/http"
	"os"
	"runtime"
	"runtime/debug"
)

type AppCrash struct {
	App        string `json:"app"`
	AppVersion string `json:"app_version"`
	ErrorData  `json:"data"`
	Message    string `json:"message"`
	Name       string `json:"name"`
}

type ErrorData struct {
	ShortAccessToken string `json:"short_access_token"`
	ProjectID        string `json:"project_id"`
	Arch             string `json:"arch"`
	Os               string `json:"os"`
	Stack            string `json:"stack_trace"`
	ClientInfo       string `json:"client_info"`
}

func identification(cfg *phraseapp.Config) (string, string) {
	var shortToken string
	var projectID string
	if cfg != nil {
		if len(cfg.Token) == 64 {
			shortToken = cfg.Token[len(cfg.Token)-8:]
		}
		projectID = cfg.DefaultProjectID
	}

	return shortToken, projectID
}

func ReportError(name string, r interface{}, cfg *phraseapp.Config) {
	message := fmt.Sprintf("%s", r)

	shortToken, projectID := identification(cfg)
	body, err := createBody(name, message, shortToken, projectID)
	if err != nil {
		return
	}
	response, err := http.Post("https://phraseapp.com/errors", "application/json", bytes.NewBuffer(body))

	if err != nil {
		return
	}

	response.Body.Close()
}

func createBody(name, message, shortToken, projectID string) ([]byte, error) {
	crash := &AppCrash{
		App:        "phraseapp-client",
		AppVersion: PHRASEAPP_CLIENT_VERSION,
		ErrorData: ErrorData{
			ShortAccessToken: shortToken,
			ProjectID:        projectID,
			Arch:             runtime.GOARCH,
			Os:               runtime.GOOS,
			Stack:            string(debug.Stack()),
			ClientInfo:       GetInfo(),
		},
		Name:    name,
		Message: message,
	}

	return json.Marshal(crash)
}

func printErr(err error) {
	ct.Foreground(ct.Red, true)
	fmt.Fprintf(os.Stderr, "\nERROR: %s\n", err)
	ct.ResetColor()
}
