package spotifyprivateapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/FrostBreker/spotify-private-api/responseTypes"
)

// FetchTrack fetches information about a track from Spotify's private API.
// It takes a trackId parameter, which is the Spotify ID of the track.
// The function returns a TrackResponse object and an error.
func (c *Client) FetchTrack(trackId string) (responseTypes.TrackResponse, error) {
	LogInfo("Fetching track: " + trackId)

	// Construct the request URL
	requestURL := "https://api-partner.spotify.com/pathfinder/v1/query?operationName=getTrack&variables=%7B%22uri%22%3A%22spotify%3Atrack%3A" + trackId + "%22%7D&extensions=%7B%22persistedQuery%22%3A%7B%22version%22%3A1%2C%22sha256Hash%22%3A%22e101aead6d78faa11d75bec5e36385a07b2f1c4a0420932d374d89ee17c70dd6%22%7D%7D"

	// Create a new GET request
	req, _ := http.NewRequest("GET", requestURL, nil)

	// Set headers for the Spotify request
	req, err := SetHeadersForSpotifyRequest(req)
	if err != nil {
		fmt.Println(err)
		LogError("Errored when setting headers for the spotifyAPI")
		return responseTypes.TrackResponse{}, err
	}

	// Send the request
	resp, err := HttpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		LogError("Errored when sending request to the spotifyAPI")
		return responseTypes.TrackResponse{}, err
	}

	defer resp.Body.Close()

	// Read the response body
	jsonStr, _ := io.ReadAll(resp.Body)

	// Unmarshal the JSON response into a TrackResponse object
	var trackResponse responseTypes.TrackResponse
	json.Unmarshal(jsonStr, &trackResponse)

	LogInfo("Successfully fetched track: " + trackId)
	return trackResponse, nil
}

// FetchTrackPlayCount retrieves the play count for a track.
// It takes a trackId parameter, which is the Spotify ID of the track.
// The function internally calls the FetchTrack function to fetch the track information
// and then extracts the play count from the response.
// It returns a TrackPlayCount object and an error.
func (c *Client) FetchTrackPlayCount(trackId string) (responseTypes.TrackPlayCount, error) {
	LogInfo("Retrieving track playcount for track: " + trackId)

	// Fetch the track information
	trackResponse, err := c.FetchTrack(trackId)
	if err != nil {
		LogError("Error fetching track (" + trackId + "): " + err.Error())
		return responseTypes.TrackPlayCount{}, err
	}

	// Convert the play count to an integer
	playCount, err := strconv.Atoi(trackResponse.Data.TrackUnion.Playcount)
	if err != nil {
		LogError("Error during conversion")
		return responseTypes.TrackPlayCount{}, err
	}

	// Create a TrackPlayCount object
	playCountObject := responseTypes.TrackPlayCount{
		TrackId:   trackId,
		PlayCount: playCount,
	}
	LogInfo("Successfully retrieved track playcount for track: " + trackId)
	return playCountObject, nil
}

// The `FetchArtist` function is a method of the `Client` struct in the `spotifyprivateapi` package.
// It is used to fetch information about an artist from Spotify's private API.
// It takes an artistId parameter, which is the Spotify ID of the artist.
// The function returns an Artist object and an error.
func (c *Client) FetchArtist(artistId string) (responseTypes.ArtistResponseType, error) {
	LogInfo("Fetching artist: " + artistId)

	// Construct the request URL
	requestURL := "https://api-partner.spotify.com/pathfinder/v1/query?operationName=queryArtistOverview&variables=%7B%22uri%22%3A%22spotify%3Aartist%3A" + artistId + "%22%2C%22locale%22%3A%22%22%2C%22includePrerelease%22%3Afalse%7D&extensions=%7B%22persistedQuery%22%3A%7B%22version%22%3A1%2C%22sha256Hash%22%3A%2235648a112beb1794e39ab931365f6ae4a8d45e65396d641eeda94e4003d41497%22%7D%7D"

	// Create a new GET request
	req, _ := http.NewRequest("GET", requestURL, nil)

	// Set headers for the Spotify request
	req, err := SetHeadersForSpotifyRequest(req)
	if err != nil {
		fmt.Println(err)
		LogError("Errored when setting headers for the spotifyAPI")
		return responseTypes.ArtistResponseType{}, err
	}

	// Send the request
	resp, err := HttpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		LogError("Errored when sending request to the spotifyAPI")
		return responseTypes.ArtistResponseType{}, err
	}

	defer resp.Body.Close()

	// Read the response body
	jsonStr, _ := io.ReadAll(resp.Body)

	// Unmarshal the JSON response into a Artist object
	var artistResponse responseTypes.ArtistResponseType
	json.Unmarshal(jsonStr, &artistResponse)

	LogInfo("Successfully fetched artist: " + artistId)
	return artistResponse, nil
}

// The `FetchAlbum` function is a method of the `Client` struct in the `spotifyprivateapi` package.
// It is used to fetch information about an album from Spotify's private API.
// It takes an albumId parameter, which is the Spotify ID of the album.
// The function returns an AlbumResponseType object and an error.
func (c *Client) FetchAlbum(albumId string) (responseTypes.AlbumResponseType, error) {
	LogInfo("Fetching album: " + albumId)
	// The `requestURL` variable is a string that contains the URL for making a GET request to Spotify's
	// private API to fetch information about an album. It includes the `albumId` parameter in the URL to
	// specify which album to fetch.
	requestURL := "https://api-partner.spotify.com/pathfinder/v1/query?operationName=getAlbum&variables=%7B%22uri%22%3A%22spotify%3Aalbum%3A" + albumId + "%22%2C%22locale%22%3A%22%22%2C%22offset%22%3A0%2C%22limit%22%3A50%7D&extensions=%7B%22persistedQuery%22%3A%7B%22version%22%3A1%2C%22sha256Hash%22%3A%2246ae954ef2d2fe7732b4b2b4022157b2e18b7ea84f70591ceb164e4de1b5d5d3%22%7D%7D"
	client := &http.Client{}

	req, _ := http.NewRequest("GET", requestURL, nil)
	req, err := SetHeadersForSpotifyRequest(req)
	if err != nil {
		fmt.Println(err)
		LogError("Errored when setting headers for the spotifyAPI")
		return responseTypes.AlbumResponseType{}, err
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		LogError("Errored when sending request to the spotifyAPI")
		return responseTypes.AlbumResponseType{}, err
	}

	defer resp.Body.Close()
	// `jsonStr, _ := io.ReadAll(resp.Body)` is reading the response body from the HTTP response and
	// storing it in the `jsonStr` variable. The `io.ReadAll()` function reads all the data from the
	// `resp.Body` and returns it as a byte slice. The `_` is used to ignore the error returned by
	// `io.ReadAll()` since we are assuming that the read operation will be successful.
	jsonStr, _ := io.ReadAll(resp.Body)

	// The code snippet is unmarshaling the JSON response from the Spotify API into the `albumResponse`
	// variable, which is of type `responseTypes.AlbumResponseType`.
	var albumResponse responseTypes.AlbumResponseType
	json.Unmarshal(jsonStr, &albumResponse)
	LogInfo("Succesfully fetched album: " + albumId)
	return albumResponse, nil
}
