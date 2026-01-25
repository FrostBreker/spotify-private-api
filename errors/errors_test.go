package errors

import (
	"errors"
	"testing"
)

func TestAPIError(t *testing.T) {
	t.Run("Error message", func(t *testing.T) {
		err := NewAPIError(404, "not found")
		expected := "spotify API error (status 404): not found"
		if err.Error() != expected {
			t.Errorf("Error() = %v, want %v", err.Error(), expected)
		}
	})

	t.Run("IsAPIError", func(t *testing.T) {
		err := NewAPIError(400, "bad request")
		if !IsAPIError(err) {
			t.Error("IsAPIError() should return true for APIError")
		}

		regularErr := errors.New("regular error")
		if IsAPIError(regularErr) {
			t.Error("IsAPIError() should return false for regular error")
		}
	})
}

func TestParseError(t *testing.T) {
	t.Run("Error message", func(t *testing.T) {
		innerErr := errors.New("invalid json")
		err := NewParseError("track", innerErr)
		if err.Error() != "failed to parse track response: invalid json" {
			t.Errorf("Error() = %v", err.Error())
		}
	})

	t.Run("Unwrap", func(t *testing.T) {
		innerErr := errors.New("inner error")
		err := NewParseError("track", innerErr)
		if !errors.Is(err, innerErr) {
			t.Error("Unwrap() should return inner error")
		}
	})

	t.Run("IsParseError", func(t *testing.T) {
		err := NewParseError("track", nil)
		if !IsParseError(err) {
			t.Error("IsParseError() should return true for ParseError")
		}
	})
}

func TestRequestError(t *testing.T) {
	t.Run("Error message", func(t *testing.T) {
		innerErr := errors.New("connection refused")
		err := NewRequestError("fetch", "https://example.com", innerErr)
		expected := "request error during fetch to https://example.com: connection refused"
		if err.Error() != expected {
			t.Errorf("Error() = %v, want %v", err.Error(), expected)
		}
	})

	t.Run("Unwrap", func(t *testing.T) {
		innerErr := errors.New("inner error")
		err := NewRequestError("operation", "url", innerErr)
		if !errors.Is(err, innerErr) {
			t.Error("Unwrap() should return inner error")
		}
	})
}

func TestSentinelErrors(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{"ErrInvalidID", ErrInvalidID, "invalid spotify ID"},
		{"ErrEmptyResponse", ErrEmptyResponse, "empty response from API"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err.Error() != tt.want {
				t.Errorf("%s.Error() = %v, want %v", tt.name, tt.err.Error(), tt.want)
			}
		})
	}
}
