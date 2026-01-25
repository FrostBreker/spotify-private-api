package spotifyprivateapi

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

type mockHTTPClient struct {
	doFunc func(req *http.Request) (*http.Response, error)
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.doFunc(req)
}

func mockResponse(statusCode int, body string) *http.Response {
	return &http.Response{
		StatusCode: statusCode,
		Body:       io.NopCloser(strings.NewReader(body)),
		Header:     make(http.Header),
	}
}

func newTestClient(httpClient *mockHTTPClient) *Client {
	return NewClient(
		WithHTTPDoer(httpClient),
		WithDebug(false),
		WithLogger(nil),
	)
}

func TestNewClient(t *testing.T) {
	tests := []struct {
		name string
		opts []Option
	}{
		{
			name: "default client",
			opts: nil,
		},
		{
			name: "with debug",
			opts: []Option{WithDebug(true)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(tt.opts...)
			if client == nil {
				t.Error("NewClient returned nil")
			}
		})
	}
}
