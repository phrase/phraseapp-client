package main

import (
	"crypto/tls"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/phrase/phraseapp-go/phraseapp"
)

func newClient(creds phraseapp.Credentials, debug bool, cache bool) (*phraseapp.Client, error) {
	client, err := phraseapp.NewClient(creds, debug)
	if err != nil {
		return nil, err
	}

	if cache {
		err := client.EnableCaching(true)
		if err != nil {
			return client, err
		}
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
		client.Client = http.Client{Transport: tr}
	}

	return client, nil
}
