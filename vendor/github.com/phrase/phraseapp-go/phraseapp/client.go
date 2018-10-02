// Package phraseapp is a library for easier usage of the PhraseApp API
package phraseapp

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/bgentry/speakeasy"
)

// Client is a generic PhraseApp client. It manages a connection to the PhraseApp API
type Client struct {
	http.Client
	Credentials Credentials
	debug       bool
}

// Credentials contains all information to authenticate against phraseapp.com or a custom host.
type Credentials struct {
	Username string `cli:"opt --username -u desc='username used for authentication'"`
	Token    string `cli:"opt --access-token -t desc='access token used for authentication'"`
	TFA      bool   `cli:"opt --tfa desc='use Two-Factor Authentication'"`
	Host     string `cli:"opt --host desc='Host to send Request to'"`
}

// NewClient initializes a new client.
// Uses PHRASEAPP_HOST and PHRASEAPP_ACCESS_TOKEN environment variables for host and access token with specified in environment.
func NewClient(credentials Credentials, debug bool) (*Client, error) {
	client := &Client{
		Credentials: credentials,
		debug:       debug,
	}

	envToken := os.Getenv("PHRASEAPP_ACCESS_TOKEN")
	if envToken != "" && credentials.Token == "" && credentials.Username == "" {
		client.Credentials.Token = envToken
	}

	if client.Credentials.Host == "" {
		envHost := os.Getenv("PHRASEAPP_HOST")
		if envHost != "" {
			client.Credentials.Host = envHost
		} else {
			client.Credentials.Host = "https://api.phraseapp.com"
		}
	}

	return client, nil
}

func (client *Client) authenticate(req *http.Request) error {
	if client.Credentials.Token != "" {
		req.Header.Set("Authorization", "token "+client.Credentials.Token)
	} else if client.Credentials.Username != "" {
		pwd, err := speakeasy.Ask("Password: ")
		if err != nil {
			return err
		}
		req.SetBasicAuth(client.Credentials.Username, pwd)

		if client.Credentials.TFA { // TFA only required for username+password based login.
			token, err := speakeasy.Ask("TFA-Token: ")
			if err != nil {
				return err
			}
			req.Header.Set("X-PhraseApp-OTP", token)
		}
	} else {
		return fmt.Errorf("either username or token must be given")
	}

	req.Header.Set("User-Agent", GetUserAgent())

	return nil
}

func (client *Client) sendRequestPaginated(method, urlPath, contentType string, body io.Reader, expectedStatus, page, perPage int) (io.ReadCloser, error) {
	endpointURL, err := url.Parse(client.Credentials.Host + urlPath)
	if err != nil {
		return nil, err
	}

	addPagination(endpointURL, page, perPage)

	req, err := client.buildRequest(method, endpointURL, body, contentType)
	if err != nil {
		return nil, err
	}

	resp, err := client.send(req, expectedStatus)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func (client *Client) sendRequest(method, urlPath, contentType string, body io.Reader, expectedStatus int) (io.ReadCloser, error) {
	endpointURL, err := url.Parse(client.Credentials.Host + urlPath)
	if err != nil {
		return nil, err
	}

	req, err := client.buildRequest(method, endpointURL, body, contentType)
	if err != nil {
		return nil, err
	}

	resp, err := client.send(req, expectedStatus)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (client *Client) send(req *http.Request, expectedStatus int) (*http.Response, error) {
	err := client.authenticate(req)
	if err != nil {
		return nil, err
	}

	if client.debug {
		b := new(bytes.Buffer)
		err = req.Header.Write(b)
		if err != nil {
			return nil, err
		}

		fmt.Fprintln(os.Stderr, "Header:", b.String())
	}

	resp, err := client.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if client.debug {
		fmt.Fprintf(os.Stderr, "\nResponse HTTP Status Code: %s\n", resp.Status)
	}

	err = handleResponseStatus(resp, expectedStatus)
	if err != nil {
		resp.Body.Close()
	}
	return resp, err
}

func addPagination(u *url.URL, page, perPage int) {
	query := u.Query()
	query.Add("page", strconv.Itoa(page))
	query.Add("per_page", strconv.Itoa(perPage))

	u.RawQuery = query.Encode()
}

func (client *Client) buildRequest(method string, u *url.URL, body io.Reader, contentType string) (*http.Request, error) {
	if client.debug {
		fmt.Fprintln(os.Stderr, "Method:", method)
		fmt.Fprintln(os.Stderr, "URL:", u)

		if body != nil {
			bodyBytes, err := ioutil.ReadAll(body)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error reading body:", err.Error())
			}

			fmt.Fprintln(os.Stderr, "Body:", string(bodyBytes))
			body = bytes.NewReader(bodyBytes)
		}
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	if contentType != "" {
		req.Header.Add("Content-Type", contentType)
	}

	return req, nil
}
