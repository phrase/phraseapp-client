package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dynport/dgtk/version"
)

const cliLandingPageUrl = "https://phraseapp.com/en/cli"

var releaseURL = "https://github.com/phrase/phraseapp-client/releases/latest"

var PHRASEAPP_VERSION_TMP_FILE = "/tmp/.phraseapp.version"

func ValidateVersion() {
	if err := validateVersionWithErr(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n\n")
	}
}

func validateVersionWithErr() error {
	currentVersion, err := readCurrentVersion()
	if err != nil {
		return err
	}
	lowerVersion := strings.ToLower(PHRASEAPP_CLIENT_VERSION)
	if strings.Contains(lowerVersion, "test") || strings.Contains(lowerVersion, "dev") {
		return fmt.Errorf("You're running a development version (%s) of the PhraseApp client! Latest version is %s", PHRASEAPP_CLIENT_VERSION, currentVersion)
	}
	clientVersion, err := version.NewFromString(PHRASEAPP_CLIENT_VERSION)
	if err != nil {
		return err
	}
	if clientVersion.Less(currentVersion) {
		return fmt.Errorf("Please consider updating the PhraseApp CLI client (%s < %s)\nSee %s", PHRASEAPP_CLIENT_VERSION, currentVersion, cliLandingPageUrl)
	}
	return nil
}

func getLatestReleaseVersion() (string, error) {
	req, err := http.NewRequest("HEAD", releaseURL, nil)
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
		return "", fmt.Errorf("error requesting %s, expected status %d was %d", releaseURL, 302, resp.StatusCode)
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

func readCurrentVersion() (*version.Version, error) {
	v, err := readCurrentVersionString()
	if err != nil {
		return nil, err
	}
	return version.NewFromString(v)
}

func readCurrentVersionString() (string, error) {
	cached, modTime, err := readCachedVersionString()
	if err != nil || time.Since(*modTime) > time.Hour {
		return updateCachedVersion()
	}
	return cached, nil
}

func updateCachedVersion() (string, error) {
	currentVersion, err := getLatestReleaseVersion()
	if err != nil {
		return "", err
	}
	if err == nil { // persist the version for the next hour
		_ = ioutil.WriteFile(PHRASEAPP_VERSION_TMP_FILE, []byte(currentVersion), 0600)
	}
	return currentVersion, nil
}

func readCachedVersionString() (string, *time.Time, error) {
	f, err := os.Open(PHRASEAPP_VERSION_TMP_FILE)
	if err != nil {
		return "", nil, err
	}
	defer f.Close()
	stat, err := f.Stat()
	if err != nil {
		return "", nil, err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", nil, err
	}
	if len(b) == 0 {
		return "", nil, fmt.Errorf("file %s was empty", PHRASEAPP_VERSION_TMP_FILE)
	}
	mt := stat.ModTime()
	return string(b), &mt, nil
}
