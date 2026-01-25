// Package spotifyprivateapi provides a client for interacting with Spotify's private API.
package spotifyprivateapi

import (
	"context"
	"log/slog"
	"os"

	internalhttp "github.com/FrostBreker/spotify-private-api/internal/http"
)

// Client is the main client for interacting with the Spotify API.
type Client struct {
	httpClient internalhttp.Doer
	logger     *slog.Logger
	debug      bool
}

// Option is a functional option for configuring a Client.
type Option func(*Client)

// WithHTTPDoer sets a custom HTTP doer interface for the Client.
// This is useful for testing with mock HTTP clients.
func WithHTTPDoer(doer internalhttp.Doer) Option {
	return func(c *Client) {
		c.httpClient = doer
	}
}

// WithLogger sets a custom logger for the Client.
func WithLogger(logger *slog.Logger) Option {
	return func(c *Client) {
		c.logger = logger
	}
}

// WithDebug enables debug logging.
func WithDebug(debug bool) Option {
	return func(c *Client) {
		c.debug = debug
	}
}

// NewClient creates a new Spotify client with the given options.
func NewClient(opts ...Option) *Client {
	c := &Client{
		httpClient: internalhttp.NewDefaultClient(),
		logger:     slog.New(slog.NewTextHandler(os.Stdout, nil)),
		debug:      false,
	}

	for _, opt := range opts {
		opt(c)
	}

	if c.debug {
		c.logger.Info("Spotify client initialized")
	}

	return c
}

// log logs a message if debug mode is enabled.
func (c *Client) log(level slog.Level, msg string, args ...any) {
	if c.debug {
		c.logger.Log(context.Background(), level, msg, args...)
	}
}

// logInfo logs an info message if debug mode is enabled.
func (c *Client) logInfo(msg string, args ...any) {
	c.log(slog.LevelInfo, msg, args...)
}

// logError logs an error message if debug mode is enabled.
func (c *Client) logError(msg string, args ...any) {
	c.log(slog.LevelError, msg, args...)
}
