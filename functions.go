package spotifyprivateapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/FrostBreker/spotify-private-api/responseTypes"
)

func (c *Client) FetchTrack(trackId string) (responseTypes.TrackResponse, error) {
	if c.Debug {
		LogInfo("Fetching track: " + trackId)
	}
	requestURL := "https://api-partner.spotify.com/pathfinder/v1/query?operationName=getTrack&variables=%7B%22uri%22%3A%22spotify%3Atrack%3A" + trackId + "%22%7D&extensions=%7B%22persistedQuery%22%3A%7B%22version%22%3A1%2C%22sha256Hash%22%3A%22e101aead6d78faa11d75bec5e36385a07b2f1c4a0420932d374d89ee17c70dd6%22%7D%7D"

	req, _ := http.NewRequest("GET", requestURL, nil)
	req, err := SetHeadersForSpotifyRequest(req)
	if err != nil {
		fmt.Println(err)
		fmt.Println(GetFormatedTime() + " --> Errored when setting headers for the spotifyAPI")
		return responseTypes.TrackResponse{}, err
	}

	resp, err := HttpClient.Do(req)

	if err != nil {
		fmt.Println(err)
		fmt.Println(GetFormatedTime() + " --> Errored when sending request to the spotifyAPI")
		return responseTypes.TrackResponse{}, err
	}

	defer resp.Body.Close()
	jsonStr, _ := io.ReadAll(resp.Body)

	var trackResponse responseTypes.TrackResponse
	json.Unmarshal(jsonStr, &trackResponse)
	if c.Debug {
		LogInfo("Successfully fetched track: " + trackId)
	}
	return trackResponse, nil
}

func (c *Client) FetchTrackPlayCount(trackId string) (responseTypes.TrackPlayCount, error) {
	if c.Debug {
		LogInfo("Retrieving track playcount for track: " + trackId)
	}
	trackResponse, err := c.FetchTrack(trackId)
	if err != nil {
		if c.Debug {
			LogError("Error fetching track (" + trackId + "): " + err.Error())
		}
		return responseTypes.TrackPlayCount{}, err
	}
	playCount, err := strconv.Atoi(trackResponse.Data.TrackUnion.Playcount)
	if err != nil {
		LogError("Error during conversion")
		return responseTypes.TrackPlayCount{}, err
	}

	playCountObject := responseTypes.TrackPlayCount{
		TrackId:   trackId,
		PlayCount: playCount,
	}
	return playCountObject, nil
}
