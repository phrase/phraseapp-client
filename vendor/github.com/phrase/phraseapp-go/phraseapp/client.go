package phraseapp

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"os"

	"github.com/bgentry/speakeasy"
)

var Debug bool

func EnableDebug() {
	Debug = true
}

type Client struct {
	http.Client
	Credentials *Credentials
}

type Credentials struct {
	Username string `cli:"opt --username -u desc='username used for authentication'"`
	Token    string `cli:"opt --access-token -t desc='access token used for authentication'"`
	TFA      bool   `cli:"opt --tfa desc='use Two-Factor Authentication'"`
	Host     string `cli:"opt --host desc='Host to send Request to'"`
	Debug    bool   `cli:"opt --verbose -v desc='Verbose output'"`
}

func NewClient(credentials *Credentials) (*Client, error) {
	client := &Client{Credentials: credentials}

	envToken := os.Getenv("PHRASEAPP_ACCESS_TOKEN")
	if envToken != "" && credentials.Token == "" && credentials.Username == "" {
		client.Credentials.Token = envToken
	}

	if credentials.Debug == true {
		EnableDebug()
	}

	if credentials.Host == "" {
		client.Credentials.Host = "https://api.phraseapp.com"
	}

	return client, nil
}

func (client *Client) authenticate(req *http.Request) error {
	if client.Credentials == nil {
		return fmt.Errorf("no auth handler registered")
	}

	if err := client.Credentials.validate(); err != nil {
		return err
	}

	req.Header.Set("User-Agent", GetUserAgent())
	switch {
	case client.Credentials.Token != "":
		req.Header.Set("Authorization", "token "+client.Credentials.Token)
	case client.Credentials.Username != "":
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
	}

	return nil
}

func (ah *Credentials) validate() error {
	switch {
	case ah.Username == "" && ah.Token == "":
		return fmt.Errorf("either username or token must be given")
	default:
		return nil
	}
}

func (client *Client) sendRequestPaginated(method, rawurl, ctype string, r io.Reader, status, page, perPage int) (io.ReadCloser, error) {
	endpointUrl := client.Credentials.Host + rawurl
	u, err := url.Parse(endpointUrl)
	if err != nil {
		return nil, err
	}

	query := u.Query()
	query.Add("page", strconv.Itoa(page))
	query.Add("per_page", strconv.Itoa(perPage))

	u.RawQuery = query.Encode()

	if Debug {
		fmt.Fprintln(os.Stderr, method, u)
		if r != nil {
			bytes, err := ioutil.ReadAll(r)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
			}
			str := string(bytes)
			fmt.Fprintln(os.Stderr, str)
			r = strings.NewReader(str)
		}
	}

	req, err := http.NewRequest(method, u.String(), r)
	if err != nil {
		return nil, err
	}

	if ctype != "" {
		req.Header.Add("Content-Type", ctype)
	}

	resp, err := client.send(req, status)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func (client *Client) sendRequest(method, url, ctype string, r io.Reader, status int) (io.ReadCloser, error) {
	endpointUrl := client.Credentials.Host + url
	if Debug {
		fmt.Fprintln(os.Stderr, method, url)
		if r != nil {
			bytes, err := ioutil.ReadAll(r)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
			}
			str := string(bytes)
			fmt.Fprintln(os.Stderr, str)
			r = strings.NewReader(str)
		}
	}
	req, err := http.NewRequest(method, endpointUrl, r)
	if err != nil {
		return nil, err
	}

	if ctype != "" {
		req.Header.Add("Content-Type", ctype)
	}

	resp, err := client.send(req, status)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (client *Client) send(req *http.Request, status int) (*http.Response, error) {
	err := client.authenticate(req)
	if err != nil {
		return nil, err
	}

	if Debug {
		b := new(bytes.Buffer)
		err = req.Header.Write(b)
		if err != nil {
			return nil, err
		}

		bytes, err := ioutil.ReadAll(b)
		if err != nil {
			return nil, err
		}
		fmt.Fprintln(os.Stderr, string(bytes))

		r := req.Body
		if r != nil {
			var by []byte
			_, err = r.Read(by)
			fmt.Fprintln(os.Stderr, string(by))
		}
	}
	resp, err := client.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if Debug {
		fmt.Fprintf(os.Stderr, "\nResponse HTTP Status Code: %s\n", resp.Status)
	}

	err = handleResponseStatus(resp, status)
	if err != nil {
		resp.Body.Close()
	}
	return resp, err
}
