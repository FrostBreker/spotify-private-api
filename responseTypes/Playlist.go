package responseTypes

import "time"

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
			Sources []struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"sources"`
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

type PlaylistResponse struct {
	Data struct {
		PlaylistV2 struct {
			Typename string `json:"__typename"`
			Content  struct {
				Typename   string `json:"__typename"`
				TotalCount int    `json:"totalCount"`
				PagingInfo struct {
					Offset int `json:"offset"`
					Limit  int `json:"limit"`
				} `json:"pagingInfo"`
				Items []struct {
					UID     string `json:"uid"`
					AddedAt struct {
						IsoString time.Time `json:"isoString"`
					} `json:"addedAt"`
					AddedBy    any `json:"addedBy"`
					Attributes []struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"attributes"`
					ItemV2 struct {
						Typename string        `json:"__typename"`
						Data     PlaylistTrack `json:"data"`
					} `json:"itemV2"`
				} `json:"items"`
			} `json:"content"`
		} `json:"playlistV2"`
	} `json:"data"`
	Extensions struct {
	} `json:"extensions"`
}
