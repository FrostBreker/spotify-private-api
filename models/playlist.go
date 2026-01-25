package models

import "time"

// PlaylistResponse represents the response from fetching a playlist.
type PlaylistResponse struct {
	Data struct {
		PlaylistV2 PlaylistV2 `json:"playlistV2"`
	} `json:"data"`
	Extensions struct{} `json:"extensions"`
}

// PlaylistV2 contains the main playlist data.
type PlaylistV2 struct {
	Typename string          `json:"__typename"`
	Content  PlaylistContent `json:"content"`
}

// PlaylistContent contains the playlist content.
type PlaylistContent struct {
	Typename   string `json:"__typename"`
	TotalCount int    `json:"totalCount"`
	PagingInfo struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"pagingInfo"`
	Items []PlaylistItem `json:"items"`
}

// PlaylistItem represents an item in a playlist.
type PlaylistItem struct {
	UID     string `json:"uid"`
	AddedAt struct {
		IsoString time.Time `json:"isoString"`
	} `json:"addedAt"`
	AddedBy    interface{} `json:"addedBy"`
	Attributes []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"attributes"`
	ItemV2 struct {
		Typename string        `json:"__typename"`
		Data     PlaylistTrack `json:"data"`
	} `json:"itemV2"`
}

// PlaylistTrack represents a track in a playlist.
type PlaylistTrack struct {
	Typename      string `json:"__typename"`
	URI           string `json:"uri"`
	Name          string `json:"name"`
	TrackDuration struct {
		TotalMilliseconds int `json:"totalMilliseconds"`
	} `json:"trackDuration"`
	Playcount    string `json:"playcount"`
	AlbumOfTrack struct {
		URI     string `json:"uri"`
		Name    string `json:"name"`
		Artists struct {
			Items []struct {
				URI     string `json:"uri"`
				Profile struct {
					Name string `json:"name"`
				} `json:"profile"`
			} `json:"items"`
		} `json:"artists"`
		CoverArt struct {
			Sources []ImageSource `json:"sources"`
		} `json:"coverArt"`
	} `json:"albumOfTrack"`
	Artists struct {
		Items []struct {
			URI     string `json:"uri"`
			Profile struct {
				Name string `json:"name"`
			} `json:"profile"`
		} `json:"items"`
	} `json:"artists"`
	DiscNumber  int `json:"discNumber"`
	TrackNumber int `json:"trackNumber"`
	Playability struct {
		Playable bool   `json:"playable"`
		Reason   string `json:"reason"`
	} `json:"playability"`
	ContentRating struct {
		Label string `json:"label"`
	} `json:"contentRating"`
}
