package main

import (
	"os"

	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/dynport/dgtk/cli"
	"github.com/phrase/phraseapp-go/phraseapp"
)

func main() {
	Run()
}

func Run() {
	validateVersion()

	cfg, err := phraseapp.ReadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(2)
	}

	r, err := router(cfg)
	if err != nil {
		printErr(err)
		os.Exit(3)
	}

	switch err := r.RunWithArgs(); err {
	case cli.ErrorHelpRequested, cli.ErrorNoRoute:
		os.Exit(1)
	case nil:
		os.Exit(0)
	default:
		printErr(err)
		os.Exit(1)
	}
}

const PHRASEAPP_VERSION_TMP_FILE = "/tmp/.phraseapp.version"

func validateVersion() {
	var version string
	stat, err := os.Stat(PHRASEAPP_VERSION_TMP_FILE)
	if PHRASEAPP_CLIENT_VERSION == "test" {
		fmt.Fprintf(os.Stderr, "You're running a development version of the PhraseApp CLI client!\n\n")
		return
	} else if os.IsNotExist(err) || time.Since(stat.ModTime()) > time.Hour {
		// fetch new version, if not done so or over an hour ago
		version, err = getCurrentVersion()
		if err == nil { // persist the version for the next hour
			err = ioutil.WriteFile(PHRASEAPP_VERSION_TMP_FILE, []byte(version), 0600)
		}
	} else if err == nil {
		// otherwise load the version (fetched less than an hour ago) from the temp file
		var buf []byte
		buf, err = ioutil.ReadFile(PHRASEAPP_VERSION_TMP_FILE)
		if err == nil {
			version = string(buf)
		}
	}

	if err == nil && version != PHRASEAPP_CLIENT_VERSION {
		fmt.Fprintf(os.Stderr, "Please consider updating the PhraseApp CLI client (%s < %s)\nSee https://phraseapp.com/en/cli\n\n", PHRASEAPP_CLIENT_VERSION, version)
	}
}

func getCurrentVersion() (string, error) {
	req, err := http.NewRequest("HEAD", "https://github.com/phrase/phraseapp-client/releases/latest", nil)
	if err != nil {
		return "", err
	}

	transport := http.Transport{}
	resp, err := transport.RoundTrip(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close() // body is empty as it is only a HEAD request

	if resp.StatusCode != 302 {
		return "", fmt.Errorf("failed to request the file")
	}

	url, err := resp.Location()
	if err != nil {
		return "", err
	}

	segments := strings.Split(url.Path, "/")
	for i := len(segments) - 1; i >= 0; i-- {
		if segments[i] != "" {
			return segments[i], nil
		}
	}
	return "", fmt.Errorf("no valid version segment found")
}
