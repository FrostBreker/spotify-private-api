# Spotify Private API

This package provides a simple way to interact with Spotify's private API. It includes methods to fetch track information and play count.

## Installation

```bash
go get github.com/FrostBreker/spotify-private-api@v1.0.2
```

## Usage

```go
import "github.com/FrostBreker/spotify-private-api"

client := spotifyprivateapi.NewClient(spotifyprivateapi.Client{
		Debug: true,
})

// Fetch track information
track, err := client.FetchTrack("trackId")
if err != nil {
    log.Fatal(err)
}

// Fetch track play count
playCount, err := client.FetchTrackPlayCount("trackId")
if err != nil {
    log.Fatal(err)
}
```

## Documentation

### FetchTrack(trackId string) (responseTypes.TrackResponse, error)

FetchTrack sends a GET request to Spotify's private API to fetch information about a track. The trackId parameter is the Spotify ID of the track. It returns a TrackResponse object and an error.

### FetchTrackPlayCount(trackId string) (responseTypes.TrackPlayCount, error)

FetchTrackPlayCount first calls FetchTrack to get the track information, then it extracts the play count from the response. The trackId parameter is the Spotify ID of the track. It returns a TrackPlayCount object and an error.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)

**Note:** This package uses Spotify's private API and is not officially supported by Spotify. Use at your own risk.
