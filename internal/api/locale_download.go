package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/phrase/phraseapp-go/phraseapp"
)

func LocaleDownload(client *phraseapp.Client, project_id, id string, params *phraseapp.LocaleDownloadParams) (*http.Response, error) {
	urlPath := fmt.Sprintf("/v2/projects/%s/locales/%s/download", url.QueryEscape(project_id), url.QueryEscape(id))
	endpointURL, err := url.Parse(client.Credentials.Host + urlPath)
	if err != nil {
		return nil, err
	}

	paramsBuf := bytes.NewBuffer(nil)
	err = json.NewEncoder(paramsBuf).Encode(&params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", endpointURL.String(), paramsBuf)
	if err != nil {
		return nil, err
	}

	err = client.Authenticate(req)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Client.Do(req)
	if err != nil {
		return nil, err
	}

	err = phraseapp.HandleResponseStatus(resp, 200)
	if err != nil {
		resp.Body.Close()
		return nil, err
	}

	return resp, err
}
