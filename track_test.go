package spotifyprivateapi

import (
	"context"
	"encoding/base64"
	"net/http"
	"testing"

	"github.com/FrostBreker/spotify-private-api/errors"
)

// mockHTMLResponse creates a mock HTML response with embedded initialState for testing
func mockHTMLResponse(entityType, entityID, name string) *http.Response {
	initialStateJSON := `{"entities":{"items":{"spotify:` + entityType + `:` + entityID + `":{"__typename":"` + entityType + `","uri":"spotify:` + entityType + `:` + entityID + `","name":"` + name + `","playcount":"1000000","duration":{"totalMilliseconds":180000}}}}}`
	base64Data := base64.StdEncoding.EncodeToString([]byte(initialStateJSON))
	htmlBody := `<html><script id="initialState" type="text/plain">` + base64Data + `</script></html>`
	return mockResponse(http.StatusOK, htmlBody)
}

func TestClient_FetchTrack(t *testing.T) {
	tests := []struct {
		name        string
		trackID     string
		mockResp    *http.Response
		mockErr     error
		wantErr     bool
		wantErrType error
	}{
		{
			name:        "empty track ID",
			trackID:     "",
			wantErr:     true,
			wantErrType: errors.ErrInvalidID,
		},
		{
			name:     "successful fetch",
			trackID:  "test123",
			mockResp: mockHTMLResponse("track", "test123", "Test Track"),
			wantErr:  false,
		},
		{
			name:     "API error",
			trackID:  "test123",
			mockResp: mockResponse(http.StatusNotFound, `{"error": "not found"}`),
			wantErr:  true,
		},
		{
			name:     "invalid JSON response",
			trackID:  "test123",
			mockResp: mockResponse(http.StatusOK, `invalid json`),
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &mockHTTPClient{
				doFunc: func(req *http.Request) (*http.Response, error) {
					if tt.mockErr != nil {
						return nil, tt.mockErr
					}
					return tt.mockResp, nil
				},
			}

			client := newTestClient(mockClient)
			result, err := client.FetchTrack(context.Background(), tt.trackID)

			if (err != nil) != tt.wantErr {
				t.Errorf("FetchTrack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErrType != nil && err != tt.wantErrType {
				t.Errorf("FetchTrack() error = %v, wantErrType %v", err, tt.wantErrType)
				return
			}

			if !tt.wantErr && result == nil {
				t.Error("FetchTrack() returned nil result")
			}

			if !tt.wantErr && result != nil {
				if result.Data.TrackUnion.Name != "Test Track" {
					t.Errorf("FetchTrack() got Name = %v, want Test Track", result.Data.TrackUnion.Name)
				}
			}
		})
	}
}

func TestClient_FetchTrackPlayCount(t *testing.T) {
	tests := []struct {
		name      string
		trackID   string
		mockResp  *http.Response
		wantErr   bool
		wantCount int
	}{
		{
			name:    "empty track ID",
			trackID: "",
			wantErr: true,
		},
		{
			name:      "successful fetch",
			trackID:   "test123",
			mockResp:  mockHTMLResponse("track", "test123", "Test Track"),
			wantErr:   false,
			wantCount: 1000000,
		},
		{
			name:     "invalid playcount",
			trackID:  "test123",
			mockResp: mockResponse(http.StatusOK, `<html><script id="initialState" type="text/plain">`+base64.StdEncoding.EncodeToString([]byte(`{"entities":{"items":{"spotify:track:test123":{"__typename":"track","uri":"spotify:track:test123","name":"Test","playcount":"invalid"}}}}`))+`</script></html>`),
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
			result, err := client.FetchTrackPlayCount(context.Background(), tt.trackID)

			if (err != nil) != tt.wantErr {
				t.Errorf("FetchTrackPlayCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && result == nil {
				t.Error("FetchTrackPlayCount() returned nil result")
			}

			if !tt.wantErr && result != nil && result.PlayCount != tt.wantCount {
				t.Errorf("FetchTrackPlayCount() got PlayCount = %v, want %v", result.PlayCount, tt.wantCount)
			}
		})
	}
}

func TestClient_FetchTrack_ContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	mockClient := &mockHTTPClient{
		doFunc: func(req *http.Request) (*http.Response, error) {
			return nil, context.Canceled
		},
	}

	client := newTestClient(mockClient)
	_, err := client.FetchTrack(ctx, "test123")

	if err == nil {
		t.Error("FetchTrack() expected error for canceled context")
	}
}
