package phraseapp

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ErrorResponse struct {
	Message string
}

func (err *ErrorResponse) Error() string {
	return err.Message
}

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

type ValidationErrorMessage struct {
	Resource string
	Field    string
	Message  string
}

func (msg *ValidationErrorMessage) String() string {
	return fmt.Sprintf("\t[%s:%s] %s", msg.Resource, msg.Field, msg.Message)
}

type RateLimitingError struct {
	Limit     int
	Remaining int
	Reset     time.Time
}

func NewRateLimitError(resp *http.Response) (*RateLimitingError, error) {
	var err error
	re := new(RateLimitingError)

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
	return fmt.Sprintf("Rate limit exceeded: from %d requests %d are remaning (reset in %d seconds)", rle.Limit, rle.Remaining, int64(rle.Reset.Sub(time.Now()).Seconds()))
}
