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

const spotifyArtistURL = "https://open.spotify.com/artist/"

// FetchArtist fetches information about an artist from Spotify.
func (c *Client) FetchArtist(ctx context.Context, artistID string) (*models.ArtistResponse, error) {
	if artistID == "" {
		return nil, errors.ErrInvalidID
	}

	c.logInfo("Fetching artist", "artistID", artistID)

	artistURL := spotifyArtistURL + artistID
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, artistURL, nil)
	if err != nil {
		return nil, errors.NewRequestError("create request", artistURL, err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.8")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.NewRequestError("fetch page", artistURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.NewAPIError(resp.StatusCode, "failed to fetch artist page")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.NewRequestError("read response", artistURL, err)
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

	artistKey := fmt.Sprintf("spotify:artist:%s", artistID)
	artistRaw, ok := initialState.Entities.Items[artistKey]
	if !ok {
		return nil, errors.NewParseError("artist not found in initialState", nil)
	}

	var htmlArtist htmlArtistData
	if err := json.Unmarshal(artistRaw, &htmlArtist); err != nil {
		return nil, errors.NewParseError("parse artist data", err)
	}

	response := htmlArtist.toArtistResponse(artistID)
	c.logInfo("Successfully fetched artist", "artistID", artistID, "name", response.Data.ArtistUnion.Profile.Name)
	return response, nil
}

type htmlArtistData struct {
	Typename string `json:"__typename"`
	URI      string `json:"uri"`
	Profile  struct {
		Name      string `json:"name"`
		Biography struct {
			Text string `json:"text"`
		} `json:"biography"`
		ExternalLinks struct {
			Items []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"items"`
		} `json:"externalLinks"`
	} `json:"profile"`
	Stats struct {
		Followers        int `json:"followers"`
		MonthlyListeners int `json:"monthlyListeners"`
	} `json:"stats"`
	Visuals struct {
		AvatarImage struct {
			Sources []struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"sources"`
			ExtractedColors struct {
				ColorRaw struct {
					Hex string `json:"hex"`
				} `json:"colorRaw"`
			} `json:"extractedColors"`
		} `json:"avatarImage"`
	} `json:"visuals"`
}

func (h *htmlArtistData) toArtistResponse(artistID string) *models.ArtistResponse {
	response := &models.ArtistResponse{}
	response.Data.ArtistUnion.Typename = h.Typename
	response.Data.ArtistUnion.URI = h.URI
	response.Data.ArtistUnion.ID = artistID
	response.Data.ArtistUnion.Profile.Name = h.Profile.Name
	response.Data.ArtistUnion.Profile.Biography.Text = h.Profile.Biography.Text
	response.Data.ArtistUnion.Stats.Followers = h.Stats.Followers
	response.Data.ArtistUnion.Stats.MonthlyListeners = h.Stats.MonthlyListeners

	for _, link := range h.Profile.ExternalLinks.Items {
		response.Data.ArtistUnion.Profile.ExternalLinks.Items = append(
			response.Data.ArtistUnion.Profile.ExternalLinks.Items,
			struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			}{Name: link.Name, URL: link.URL},
		)
	}

	for _, src := range h.Visuals.AvatarImage.Sources {
		response.Data.ArtistUnion.Visuals.AvatarImage.Sources = append(
			response.Data.ArtistUnion.Visuals.AvatarImage.Sources,
			models.ImageSource{URL: src.URL, Width: src.Width, Height: src.Height},
		)
	}
	response.Data.ArtistUnion.Visuals.AvatarImage.ExtractedColors.ColorRaw.Hex = h.Visuals.AvatarImage.ExtractedColors.ColorRaw.Hex

	return response
}
