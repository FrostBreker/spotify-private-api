package models

import "time"

// AlbumResponse represents the response from fetching an album.
type AlbumResponse struct {
	Data struct {
		AlbumUnion AlbumUnion `json:"albumUnion"`
	} `json:"data"`
	Extensions struct{} `json:"extensions"`
}

// AlbumUnion contains the main album data.
type AlbumUnion struct {
	Typename string `json:"__typename"`
	URI      string `json:"uri"`
	Name     string `json:"name"`
	Artists  struct {
		TotalCount int           `json:"totalCount"`
		Items      []AlbumArtist `json:"items"`
	} `json:"artists"`
	CoverArt CoverArt `json:"coverArt"`
	Discs    struct {
		TotalCount int `json:"totalCount"`
		Items      []struct {
			Number int `json:"number"`
			Tracks struct {
				TotalCount int `json:"totalCount"`
			} `json:"tracks"`
		} `json:"items"`
	} `json:"discs"`
	Releases struct {
		TotalCount int           `json:"totalCount"`
		Items      []interface{} `json:"items"`
	} `json:"releases"`
	Type string `json:"type"`
	Date struct {
		IsoString time.Time `json:"isoString"`
		Precision string    `json:"precision"`
	} `json:"date"`
	Playability struct {
		Playable bool   `json:"playable"`
		Reason   string `json:"reason"`
	} `json:"playability"`
	Label     string `json:"label"`
	Copyright struct {
		TotalCount int `json:"totalCount"`
		Items      []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"items"`
	} `json:"copyright"`
	CourtesyLine string `json:"courtesyLine"`
	Saved        bool   `json:"saved"`
	SharingInfo  struct {
		ShareURL string `json:"shareUrl"`
		ShareID  string `json:"shareId"`
	} `json:"sharingInfo"`
	Tracks struct {
		TotalCount int          `json:"totalCount"`
		Items      []AlbumTrack `json:"items"`
	} `json:"tracks"`
	MoreAlbumsByArtist struct {
		Items []struct {
			Discography struct {
				PopularReleasesAlbums struct {
					Items []struct {
						ID   string `json:"id"`
						URI  string `json:"uri"`
						Name string `json:"name"`
						Date struct {
							Year int `json:"year"`
						} `json:"date"`
						CoverArt    CoverArt `json:"coverArt"`
						Playability struct {
							Playable bool   `json:"playable"`
							Reason   string `json:"reason"`
						} `json:"playability"`
						SharingInfo struct {
							ShareID  string `json:"shareId"`
							ShareURL string `json:"shareUrl"`
						} `json:"sharingInfo"`
						Type string `json:"type"`
					} `json:"items"`
				} `json:"popularReleasesAlbums"`
			} `json:"discography"`
		} `json:"items"`
	} `json:"moreAlbumsByArtist"`
}

// AlbumArtist represents an artist in an album response.
type AlbumArtist struct {
	ID      string `json:"id"`
	URI     string `json:"uri"`
	Profile struct {
		Name string `json:"name"`
	} `json:"profile"`
	Visuals struct {
		AvatarImage struct {
			Sources []ImageSource `json:"sources"`
		} `json:"avatarImage"`
	} `json:"visuals"`
	SharingInfo struct {
		ShareURL string `json:"shareUrl"`
	} `json:"sharingInfo"`
}

// AlbumTrack represents a track in an album response.
type AlbumTrack struct {
	UID   string `json:"uid"`
	Track struct {
		Saved         bool   `json:"saved"`
		URI           string `json:"uri"`
		Name          string `json:"name"`
		Playcount     string `json:"playcount"`
		DiscNumber    int    `json:"discNumber"`
		TrackNumber   int    `json:"trackNumber"`
		ContentRating struct {
			Label string `json:"label"`
		} `json:"contentRating"`
		RelinkingInformation interface{} `json:"relinkingInformation"`
		Duration             struct {
			TotalMilliseconds int `json:"totalMilliseconds"`
		} `json:"duration"`
		Playability struct {
			Playable bool `json:"playable"`
		} `json:"playability"`
		Artists struct {
			Items []struct {
				URI     string `json:"uri"`
				Profile struct {
					Name string `json:"name"`
				} `json:"profile"`
			} `json:"items"`
		} `json:"artists"`
	} `json:"track"`
}
