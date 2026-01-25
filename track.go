package spotifyprivateapi

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/FrostBreker/spotify-private-api/errors"
	"github.com/FrostBreker/spotify-private-api/models"
)

const spotifyTrackURL = "https://open.spotify.com/track/"

// FetchTrack fetches information about a track from Spotify.
func (c *Client) FetchTrack(ctx context.Context, trackID string) (*models.TrackResponse, error) {
	if trackID == "" {
		return nil, errors.ErrInvalidID
	}

	c.logInfo("Fetching track", "trackID", trackID)

	trackURL := spotifyTrackURL + trackID
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, trackURL, nil)
	if err != nil {
		return nil, errors.NewRequestError("create request", trackURL, err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.8")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.NewRequestError("fetch page", trackURL, err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			c.logError("Failed to close response body", "error", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, errors.NewAPIError(resp.StatusCode, "failed to fetch track page")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.NewRequestError("read response", trackURL, err)
	}

	base64Data := extractInitialState(string(body))
	if base64Data == "" {
		return nil, errors.NewParseError("initialState not found in HTML", nil)
	}

	decoded, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return nil, errors.NewParseError("decode base64", err)
	}

	var initialState struct {
		Entities struct {
			Items map[string]json.RawMessage `json:"items"`
		} `json:"entities"`
	}

	if err := json.Unmarshal(decoded, &initialState); err != nil {
		return nil, errors.NewParseError("parse initialState", err)
	}

	trackKey := fmt.Sprintf("spotify:track:%s", trackID)
	trackRaw, ok := initialState.Entities.Items[trackKey]
	if !ok {
		return nil, errors.NewParseError("track not found in initialState", nil)
	}

	var htmlTrack htmlTrackData
	if err := json.Unmarshal(trackRaw, &htmlTrack); err != nil {
		return nil, errors.NewParseError("parse track data", err)
	}

	response := htmlTrack.toTrackResponse()
	c.logInfo("Successfully fetched track", "trackID", trackID, "name", response.Data.TrackUnion.Name)
	return response, nil
}

// FetchTrackPlayCount retrieves the play count for a track.
func (c *Client) FetchTrackPlayCount(ctx context.Context, trackID string) (*models.TrackPlayCount, error) {
	if trackID == "" {
		return nil, errors.ErrInvalidID
	}

	c.logInfo("Fetching track play count", "trackID", trackID)

	trackResponse, err := c.FetchTrack(ctx, trackID)
	if err != nil {
		return nil, err
	}

	var playCount int
	if trackResponse.Data.TrackUnion.Playcount != "" {
		playCount, err = strconv.Atoi(trackResponse.Data.TrackUnion.Playcount)
		if err != nil {
			return nil, errors.NewParseError("play count", err)
		}
	}

	c.logInfo("Successfully fetched track play count", "trackID", trackID, "playCount", playCount)
	return &models.TrackPlayCount{TrackID: trackID, PlayCount: playCount}, nil
}

type htmlTrackData struct {
	Typename  string `json:"__typename"`
	URI       string `json:"uri"`
	Name      string `json:"name"`
	Playcount string `json:"playcount"`
	Duration  struct {
		TotalMilliseconds int `json:"totalMilliseconds"`
	} `json:"duration"`
	AlbumOfTrack struct {
		URI      string `json:"uri"`
		Name     string `json:"name"`
		CoverArt struct {
			Sources []struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"sources"`
		} `json:"coverArt"`
		Date struct {
			Year int `json:"year"`
		} `json:"date"`
	} `json:"albumOfTrack"`
	FirstArtist struct {
		Items []struct {
			URI     string `json:"uri"`
			Profile struct {
				Name string `json:"name"`
			} `json:"profile"`
		} `json:"items"`
	} `json:"firstArtist"`
	ContentRating struct {
		Label string `json:"label"`
	} `json:"contentRating"`
}

func (h *htmlTrackData) toTrackResponse() *models.TrackResponse {
	response := &models.TrackResponse{}
	response.Data.TrackUnion.Typename = h.Typename
	response.Data.TrackUnion.URI = h.URI
	response.Data.TrackUnion.Name = h.Name
	response.Data.TrackUnion.Playcount = h.Playcount
	response.Data.TrackUnion.Duration.TotalMilliseconds = h.Duration.TotalMilliseconds
	response.Data.TrackUnion.ContentRating.Label = h.ContentRating.Label
	response.Data.TrackUnion.AlbumOfTrack.URI = h.AlbumOfTrack.URI
	response.Data.TrackUnion.AlbumOfTrack.Name = h.AlbumOfTrack.Name
	response.Data.TrackUnion.AlbumOfTrack.Date.Year = h.AlbumOfTrack.Date.Year

	for _, src := range h.AlbumOfTrack.CoverArt.Sources {
		response.Data.TrackUnion.AlbumOfTrack.CoverArt.Sources = append(
			response.Data.TrackUnion.AlbumOfTrack.CoverArt.Sources,
			models.ImageSource{URL: src.URL, Width: src.Width, Height: src.Height},
		)
	}

	for _, artist := range h.FirstArtist.Items {
		response.Data.TrackUnion.FirstArtist.Items = append(
			response.Data.TrackUnion.FirstArtist.Items,
			models.TrackArtist{
				URI: artist.URI,
				Profile: struct {
					Name string `json:"name"`
				}{Name: artist.Profile.Name},
			},
		)
	}

	return response
}

func extractInitialState(body string) string {
	patterns := []string{
		`<script id="initialState" type="text/plain">`,
		`<script id='initialState' type='text/plain'>`,
	}

	for _, pattern := range patterns {
		if idx := strings.Index(body, pattern); idx != -1 {
			start := idx + len(pattern)
			if end := strings.Index(body[start:], `</script>`); end != -1 {
				return strings.TrimSpace(body[start : start+end])
			}
		}
	}

	if idx := strings.Index(body, `id="initialState"`); idx != -1 {
		if tagEnd := strings.Index(body[idx:], ">"); tagEnd != -1 {
			start := idx + tagEnd + 1
			if end := strings.Index(body[start:], `</script>`); end != -1 {
				return strings.TrimSpace(body[start : start+end])
			}
		}
	}

	return ""
}
