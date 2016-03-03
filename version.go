package main

import (
	"fmt"
	"github.com/phrase/phraseapp-client/Godeps/_workspace/src/github.com/phrase/phraseapp-go/phraseapp"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const PHRASEAPP_VERSION_TMP_FILE = "/tmp/.phraseapp.version"

func ValidateVersion() {
	var version string
	stat, err := os.Stat(PHRASEAPP_VERSION_TMP_FILE)
	if os.IsNotExist(err) || time.Since(stat.ModTime()) > time.Hour {
		// fetch new version, if not done so or over an hour ago
		version, err = getLatestReleaseVersion()

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

	tmpVersion := strings.ToLower(phraseapp.ClientVersion)
	if strings.Contains(tmpVersion, "test") || strings.Contains(tmpVersion, "dev") {
		fmt.Fprintf(os.Stderr, "You're running a development version (%s) of the PhraseApp client! Latest version is %s\n\n", phraseapp.ClientVersion, version)
		return
	}
	if err == nil && isClientVersionOutdated(version) {
		fmt.Fprintf(os.Stderr, "Please consider updating the PhraseApp CLI client (%s < %s)\nSee https://phraseapp.com/en/cli\n\n", phraseapp.ClientVersion, version)
	}
}

func isClientVersionOutdated(version string) bool {
	if version == phraseapp.ClientVersion {
		return false
	}

	versionParts := strings.Split(version, ".")[:3]
	clientVersionParts := strings.Split(phraseapp.ClientVersion, ".")[:3]

	versionNum, err := strconv.Atoi(strings.Join(versionParts, ""))
	if err != nil {
		return false
	}
	clientVersionNum, err := strconv.Atoi(strings.Join(clientVersionParts, ""))
	if err != nil {
		return false
	}
	return versionNum > clientVersionNum
}

func getLatestReleaseVersion() (string, error) {
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
