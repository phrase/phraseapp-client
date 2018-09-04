package phraseapp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func IsErrNotFound(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(ErrNotFound)
	return ok
}

// ErrNotFound represents an error for requests of non existing resources
type ErrNotFound struct {
	Message string
}

func (e ErrNotFound) Error() string {
	return e.Message
}

type ErrorResponse struct {
	Message string
}

func (err *ErrorResponse) Error() string {
	return err.Message
}

// ValidationErrorResponse represents the response for a failed validation of content
type ValidationErrorResponse struct {
	ErrorResponse

	Errors []ValidationErrorMessage
}

func (err *ValidationErrorResponse) Error() string {
	msgs := make([]string, len(err.Errors))
	for i := range err.Errors {
		msgs[i] = err.Errors[i].String()
	}
	return fmt.Sprintf("%s\n%s", err.Message, strings.Join(msgs, "\n"))
}

// ValidationErrorMessage represents an error for a failed validation of content
type ValidationErrorMessage struct {
	Resource string
	Field    string
	Message  string
}

func (msg *ValidationErrorMessage) String() string {
	return fmt.Sprintf("\t[%s:%s] %s", msg.Resource, msg.Field, msg.Message)
}

// RateLimitingError is returned when hitting the API rate limit
type RateLimitingError struct {
	Limit           int
	Remaining       int
	Reset           time.Time
	TooManyRequests bool
}

const errorConcurrencyLimit = "Concurrency limit exceeded"

func NewRateLimitError(resp *http.Response) (*RateLimitingError, error) {
	var err error
	re := new(RateLimitingError)
	b, err := ioutil.ReadAll(resp.Body)
	if err == nil && strings.TrimSpace(string(b)) == errorConcurrencyLimit {
		re.TooManyRequests = true
	}

	limit := resp.Header.Get("X-Rate-Limit-Limit")
	re.Limit, err = strconv.Atoi(limit)
	if err != nil {
		return nil, err
	}

	remaining := resp.Header.Get("X-Rate-Limit-Remaining")
	re.Remaining, err = strconv.Atoi(remaining)
	if err != nil {
		return nil, err
	}

	reset := resp.Header.Get("X-Rate-Limit-Reset")
	sinceEpoch, err := strconv.ParseInt(reset, 10, 64)
	if err != nil {
		return nil, err
	}
	re.Reset = time.Unix(sinceEpoch, 0)

	return re, nil
}

func (rle *RateLimitingError) Error() string {
	if rle.TooManyRequests {
		return fmt.Sprintf("Rate limit exceeded: too many parallel requests")
	}
	return fmt.Sprintf("Rate limit exceeded: from %d requests %d are remaining (reset in %d seconds)", rle.Limit, rle.Remaining, int64(rle.Reset.Sub(time.Now()).Seconds()))
}
