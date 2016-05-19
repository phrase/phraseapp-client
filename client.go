package main

import (
	"crypto/tls"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/phrase/phraseapp-client/Godeps/_workspace/src/github.com/phrase/phraseapp-go/phraseapp"
)

func newClient(creds *phraseapp.Credentials) (*phraseapp.Client, error) {
	c, err := phraseapp.NewClient(creds)
	if err != nil {
		return nil, err
	}
	if os.Getenv("PHRASEAPP_INSECURE_SKIP_VERIFY") == "true" {
		tr := &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		}
		c.Client = http.Client{Transport: tr}
	}
	return c, nil
}
