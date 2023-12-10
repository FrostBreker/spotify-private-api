# Spotify Private API

This package provides a simple way to interact with Spotify's private API. It includes methods to fetch track information and play count.

## Installation

```bash
go get github.com/FrostBreker/spotify-private-api
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

// Fetch artist information
artist, err := client.FetchArtist("artistId")
if err != nil {
    log.Fatal(err)
}

// Fetch album information
album, err := client.FetchAlbum("albumId")
if err != nil {
    log.Fatal(err)
}

// Search in spotify
searchResults, err := client.Search("query")
if err != nil {
    log.Fatal(err)
}

// Search for top results in spotify
topTracks, err := client.SearchTopResult("query")
if err != nil {
    log.Fatal(err)
}
```

## Documentation

### FetchTrack(trackId string) (responseTypes.TrackResponse, error)

FetchTrack sends a GET request to Spotify's private API to fetch information about a track. The trackId parameter is the Spotify ID of the track. It returns a TrackResponse object and an error.

### FetchTrackPlayCount(trackId string) (responseTypes.TrackPlayCount, error)

FetchTrackPlayCount first calls FetchTrack to get the track information, then it extracts the play count from the response. The trackId parameter is the Spotify ID of the track. It returns a TrackPlayCount object and an error.

### FetchArtist(artistId string) (responseTypes.Artist, error)

FetchArtist fetches information about an artist from Spotify's private API. It takes an artistId parameter, which is the Spotify ID of the artist. It returns an Artist object and an error.

### FetchAlbum(albumId string) (responseTypes.Album, error)

FetchAlbum fetches information about an album from Spotify's private API. It takes an albumId parameter, which is the Spotify ID of the album. It returns an AlbumResponseType object and an error.

### Search(query string) (responseTypes.SearchResponseType, error)

Search sends a GET request to Spotify's private API to search for tracks, albums, artists and playlists. It takes a query parameter, which is the search query. It returns a SearchResponseType object and an error.

### SearchTopResult(query string) (responseTypes.SearchTopResultResponseType, error)

SearchTopResult first calls Search to get the search results, then it extracts the top results from the response. It takes a query parameter, which is the search query. It returns a SearchTopResultResponseType object and an error.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)

**Note:** This package uses Spotify's private API and is not officially supported by Spotify. Use at your own risk.
