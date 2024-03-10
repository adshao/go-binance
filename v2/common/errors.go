package common

import (
	"fmt"
)

// APIError define API error when response status is 4xx or 5xx
type APIError struct {
	StatusCode int    `json:"-"`
	Code       int64  `json:"code"`
	Message    string `json:"msg"`
}

// Error return error code and message
func (e APIError) Error() string {
	return fmt.Sprintf("<APIError> status_code=%d, code=%d, msg=%s", e.StatusCode, e.Code, e.Message)
}

// IsAPIError check if e is an API error
func IsAPIError(e error) bool {
	_, ok := e.(*APIError)
	return ok
}
