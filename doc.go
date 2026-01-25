// Package spotifyprivateapi provides a client for accessing Spotify data.
//
// This package scrapes Spotify's public web pages to fetch track, artist,
// and album information without requiring any authentication tokens.
//
// Basic usage:
//
//	client := spotifyprivateapi.NewClient()
//	track, err := client.FetchTrack(ctx, "trackID")
//
// For more examples, see the example directory.
package spotifyprivateapi

// Package structure:
//   - github.com/FrostBreker/spotify-private-api (main client)
//   - github.com/FrostBreker/spotify-private-api/models (response types)
//   - github.com/FrostBreker/spotify-private-api/errors (error types)
