package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/coreos/go-semver/semver"
)

var (
	downloadPageURL = "https://phraseapp.com/en/cli"
	releasesURL     = "https://github.com/phrase/phraseapp-client/releases/latest"

	versionCacheFilename = filepath.Join(os.TempDir(), ".phraseapp.version")
)

func CheckForUpdate(w io.Writer) {
	latestVersion, err := getLatestVersion()
	if err != nil {
		printError(err)
		return
	}

	if containsAnySub(strings.ToLower(PHRASEAPP_CLIENT_VERSION), []string{"dev", "test"}) {
		fmt.Fprintf(w, "You're running a development version (%s) of the PhraseApp client! Latest version is %s.\n", PHRASEAPP_CLIENT_VERSION, latestVersion)
		return
	}

	version, err := semver.NewVersion(PHRASEAPP_CLIENT_VERSION)
	if err != nil {
		printError(err)
		return
	}

	if version.LessThan(*latestVersion) {
		fmt.Fprintf(w, "Please consider updating the PhraseApp CLI client (%s < %s)\nYou can get the latest version from %s.\n", version, latestVersion, downloadPageURL)
	}
}

func getLatestVersion() (*semver.Version, error) {
	version, modified, err := getLatestVersionFromCache()

	if err != nil || time.Since(modified) > 24*time.Hour {
		versionOnline, err := getLatestVersionFromURL()
		if err == nil {
			ioutil.WriteFile(versionCacheFilename, []byte(versionOnline.String()), 0600)
		}

		return versionOnline, err
	}

	return version, nil
}

func getLatestVersionFromCache() (version *semver.Version, modified time.Time, err error) {
	f, err := os.Open(versionCacheFilename)
	if err != nil {
		return nil, modified, err
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		return nil, modified, err
	}
	modified = stat.ModTime()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, modified, err
	}

	version, err = semver.NewVersion(string(content))
	return version, modified, err
}

func getLatestVersionFromURL() (*semver.Version, error) {
	req, err := http.NewRequest("HEAD", releasesURL, nil)
	if err != nil {
		return nil, err
	}

	transport := http.Transport{}
	resp, err := transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // body is empty as it is only a HEAD request

	if resp.StatusCode != http.StatusFound {
		return nil, fmt.Errorf("error requesting %s, expected status %d was %d", releasesURL, http.StatusFound, resp.StatusCode)
	}

	url, err := resp.Location()
	if err != nil {
		return nil, err
	}

	latest := path.Base(url.Path)
	if latest == "." || latest == "/" {
		return nil, fmt.Errorf("no valid version segment found")
	}

	return semver.NewVersion(latest)
}
