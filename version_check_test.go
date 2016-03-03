package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVersionCheck(t *testing.T) {
	defer expectReleaseVersion("1.1.3", 302)()

	v, err := getLatestReleaseVersion()
	if err != nil {
		t.Fatal(err)
	}
	if x, v := "1.1.3", v; x != v {
		t.Errorf("expected release version to be %q, was %q", x, v)
	}
}

func TestValidateVersionWithErr(t *testing.T) {
	oldVersion := PHRASEAPP_CLIENT_VERSION
	currentVersion := "1.1.11"

	defer expectReleaseVersion(currentVersion, 302)()

	errorCases := []string{"1.0.0", "1.1.10", "1.0.11", "1.1.10-dev"}
	for _, version := range errorCases {
		PHRASEAPP_CLIENT_VERSION = version
		err := validateVersionWithErr()
		if err == nil {
			t.Errorf("expected an error but got none for %s >= %s", version, currentVersion)
		}
	}

	noErrorCases := []string{"1.1.11", "1.1.12", "1.2.10", "2.0.0"}
	for _, version := range noErrorCases {
		PHRASEAPP_CLIENT_VERSION = version
		err := validateVersionWithErr()
		if err != nil {
			t.Errorf("expected no error but got %s", err)
		}
	}

	PHRASEAPP_CLIENT_VERSION = oldVersion
}

func expectReleaseVersion(version string, statusCode int) func() {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", "https://github.com/phrase/phraseapp-client/releases/tag/"+version)
		w.WriteHeader(statusCode)
	}
	s := httptest.NewServer(handler)
	old := releaseURL
	releaseURL = s.URL
	return func() {
		s.Close()
		releaseURL = old
	}
}
