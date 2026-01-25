// Package errors provides custom error types for the spotify-private-api package.
package errors

import (
	"errors"
	"fmt"
)

// Sentinel errors for common error conditions.
var (
	ErrInvalidID     = errors.New("invalid spotify ID")
	ErrEmptyResponse = errors.New("empty response from API")
)

// APIError represents an error returned by the Spotify API.
type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("spotify API error (status %d): %s", e.StatusCode, e.Message)
}

// NewAPIError creates a new APIError with the given status code and message.
func NewAPIError(statusCode int, message string) *APIError {
	return &APIError{StatusCode: statusCode, Message: message}
}

// ParseError represents a JSON parsing error.
type ParseError struct {
	Type string
	Err  error
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("failed to parse %s response: %v", e.Type, e.Err)
}

func (e *ParseError) Unwrap() error {
	return e.Err
}

// NewParseError creates a new ParseError.
func NewParseError(responseType string, err error) *ParseError {
	return &ParseError{Type: responseType, Err: err}
}

// RequestError represents an HTTP request error.
type RequestError struct {
	Operation string
	URL       string
	Err       error
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("request error during %s to %s: %v", e.Operation, e.URL, e.Err)
}

func (e *RequestError) Unwrap() error {
	return e.Err
}

// NewRequestError creates a new RequestError.
func NewRequestError(operation, url string, err error) *RequestError {
	return &RequestError{Operation: operation, URL: url, Err: err}
}

// IsAPIError checks if the error is an APIError.
func IsAPIError(err error) bool {
	var apiErr *APIError
	return errors.As(err, &apiErr)
}

// IsParseError checks if the error is a ParseError.
func IsParseError(err error) bool {
	var parseErr *ParseError
	return errors.As(err, &parseErr)
}
