# Spotify Private API Go Package

[![Go Version](https://img.shields.io/badge/Go-1.24+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE.MD)
[![Tests](https://github.com/FrostBreker/spotify-private-api/actions/workflows/tests.yml/badge.svg)](https://github.com/FrostBreker/spotify-private-api/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/FrostBreker/spotify-private-api)](https://goreportcard.com/report/github.com/FrostBreker/spotify-private-api)

A Go package for accessing Spotify data by scraping public pages. No authentication or API tokens required.

## Features

- 🎵 **Track Information** - Track name, duration, album, artists, and cover art
- 🎤 **Artist Profiles** - Artist stats, biography, followers, and monthly listeners
- 💿 **Album Details** - Album info with track listings and artwork
- 🔓 **No Authentication** - Works without any API tokens or OAuth
- 🌐 **Browser-Like Requests** - Uses TLS fingerprinting to mimic real browsers
- ⚡ **Context Support** - All operations support cancellation and timeouts
- 🧪 **Fully Testable** - Mock support for comprehensive unit testing
- 🛡️ **Type-Safe Errors** - Custom error types for precise error handling

## Installation

```bash
go get github.com/FrostBreker/spotify-private-api
```

**Requirements:** Go 1.24 or higher

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    spotify "github.com/FrostBreker/spotify-private-api"
)

func main() {
    client := spotify.NewClient()
    ctx := context.Background()

    // Fetch a track
    track, err := client.FetchTrack(ctx, "4cOdK2wGLETKBW3PvgPWqT")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Track: %s\n", track.Data.TrackUnion.Name)
    fmt.Printf("Artist: %s\n", track.Data.TrackUnion.FirstArtist.Items[0].Profile.Name)
    fmt.Printf("Album: %s\n", track.Data.TrackUnion.AlbumOfTrack.Name)
}
```

## Usage Examples

### Track Information

```go
client := spotify.NewClient()
ctx := context.Background()

track, err := client.FetchTrack(ctx, "4cOdK2wGLETKBW3PvgPWqT")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Name: %s\n", track.Data.TrackUnion.Name)
fmt.Printf("URI: %s\n", track.Data.TrackUnion.URI)
fmt.Printf("Duration: %dms\n", track.Data.TrackUnion.Duration.TotalMilliseconds)
fmt.Printf("Album: %s\n", track.Data.TrackUnion.AlbumOfTrack.Name)

if len(track.Data.TrackUnion.FirstArtist.Items) > 0 {
    fmt.Printf("Artist: %s\n", track.Data.TrackUnion.FirstArtist.Items[0].Profile.Name)
}
```

### Artist Information

```go
client := spotify.NewClient()
ctx := context.Background()

artist, err := client.FetchArtist(ctx, "4tZwfgrHOc3mvqYlEYSvVi")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Artist: %s\n", artist.Data.ArtistUnion.Profile.Name)
fmt.Printf("Followers: %d\n", artist.Data.ArtistUnion.Stats.Followers)
fmt.Printf("Monthly Listeners: %d\n", artist.Data.ArtistUnion.Stats.MonthlyListeners)

// Access artist's avatar
if len(artist.Data.ArtistUnion.Visuals.AvatarImage.Sources) > 0 {
    fmt.Printf("Avatar URL: %s\n", artist.Data.ArtistUnion.Visuals.AvatarImage.Sources[0].URL)
}
```

### Album Information

```go
client := spotify.NewClient()
ctx := context.Background()

album, err := client.FetchAlbum(ctx, "2noRn2Aes5aoNVsU6iWThc")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Album: %s\n", album.Data.AlbumUnion.Name)
fmt.Printf("URI: %s\n", album.Data.AlbumUnion.URI)
fmt.Printf("Total Tracks: %d\n", album.Data.AlbumUnion.Tracks.TotalCount)

if len(album.Data.AlbumUnion.Artists.Items) > 0 {
    fmt.Printf("Artist: %s\n", album.Data.AlbumUnion.Artists.Items[0].Profile.Name)
}
```

## Configuration Options

```go
// Enable debug logging
client := spotify.NewClient(
    spotify.WithDebug(true),
)

// Custom logger (uses log/slog)
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
client := spotify.NewClient(
    spotify.WithLogger(logger),
)

// Custom HTTP client for testing
client := spotify.NewClient(
    spotify.WithHTTPDoer(mockDoer),
)
```

## API Reference

### Client Methods

| Method          | Parameters        | Description              | Returns            |
| --------------- | ----------------- | ------------------------ | ------------------ |
| `FetchTrack()`  | `ctx`, `trackID`  | Get track information    | `*TrackResponse`   |
| `FetchArtist()` | `ctx`, `artistID` | Get artist profile       | `*ArtistResponse`  |
| `FetchAlbum()`  | `ctx`, `albumID`  | Get album details        | `*AlbumResponse`   |

### Configuration Options

| Option                 | Description                        |
| ---------------------- | ---------------------------------- |
| `WithDebug(bool)`      | Enable/disable debug logging       |
| `WithHTTPDoer(Doer)`   | Use custom HTTP doer (for testing) |
| `WithLogger(*slog.Logger)` | Use custom structured logger   |

## Error Handling

```go
import "github.com/FrostBreker/spotify-private-api/errors"

track, err := client.FetchTrack(ctx, trackID)
if err != nil {
    if errors.IsAPIError(err) {
        fmt.Println("API returned an error")
    }
    if errors.IsParseError(err) {
        fmt.Println("Failed to parse response")
    }
    if err == errors.ErrInvalidID {
        fmt.Println("Invalid Spotify ID provided")
    }
}
```

### Error Types

| Error Type     | Description                            |
| -------------- | -------------------------------------- |
| `APIError`     | Spotify returned an error response     |
| `ParseError`   | JSON parsing or processing error       |
| `RequestError` | HTTP request failed                    |

### Sentinel Errors

| Error              | Description                 |
| ------------------ | --------------------------- |
| `ErrInvalidID`     | Invalid Spotify ID provided |
| `ErrEmptyResponse` | Empty response received     |

## Project Structure

```
spotify-private-api/
├── client.go          # Main client implementation
├── track.go           # Track operations
├── artist.go          # Artist operations
├── album.go           # Album operations
├── doc.go             # Package documentation
├── errors/
│   └── errors.go      # Custom error types
├── internal/
│   └── http/
│       └── http.go    # Browser-like HTTP client
├── models/
│   ├── track.go       # Track response types
│   ├── artist.go      # Artist response types
│   └── album.go       # Album response types
└── example/
    └── main.go        # Usage examples
```

## Testing

```go
// Create a mock HTTP client
type mockHTTPClient struct {
    doFunc func(req *http.Request) (*http.Response, error)
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
    return m.doFunc(req)
}

// Use in tests
client := spotify.NewClient(
    spotify.WithHTTPDoer(&mockHTTPClient{
        doFunc: func(req *http.Request) (*http.Response, error) {
            return &http.Response{
                StatusCode: 200,
                Body:       io.NopCloser(strings.NewReader(mockHTML)),
            }, nil
        },
    }),
)
```

Run tests:

```bash
go test ./...
go test -v ./...
go test -cover ./...
```

## How It Works

This package scrapes Spotify's public web pages to extract embedded data. Spotify embeds JSON data in a `<script id="initialState">` tag as base64-encoded content. The package:

1. Makes HTTP requests with browser-like TLS fingerprints (using CycleTLS)
2. Parses the HTML response to find the embedded data
3. Decodes and unmarshals the JSON into typed Go structs

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

See [CONTRIBUTING.md](CONTRIBUTING.md) for detailed guidelines.

## License

This project is licensed under the MIT License - see the [LICENSE.MD](LICENSE.MD) file for details.

## Disclaimer

⚠️ **Important**: This package scrapes Spotify's public web pages. Use responsibly.

- This package is for educational and research purposes only
- The page structure may change without notice, potentially breaking this package
- Excessive use may result in rate limiting or IP blocks
- Not affiliated with or endorsed by Spotify

---

**Happy Listening! 🎧**
