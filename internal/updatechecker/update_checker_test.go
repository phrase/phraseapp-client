package updatechecker

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestUpdateChecker_GetLatestVersionFromURL(t *testing.T) {
	expected := "1.1.3"

	tuc, _, cleanup := newTestUpateChecker("", expected, "", t)
	defer cleanup()

	v, err := tuc.getLatestVersionFromURL()
	if err != nil {
		t.Fatal(err)
	}

	if v.String() != expected {
		t.Errorf("expected latest version to be %q, was %q", expected, v)
	}
}

func TestUpdateChecker_GetLatestVersionFromCache(t *testing.T) {
	expected := "1.1.3"

	tuc, _, cleanup := newTestUpateChecker("", "", expected, t)
	defer cleanup()

	v, _, err := tuc.getLatestVersionFromCache()
	if err != nil {
		t.Errorf("could not read version from cache")
	}

	if v.String() != expected {
		t.Errorf("expected latest version to be %q, was %q", expected, v)
	}
}

func TestUpdateChecker_GetLatestVersion_withCache(t *testing.T) {
	expected := "1.1.3"

	tuc, _, cleanup := newTestUpateChecker("2.0.0", expected, "", t)
	defer cleanup()

	v, err := tuc.getLatestVersion()
	if err != nil {
		t.Fatal(err)
	}

	if v.String() != expected {
		t.Errorf("expected latest version to be %q, was %q", expected, v)
	}
}

func TestUpdateChecker_GetLatestVersion_withoutCache(t *testing.T) {
	expected := "2.0.0"

	tuc, _, cleanup := newTestUpateChecker("", expected, "", t)
	defer cleanup()

	v, err := tuc.getLatestVersion()
	if err != nil {
		t.Fatal(err)
	}

	if v.String() != expected {
		t.Errorf("expected latest version to be %q, was %q", expected, v)
	}

	cv, _, err := tuc.getLatestVersionFromCache()
	if err != nil {
		t.Errorf("could not read version from cache")
	}

	if cv.String() != expected {
		t.Errorf("expected cached version after update to be %q, was %q", expected, v)
	}
}

func TestUpdateChecker_GetLatestVersion_withInvalidCache(t *testing.T) {
	expected := "2.0.0"

	tuc, _, cleanup := newTestUpateChecker("", expected, "1.1.3", t)
	defer cleanup()

	invalidateCache(tuc.versionCacheFilename)

	v, err := tuc.getLatestVersion()
	if err != nil {
		t.Fatal(err)
	}

	if v.String() != expected {
		t.Errorf("expected latest version to be %q, was %q", expected, v)
	}

	cv, _, err := tuc.getLatestVersionFromCache()
	if err != nil {
		t.Errorf("could not read version from cache")
	}

	if cv.String() != expected {
		t.Errorf("expected cached version after update to be %q, was %q", expected, v)
	}
}

func TestUpdateChecker_GetLatestVersion_withError(t *testing.T) {
	tuc, _, cleanup := newTestUpateChecker("", "", "", t)
	defer cleanup()

	_, err := tuc.getLatestVersion()
	if err == nil {
		t.Errorf("expected an error, but got nil")
	}
}

func TestUpdateChecker_CheckForUpdate_withUpdateAvailable(t *testing.T) {
	tuc, out, cleanup := newTestUpateChecker("1.1.3", "2.0.0", "", t)
	defer cleanup()

	tuc.Check()

	expected := "Please consider updating the PhraseApp CLI client (1.1.3 < 2.0.0)\n"
	expected += "You can get the latest version from https://phraseapp.com/en/cli.\n"

	if out.String() != expected {
		t.Errorf("expected %q, got %q", expected, out.String())
	}
}

func TestUpdateChecker_CheckForUpdate_withNoUpdateAvailable(t *testing.T) {
	tuc, out, cleanup := newTestUpateChecker("1.1.7", "1.1.7", "", t)
	defer cleanup()

	tuc.Check()

	if out.String() != "" {
		t.Errorf("expected no output, got %q", out.String())
	}
}

func TestUpdateChecker_CheckForUpdate_withDevVersion(t *testing.T) {
	tuc, out, cleanup := newTestUpateChecker("1.1.3-dev", "2.0.0", "", t)
	defer cleanup()

	tuc.Check()

	expected := "You're running a development version (1.1.3-dev) of the PhraseApp client! Latest version is 2.0.0.\n"

	if out.String() != expected {
		t.Errorf("expected %q, got %q", expected, out.String())
	}
}

func newTestUpateChecker(version, latestReleaseVersion, latestCachedVersion string, t *testing.T) (*Checker, *bytes.Buffer, func()) {
	url, stopServer := setupTestServer(latestReleaseVersion)

	file, err := ioutil.TempFile("", "phraseapp_update-check_test")
	if err != nil {
		t.Fatal(err)
	}
	file.Write([]byte(latestCachedVersion))
	file.Close()

	out := &bytes.Buffer{}

	tuc := New(version, file.Name(), url, out)

	return tuc, out, func() {
		stopServer()
		os.Remove(file.Name())
	}
}

func setupTestServer(version string) (string, func()) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", "https://github.com/phrase/phraseapp-client/releases/tag/"+version)
		w.WriteHeader(http.StatusFound)
	}))

	return s.URL, s.Close
}

func invalidateCache(filename string) {
	info, _ := os.Stat(filename)
	os.Chtimes(filename, time.Now(), info.ModTime().Add(-48*time.Hour))
}
