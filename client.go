package main

import (
	"crypto/tls"
	"net/http"
	"os"

	"github.com/phrase/phraseapp-client/Godeps/_workspace/src/github.com/phrase/phraseapp-go/phraseapp"
)

func newClient(creds *phraseapp.Credentials) (*phraseapp.Client, error) {
	c, err := phraseapp.NewClient(creds)
	if err != nil {
		return nil, err
	}
	if os.Getenv("PHRASEAPP_INSECURE_SKIP_VERIFY") == "true" {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		c.Client = http.Client{Transport: tr}
	}
	return c, nil
}
