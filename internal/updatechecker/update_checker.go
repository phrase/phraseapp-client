package updatechecker

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/coreos/go-semver/semver"
	"github.com/phrase/phraseapp-client/internal/print"
	"github.com/phrase/phraseapp-client/internal/stringz"
)

const downloadPageURL = "https://phraseapp.com/en/cli"

type Checker struct {
	version              string
	versionCacheFilename string
	releasesURL          string
	output               io.Writer
}

func New(version, versionCacheFilename, releasesURL string, output io.Writer) *Checker {
	return &Checker{
		version:              version,
		versionCacheFilename: versionCacheFilename,
		releasesURL:          releasesURL,
		output:               output,
	}
}

func (uc *Checker) Check() {
	latestVersion, err := uc.getLatestVersion()
	if err != nil {
		print.Error(err)
		return
	}

	if stringz.ContainsAnySub(strings.ToLower(uc.version), []string{"dev", "test"}) {
		fmt.Fprintf(uc.output, "You're running a development version (%s) of the PhraseApp client! Latest version is %s.\n", uc.version, latestVersion)
		return
	}

	version, err := semver.NewVersion(uc.version)
	if err != nil {
		print.Error(err)
		return
	}

	if version.LessThan(*latestVersion) {
		fmt.Fprintf(uc.output, "Please consider updating the PhraseApp CLI client (%s < %s)\nYou can get the latest version from %s.\n", version, latestVersion, downloadPageURL)
	}
}

func (uc *Checker) getLatestVersion() (*semver.Version, error) {
	version, modified, err := uc.getLatestVersionFromCache()

	if err != nil || time.Since(modified) > 24*time.Hour {
		versionOnline, err := uc.getLatestVersionFromURL()
		if err == nil {
			ioutil.WriteFile(uc.versionCacheFilename, []byte(versionOnline.String()), 0600)
		}

		return versionOnline, err
	}

	return version, nil
}

func (uc *Checker) getLatestVersionFromCache() (version *semver.Version, modified time.Time, err error) {
	f, err := os.Open(uc.versionCacheFilename)
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

func (uc *Checker) getLatestVersionFromURL() (*semver.Version, error) {
	req, err := http.NewRequest("HEAD", uc.releasesURL, nil)
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
		return nil, fmt.Errorf("error requesting %s, expected status %d was %d", uc.releasesURL, http.StatusFound, resp.StatusCode)
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
