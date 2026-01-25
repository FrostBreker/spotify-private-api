package spotifyprivateapi

import (
	"context"
	"encoding/base64"
	"errors"
	"net/http"
	"testing"

	pkgErr "github.com/FrostBreker/spotify-private-api/errors"
)

// mockArtistHTMLResponse creates a mock HTML response with embedded artist data
func mockArtistHTMLResponse(artistID, name string) *http.Response {
	initialStateJSON := `{"entities":{"items":{"spotify:artist:` + artistID + `":{"__typename":"Artist","uri":"spotify:artist:` + artistID + `","profile":{"name":"` + name + `","biography":{"text":"Bio"}},"stats":{"followers":1000000,"monthlyListeners":500000}}}}}`
	base64Data := base64.StdEncoding.EncodeToString([]byte(initialStateJSON))
	htmlBody := `<html><script id="initialState" type="text/plain">` + base64Data + `</script></html>`
	return mockResponse(http.StatusOK, htmlBody)
}

func TestClient_FetchArtist(t *testing.T) {
	tests := []struct {
		name     string
		artistID string
		mockResp *http.Response
		wantErr  bool
		wantName string
	}{
		{
			name:     "empty artist ID",
			artistID: "",
			wantErr:  true,
		},
		{
			name:     "successful fetch",
			artistID: "artist123",
			mockResp: mockArtistHTMLResponse("artist123", "Test Artist"),
			wantErr:  false,
			wantName: "Test Artist",
		},
		{
			name:     "API error",
			artistID: "artist123",
			mockResp: mockResponse(http.StatusNotFound, `{"error": "not found"}`),
			wantErr:  true,
		},
		{
			name:     "invalid JSON response",
			artistID: "artist123",
			mockResp: mockResponse(http.StatusOK, `invalid json`),
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &mockHTTPClient{
				doFunc: func(req *http.Request) (*http.Response, error) {
					return tt.mockResp, nil
				},
			}

			client := newTestClient(mockClient)
			result, err := client.FetchArtist(context.Background(), tt.artistID)

			if (err != nil) != tt.wantErr {
				t.Errorf("FetchArtist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && tt.artistID == "" {
				if !errors.Is(err, pkgErr.ErrInvalidID) {
					t.Errorf("FetchArtist() expected ErrInvalidID, got %v", err)
				}
				return
			}

			if !tt.wantErr && result == nil {
				t.Error("FetchArtist() returned nil result")
			}

			if !tt.wantErr && result != nil {
				if result.Data.ArtistUnion.Profile.Name != tt.wantName {
					t.Errorf("FetchArtist() got Name = %v, want %v", result.Data.ArtistUnion.Profile.Name, tt.wantName)
				}
			}
		})
	}
}

func TestClient_FetchArtist_ContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	mockClient := &mockHTTPClient{
		doFunc: func(req *http.Request) (*http.Response, error) {
			return nil, context.Canceled
		},
	}

	client := newTestClient(mockClient)
	_, err := client.FetchArtist(ctx, "artist123")

	if err == nil {
		t.Error("FetchArtist() expected error for canceled context")
	}
}
