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
	Last8Token string `json:"last_8_token"`
	ProjectID  string `json:"project_id"`
	Arch       string `json:"arch"`
	Os         string `json:"os"`
	Stack      string `json:"stack_trace"`
	ClientInfo string `json:"client_info"`
}

func identification(cfg *phraseapp.Config) (string, string) {
	var last8 string
	var projectID string
	if cfg != nil {
		if len(cfg.Token) == 64 {
			last8 = cfg.Token[len(cfg.Token)-8:]
		}
		projectID = cfg.DefaultProjectID
	}

	return last8, projectID
}

func ReportError(name string, r interface{}, cfg *phraseapp.Config) {
	message := fmt.Sprintf("%s", r)

	last8, projectID := identification(cfg)
	body, err := createBody(name, message, last8, projectID)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", "https://phraseapp.com/errors", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	resp.Body.Close()
}

func createBody(name, message, last8, projectID string) ([]byte, error) {
	crash := &AppCrash{
		App:        "phraseapp-client",
		AppVersion: PHRASEAPP_CLIENT_VERSION,
		ErrorData: ErrorData{
			Last8Token: last8,
			ProjectID:  projectID,
			Arch:       runtime.GOARCH,
			Os:         runtime.GOOS,
			Stack:      string(debug.Stack()),
			ClientInfo: GetInfo(),
		},
		Name:    name,
		Message: message,
	}

	body, err := json.Marshal(crash)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func printErr(err error) {
	ct.Foreground(ct.Red, true)
	fmt.Fprintf(os.Stderr, "\nERROR: %s\n", err)
	ct.ResetColor()
}
