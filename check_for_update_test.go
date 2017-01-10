package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestGetLatestVersionFromURL(t *testing.T) {
	clearCache()
	defer clearCache()

	expected := "1.1.3"

	cleanupTestServer := setupTestServer(expected)
	defer cleanupTestServer()

	v, err := getLatestVersionFromURL()
	if err != nil {
		t.Fatal(err)
	}

	if v.String() != expected {
		t.Errorf("expected latest version to be %q, was %q", expected, v)
	}
}

func TestGetLatestVersionFromCache(t *testing.T) {
	expected := "1.1.3"

	setupTestCache(expected)
	defer clearCache()

	v, _, err := getLatestVersionFromCache()
	if err != nil {
		t.Errorf("could not read version from cache!")
	}

	if v.String() != expected {
		t.Errorf("expected latest version to be %q, was %q", expected, v)
	}
}

func TestGetLatestVersionWithCache(t *testing.T) {
	expected := "1.1.3"

	setupTestCache(expected)
	defer clearCache()

	cleanupTestServer := setupTestServer("2.0.0")
	defer cleanupTestServer()

	v, err := getLatestVersion()
	if err != nil {
		t.Fatal(err)
	}

	if v.String() != expected {
		t.Errorf("expected latest version to be %q, was %q", expected, v)
	}
}

func TestGetLatestVersionWithoutCache(t *testing.T) {
	clearCache()
	defer clearCache()

	expected := "2.0.0"

	cleanupTestServer := setupTestServer(expected)
	defer cleanupTestServer()

	v, err := getLatestVersion()
	if err != nil {
		t.Fatal(err)
	}

	if v.String() != expected {
		t.Errorf("expected latest version to be %q, was %q", expected, v)
	}

	cv, _, err := getLatestVersionFromCache()
	if err != nil {
		t.Errorf("could not read version from cache!")
	}

	if cv.String() != expected {
		t.Errorf("expected cached version after update to be %q, was %q", expected, v)
	}
}

func TestGetLatestVersionWithInvalidCache(t *testing.T) {
	setupTestCache("1.1.3")
	invalidateCache()
	defer clearCache()

	expected := "2.0.0"

	cleanupTestServer := setupTestServer(expected)
	defer cleanupTestServer()

	v, err := getLatestVersion()
	if err != nil {
		t.Fatal(err)
	}

	if v.String() != expected {
		t.Errorf("expected latest version to be %q, was %q", expected, v)
	}

	cv, _, err := getLatestVersionFromCache()
	if err != nil {
		t.Errorf("could not read version from cache!")
	}

	if cv.String() != expected {
		t.Errorf("expected cached version after update to be %q, was %q", expected, v)
	}
}

func ExampleCheckForUpdate_withUpdateAvailable() {
	clearCache()
	defer clearCache()

	resetVersion := setupFakeVersion("1.1.3")
	defer resetVersion()

	expected := "2.0.0"

	cleanupTestServer := setupTestServer(expected)
	defer cleanupTestServer()

	CheckForUpdate()

	// Output:
	// Please consider updating the PhraseApp CLI client (1.1.3 < 2.0.0)
	// You can get the latest version from https://phraseapp.com/en/cli.
}

func ExampleCheckForUpdate_withNoUpdateAvailable() {
	clearCache()
	defer clearCache()

	latest := "1.7.0"

	resetVersion := setupFakeVersion(latest)
	defer resetVersion()

	cleanupTestServer := setupTestServer(latest)
	defer cleanupTestServer()

	CheckForUpdate()

	// Output:
}

func ExampleCheckForUpdate_withDevVersion() {
	clearCache()
	defer clearCache()

	resetVersion := setupFakeVersion("1.1.3-dev")
	defer resetVersion()

	cleanupTestServer := setupTestServer("2.0.0")
	defer cleanupTestServer()

	CheckForUpdate()

	// Output: You're running a development version (1.1.3-dev) of the PhraseApp client! Latest version is 2.0.0.
}

func setupTestServer(version string) func() {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", "https://github.com/phrase/phraseapp-client/releases/tag/"+version)
		w.WriteHeader(http.StatusFound)
	}))

	old := releasesURL
	releasesURL = s.URL

	return func() {
		s.Close()
		releasesURL = old
	}
}

func setupTestCache(version string) {
	ioutil.WriteFile(versionCacheFilename, []byte(version), 0600)
}

func invalidateCache() {
	info, _ := os.Stat(versionCacheFilename)
	os.Chtimes(versionCacheFilename, time.Now(), info.ModTime().Add(-48*time.Hour))
}

func clearCache() {
	os.Remove(versionCacheFilename)
}

func setupFakeVersion(version string) func() {
	tmp := PHRASEAPP_CLIENT_VERSION
	PHRASEAPP_CLIENT_VERSION = version

	return func() {
		PHRASEAPP_CLIENT_VERSION = tmp
	}
}
