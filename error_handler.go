package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"runtime/debug"

	ct "github.com/daviddengcn/go-colortext"
	"github.com/phrase/phraseapp-go/phraseapp"
)

const errorsEndpoint = "https://phraseapp.com/errors"

type AppCrash struct {
	Message    string `json:"message"`
	App        string `json:"app"`
	AppVersion string `json:"app_version"`
	ErrorData  `json:"data"`
}

type ErrorData struct {
	Context    string `json:"context"`
	Last8      string `json:"last8"`
	ProjectID  string `json:"project_id"`
	ClientInfo `json:"client_info"`
	Arch       string   `json:"arch"`
	Os         string   `json:"os"`
	StackTrace []string `json:"stack_trace"`
	RawStack   string   `json:"raw_stack_trace"`
}

func ReportError(r interface{}, cfg *phraseapp.Config) {
	last8, projectID := identification(cfg)
	stackTrace := NewStackTrace(string(debug.Stack()))
	crash := &AppCrash{
		Message:    fmt.Sprintf("%s", r),
		App:        "phraseapp-client",
		AppVersion: PHRASEAPP_CLIENT_VERSION,
		ErrorData: ErrorData{
			Context:    stackTrace.ErrorContext(),
			Last8:      last8,
			ProjectID:  projectID,
			ClientInfo: NewInfo(),
			Arch:       runtime.GOARCH,
			Os:         runtime.GOOS,
			StackTrace: stackTrace.ErrorList(),
			RawStack:   string(debug.Stack()),
		},
	}

	body, err := json.Marshal(crash)
	if err != nil {
		return
	}

	response, err := http.Post(errorsEndpoint, "application/json", bytes.NewBuffer(body))

	if err != nil {
		return
	}

	response.Body.Close()
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

func printErr(err error) {
	ct.Foreground(ct.Red, true)
	fmt.Fprintf(os.Stderr, "\nERROR: %s\n", err)
	ct.ResetColor()
}
