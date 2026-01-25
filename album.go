package spotifyprivateapi

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/FrostBreker/spotify-private-api/errors"
	"github.com/FrostBreker/spotify-private-api/models"
)

const spotifyAlbumURL = "https://open.spotify.com/album/"

// FetchAlbum fetches information about an album from Spotify.
func (c *Client) FetchAlbum(ctx context.Context, albumID string) (*models.AlbumResponse, error) {
	if albumID == "" {
		return nil, errors.ErrInvalidID
	}

	c.logInfo("Fetching album", "albumID", albumID)

	albumURL := spotifyAlbumURL + albumID
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, albumURL, nil)
	if err != nil {
		return nil, errors.NewRequestError("create request", albumURL, err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.8")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.NewRequestError("fetch page", albumURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.NewAPIError(resp.StatusCode, "failed to fetch album page")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.NewRequestError("read response", albumURL, err)
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

	albumKey := fmt.Sprintf("spotify:album:%s", albumID)
	albumRaw, ok := initialState.Entities.Items[albumKey]
	if !ok {
		return nil, errors.NewParseError("album not found in initialState", nil)
	}

	var htmlAlbum htmlAlbumData
	if err := json.Unmarshal(albumRaw, &htmlAlbum); err != nil {
		return nil, errors.NewParseError("parse album data", err)
	}

	response := htmlAlbum.toAlbumResponse()
	c.logInfo("Successfully fetched album", "albumID", albumID, "name", response.Data.AlbumUnion.Name)
	return response, nil
}

type htmlAlbumData struct {
	Typename string `json:"__typename"`
	URI      string `json:"uri"`
	Name     string `json:"name"`
	Label    string `json:"label"`
	Artists  struct {
		Items []struct {
			URI     string `json:"uri"`
			Profile struct {
				Name string `json:"name"`
			} `json:"profile"`
		} `json:"items"`
	} `json:"artists"`
	CoverArt struct {
		Sources []struct {
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"sources"`
	} `json:"coverArt"`
	Date struct {
		Precision string `json:"precision"`
	} `json:"date"`
	Tracks struct {
		TotalCount int `json:"totalCount"`
		Items      []struct {
			Track struct {
				URI       string `json:"uri"`
				Name      string `json:"name"`
				Playcount string `json:"playcount"`
				Duration  struct {
					TotalMilliseconds int `json:"totalMilliseconds"`
				} `json:"duration"`
			} `json:"track"`
		} `json:"items"`
	} `json:"tracks"`
	Copyright struct {
		Items []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"items"`
	} `json:"copyright"`
}

func (h *htmlAlbumData) toAlbumResponse() *models.AlbumResponse {
	response := &models.AlbumResponse{}
	response.Data.AlbumUnion.Typename = h.Typename
	response.Data.AlbumUnion.URI = h.URI
	response.Data.AlbumUnion.Name = h.Name
	response.Data.AlbumUnion.Label = h.Label
	response.Data.AlbumUnion.Date.Precision = h.Date.Precision

	for _, artist := range h.Artists.Items {
		response.Data.AlbumUnion.Artists.Items = append(
			response.Data.AlbumUnion.Artists.Items,
			models.AlbumArtist{
				URI: artist.URI,
				Profile: struct {
					Name string `json:"name"`
				}{Name: artist.Profile.Name},
			},
		)
	}

	for _, src := range h.CoverArt.Sources {
		response.Data.AlbumUnion.CoverArt.Sources = append(
			response.Data.AlbumUnion.CoverArt.Sources,
			models.ImageSource{URL: src.URL, Width: src.Width, Height: src.Height},
		)
	}

	response.Data.AlbumUnion.Tracks.TotalCount = h.Tracks.TotalCount
	for _, item := range h.Tracks.Items {
		albumTrack := models.AlbumTrack{}
		albumTrack.Track.URI = item.Track.URI
		albumTrack.Track.Name = item.Track.Name
		albumTrack.Track.Playcount = item.Track.Playcount
		albumTrack.Track.Duration.TotalMilliseconds = item.Track.Duration.TotalMilliseconds
		response.Data.AlbumUnion.Tracks.Items = append(response.Data.AlbumUnion.Tracks.Items, albumTrack)
	}

	response.Data.AlbumUnion.Copyright.TotalCount = len(h.Copyright.Items)
	return response
}
