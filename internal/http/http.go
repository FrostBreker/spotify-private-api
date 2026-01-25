// Package http provides HTTP client abstractions for dependency injection.
package http

import (
	"io"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/Danny-Dasilva/CycleTLS/cycletls"
)

// Doer is an interface that wraps the Do method of http.Client.
// This allows for easy mocking in tests.
type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// ClientConfig holds configuration for the HTTP client.
type ClientConfig struct {
	Timeout             time.Duration
	MaxIdleConnsPerHost int
	IdleConnTimeout     time.Duration
}

// DefaultConfig returns the default HTTP client configuration.
func DefaultConfig() ClientConfig {
	return ClientConfig{
		Timeout:             30 * time.Second,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     90 * time.Second,
	}
}

// BrowserClient wraps CycleTLS to provide browser-like TLS fingerprinting.
type BrowserClient struct {
	client  cycletls.CycleTLS
	jar     http.CookieJar
	timeout time.Duration
}

// NewBrowserClient creates an HTTP client that mimics a Chrome browser's TLS fingerprint.
func NewBrowserClient(cfg ClientConfig) *BrowserClient {
	jar, _ := cookiejar.New(nil)
	return &BrowserClient{
		client:  cycletls.Init(),
		jar:     jar,
		timeout: cfg.Timeout,
	}
}

// Do executes an HTTP request using CycleTLS for browser-like TLS fingerprinting.
func (c *BrowserClient) Do(req *http.Request) (*http.Response, error) {
	// Build headers map from request
	headers := make(map[string]string)
	for key, values := range req.Header {
		if len(values) > 0 {
			headers[key] = values[0]
		}
	}

	// Add cookies from jar
	if c.jar != nil {
		for _, cookie := range c.jar.Cookies(req.URL) {
			if existing, ok := headers["Cookie"]; ok {
				headers["Cookie"] = existing + "; " + cookie.String()
			} else {
				headers["Cookie"] = cookie.String()
			}
		}
	}

	// Chrome 120 JA3 fingerprint
	ja3 := "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-13-18-51-45-43-27-17513,29-23-24,0"

	// Execute request
	resp, err := c.client.Do(req.URL.String(), cycletls.Options{
		Headers: headers,
		Ja3:     ja3,
		Body:    "", // Will be set below for POST requests
	}, req.Method)

	if err != nil {
		return nil, err
	}

	// Store cookies from response - CycleTLS returns []*http.Cookie
	if c.jar != nil && len(resp.Cookies) > 0 {
		c.jar.SetCookies(req.URL, resp.Cookies)
	}

	// Convert CycleTLS response to http.Response
	httpResp := &http.Response{
		StatusCode: resp.Status,
		Status:     http.StatusText(resp.Status),
		Header:     make(http.Header),
		Body:       &stringReadCloser{data: resp.Body},
	}

	for key, value := range resp.Headers {
		httpResp.Header.Set(key, value)
	}

	return httpResp, nil
}

// stringReadCloser wraps a string to implement io.ReadCloser.
type stringReadCloser struct {
	data string
	pos  int
}

func (r *stringReadCloser) Read(p []byte) (n int, err error) {
	if r.pos >= len(r.data) {
		return 0, io.EOF
	}
	n = copy(p, r.data[r.pos:])
	r.pos += n
	return n, nil
}

func (r *stringReadCloser) Close() error {
	return nil
}

// NewDefaultClient creates a browser-like HTTP client.
func NewDefaultClient() Doer {
	return NewBrowserClient(DefaultConfig())
}
