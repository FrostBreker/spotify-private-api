// Package models contains all response types for the Spotify API.
package models

import "time"

// TrackResponse represents the response from fetching a track.
type TrackResponse struct {
	Data struct {
		TrackUnion TrackUnion `json:"trackUnion"`
	} `json:"data"`
	Extensions struct{} `json:"extensions"`
}

// TrackUnion contains the main track data.
type TrackUnion struct {
	Typename      string `json:"__typename"`
	ID            string `json:"id"`
	URI           string `json:"uri"`
	Name          string `json:"name"`
	ContentRating struct {
		Label string `json:"label"`
	} `json:"contentRating"`
	Duration struct {
		TotalMilliseconds int `json:"totalMilliseconds"`
	} `json:"duration"`
	Playability struct {
		Playable bool   `json:"playable"`
		Reason   string `json:"reason"`
	} `json:"playability"`
	TrackNumber int    `json:"trackNumber"`
	Playcount   string `json:"playcount"`
	Saved       bool   `json:"saved"`
	SharingInfo struct {
		ShareURL string `json:"shareUrl"`
		ShareID  string `json:"shareId"`
	} `json:"sharingInfo"`
	FirstArtist struct {
		TotalCount int           `json:"totalCount"`
		Items      []TrackArtist `json:"items"`
	} `json:"firstArtist"`
	OtherArtists struct {
		Items []TrackArtist `json:"items"`
	} `json:"otherArtists"`
	AlbumOfTrack TrackAlbum `json:"albumOfTrack"`
}

// TrackArtist represents an artist in a track response.
type TrackArtist struct {
	ID      string `json:"id"`
	URI     string `json:"uri"`
	Visuals struct {
		AvatarImage struct {
			Sources []ImageSource `json:"sources"`
		} `json:"avatarImage"`
	} `json:"visuals"`
	Profile struct {
		Name string `json:"name"`
	} `json:"profile"`
	Discography    ArtistDiscography `json:"discography"`
	RelatedContent struct {
		RelatedArtists struct {
			TotalCount int `json:"totalCount"`
			Items      []struct {
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
			} `json:"items"`
		} `json:"relatedArtists"`
	} `json:"relatedContent"`
}

// ArtistDiscography represents an artist's discography.
type ArtistDiscography struct {
	Singles               ReleaseCollection `json:"singles"`
	Albums                ReleaseCollection `json:"albums"`
	PopularReleasesAlbums struct {
		Items []ReleaseItem `json:"items"`
	} `json:"popularReleasesAlbums"`
	TopTracks struct {
		Items []struct {
			Track TopTrack `json:"track"`
		} `json:"items"`
	} `json:"topTracks"`
}

// ReleaseCollection represents a collection of releases.
type ReleaseCollection struct {
	TotalCount int `json:"totalCount"`
	Items      []struct {
		Releases struct {
			Items []ReleaseItem `json:"items"`
		} `json:"releases"`
	} `json:"items"`
}

// ReleaseItem represents an album or single release.
type ReleaseItem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	URI         string `json:"uri"`
	Playability struct {
		Playable bool `json:"playable"`
	} `json:"playability"`
	Date struct {
		IsoString time.Time `json:"isoString"`
		Precision string    `json:"precision"`
		Year      int       `json:"year"`
	} `json:"date"`
	SharingInfo struct {
		ShareID  string `json:"shareId"`
		ShareURL string `json:"shareUrl"`
	} `json:"sharingInfo"`
	Tracks struct {
		TotalCount int `json:"totalCount"`
		Items      []struct {
			Track struct {
				URI         string `json:"uri"`
				TrackNumber int    `json:"trackNumber"`
			} `json:"track"`
		} `json:"items"`
	} `json:"tracks"`
	CoverArt CoverArt `json:"coverArt"`
}

// TopTrack represents a top track.
type TopTrack struct {
	Artists struct {
		Items []struct {
			URI     string `json:"uri"`
			Profile struct {
				Name string `json:"name"`
			} `json:"profile"`
		} `json:"items"`
	} `json:"artists"`
	AlbumOfTrack struct {
		Name     string   `json:"name"`
		URI      string   `json:"uri"`
		CoverArt CoverArt `json:"coverArt"`
	} `json:"albumOfTrack"`
	Playability struct {
		Playable bool `json:"playable"`
	} `json:"playability"`
	Playcount string `json:"playcount"`
	Previews  struct {
		AudioPreviews struct {
			Items []struct {
				URL string `json:"url"`
			} `json:"items"`
		} `json:"audioPreviews"`
	} `json:"previews"`
	Duration struct {
		TotalMilliseconds int `json:"totalMilliseconds"`
	} `json:"duration"`
	ContentRating struct {
		Label string `json:"label"`
	} `json:"contentRating"`
	Name string `json:"name"`
	URI  string `json:"uri"`
	ID   string `json:"id"`
}

// TrackAlbum represents album information in a track response.
type TrackAlbum struct {
	ID        string `json:"id"`
	Copyright struct {
		TotalCount int `json:"totalCount"`
		Items      []struct {
			Text string `json:"text"`
			Type string `json:"type"`
		} `json:"items"`
	} `json:"copyright"`
	CourtesyLine string `json:"courtesyLine"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	URI          string `json:"uri"`
	Playability  struct {
		Playable bool `json:"playable"`
	} `json:"playability"`
	Date struct {
		IsoString time.Time `json:"isoString"`
		Precision string    `json:"precision"`
		Year      int       `json:"year"`
	} `json:"date"`
	SharingInfo struct {
		ShareID  string `json:"shareId"`
		ShareURL string `json:"shareUrl"`
	} `json:"sharingInfo"`
	Tracks struct {
		TotalCount int `json:"totalCount"`
		Items      []struct {
			Track struct {
				URI         string `json:"uri"`
				TrackNumber int    `json:"trackNumber"`
			} `json:"track"`
		} `json:"items"`
	} `json:"tracks"`
	CoverArt CoverArt `json:"coverArt"`
}

// TrackPlayCount represents a track's play count.
type TrackPlayCount struct {
	TrackID   string `json:"trackId"`
	PlayCount int    `json:"playCount"`
}

// ImageSource represents an image source with dimensions.
type ImageSource struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// CoverArt represents album/track cover art.
type CoverArt struct {
	ExtractedColors struct {
		ColorRaw struct {
			Hex string `json:"hex"`
		} `json:"colorRaw"`
		ColorLight struct {
			Hex string `json:"hex"`
		} `json:"colorLight"`
		ColorDark struct {
			Hex string `json:"hex"`
		} `json:"colorDark"`
	} `json:"extractedColors"`
	Sources []ImageSource `json:"sources"`
}
