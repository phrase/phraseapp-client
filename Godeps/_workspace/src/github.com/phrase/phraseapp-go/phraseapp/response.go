package phraseapp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const docsURL = `http://docs.phraseapp.com/api/v2/`

func further() string {
	return fmt.Sprintf("\nFor further information see:\n%s", docsURL)
}

func handleResponseStatus(resp *http.Response, expectedStatus int) error {
	switch resp.StatusCode {
	case expectedStatus:
		return nil
	case 400:
		e := new(ErrorResponse)
		err := json.NewDecoder(resp.Body).Decode(&e)
		if err != nil {
			return err
		}
		return e
	case 401:
		return fmt.Errorf("401 - %s\nThe credentials you provided are invalid.%s", http.StatusText(resp.StatusCode), further())
	case 403:
		return fmt.Errorf("403 - %s\nYou are not authorized to perform the requested action on the requested resource. Check if your provided access_token has the correct scope.%s", http.StatusText(resp.StatusCode), further())
	case 404:
		return fmt.Errorf("404 - Resource Not Found\nThe resource you requested or referenced resources you required do either not exist or you do not have the authorization to request this resource.")
	case 422:
		e := new(ValidationErrorResponse)
		err := json.NewDecoder(resp.Body).Decode(&e)
		if err != nil {
			return err
		}
		return e
	case 429:
		e, err := NewRateLimitError(resp)
		if err != nil {
			return err
		}
		return e
	default:
		return fmt.Errorf("Unexpected HTTP Status Code (%d %s) received; expected %d %s.%s", resp.StatusCode, http.StatusText(resp.StatusCode), expectedStatus, http.StatusText(expectedStatus), further())
	}
}
