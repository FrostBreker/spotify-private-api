package responseTypes

import "time"

type AlbumTrackResponseType struct {
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
		RelinkingInformation any `json:"relinkingInformation"`
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

type AlbumResponseType struct {
	Data struct {
		AlbumUnion struct {
			Typename string `json:"__typename"`
			URI      string `json:"uri"`
			Name     string `json:"name"`
			Artists  struct {
				TotalCount int `json:"totalCount"`
				Items      []struct {
					ID      string `json:"id"`
					URI     string `json:"uri"`
					Profile struct {
						Name string `json:"name"`
					} `json:"profile"`
					Visuals struct {
						AvatarImage struct {
							Sources []struct {
								URL    string `json:"url"`
								Width  int    `json:"width"`
								Height int    `json:"height"`
							} `json:"sources"`
						} `json:"avatarImage"`
					} `json:"visuals"`
					SharingInfo struct {
						ShareURL string `json:"shareUrl"`
					} `json:"sharingInfo"`
				} `json:"items"`
			} `json:"artists"`
			CoverArt struct {
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
				Sources []struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"sources"`
			} `json:"coverArt"`
			Discs struct {
				TotalCount int `json:"totalCount"`
				Items      []struct {
					Number int `json:"number"`
					Tracks struct {
						TotalCount int `json:"totalCount"`
					} `json:"tracks"`
				} `json:"items"`
			} `json:"discs"`
			Releases struct {
				TotalCount int   `json:"totalCount"`
				Items      []any `json:"items"`
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
				TotalCount int                      `json:"totalCount"`
				Items      []AlbumTrackResponseType `json:"items"`
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
								CoverArt struct {
									Sources []struct {
										URL    string `json:"url"`
										Width  int    `json:"width"`
										Height int    `json:"height"`
									} `json:"sources"`
								} `json:"coverArt"`
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
		} `json:"albumUnion"`
	} `json:"data"`
	Extensions struct {
	} `json:"extensions"`
}
