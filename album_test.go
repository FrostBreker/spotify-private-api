package spotifyprivateapi

import (
	"context"
	"encoding/base64"
	"net/http"
	"testing"

	"github.com/FrostBreker/spotify-private-api/errors"
)

// mockAlbumHTMLResponse creates a mock HTML response with embedded album data
func mockAlbumHTMLResponse(albumID, name string) *http.Response {
	initialStateJSON := `{"entities":{"items":{"spotify:album:` + albumID + `":{"__typename":"Album","uri":"spotify:album:` + albumID + `","name":"` + name + `","tracks":{"totalCount":12,"items":[]}}}}}`
	base64Data := base64.StdEncoding.EncodeToString([]byte(initialStateJSON))
	htmlBody := `<html><script id="initialState" type="text/plain">` + base64Data + `</script></html>`
	return mockResponse(http.StatusOK, htmlBody)
}

func TestClient_FetchAlbum(t *testing.T) {
	tests := []struct {
		name     string
		albumID  string
		mockResp *http.Response
		wantErr  bool
		wantName string
	}{
		{
			name:    "empty album ID",
			albumID: "",
			wantErr: true,
		},
		{
			name:     "successful fetch",
			albumID:  "album123",
			mockResp: mockAlbumHTMLResponse("album123", "Test Album"),
			wantErr:  false,
			wantName: "Test Album",
		},
		{
			name:     "API error",
			albumID:  "album123",
			mockResp: mockResponse(http.StatusNotFound, `{"error": "not found"}`),
			wantErr:  true,
		},
		{
			name:     "invalid JSON response",
			albumID:  "album123",
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
			result, err := client.FetchAlbum(context.Background(), tt.albumID)

			if (err != nil) != tt.wantErr {
				t.Errorf("FetchAlbum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && tt.albumID == "" {
				if err != errors.ErrInvalidID {
					t.Errorf("FetchAlbum() expected ErrInvalidID, got %v", err)
				}
				return
			}

			if !tt.wantErr && result == nil {
				t.Error("FetchAlbum() returned nil result")
			}

			if !tt.wantErr && result != nil {
				if result.Data.AlbumUnion.Name != tt.wantName {
					t.Errorf("FetchAlbum() got Name = %v, want %v", result.Data.AlbumUnion.Name, tt.wantName)
				}
			}
		})
	}
}

func TestClient_FetchAlbum_ContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	mockClient := &mockHTTPClient{
		doFunc: func(req *http.Request) (*http.Response, error) {
			return nil, context.Canceled
		},
	}

	client := newTestClient(mockClient)
	_, err := client.FetchAlbum(ctx, "album123")

	if err == nil {
		t.Error("FetchAlbum() expected error for canceled context")
	}
}
