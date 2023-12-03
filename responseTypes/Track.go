package responseTypes

import "time"

type Artist struct {
	ID      string `json:"id"`
	URI     string `json:"uri"`
	Visuals struct {
		AvatarImage struct {
			Sources []struct {
				Width  int    `json:"width"`
				Height int    `json:"height"`
				URL    string `json:"url"`
			} `json:"sources"`
		} `json:"avatarImage"`
	} `json:"visuals"`
	Profile struct {
		Name string `json:"name"`
	} `json:"profile"`
	Discography struct {
		Singles struct {
			TotalCount int `json:"totalCount"`
			Items      []struct {
				Releases struct {
					Items []struct {
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
						CoverArt struct {
							ExtractedColors struct {
								ColorRaw struct {
									Hex string `json:"hex"`
								} `json:"colorRaw"`
							} `json:"extractedColors"`
							Sources []struct {
								URL    string `json:"url"`
								Width  int    `json:"width"`
								Height int    `json:"height"`
							} `json:"sources"`
						} `json:"coverArt"`
					} `json:"items"`
				} `json:"releases"`
			} `json:"items"`
		} `json:"singles"`
		Albums struct {
			TotalCount int `json:"totalCount"`
			Items      []struct {
				Releases struct {
					Items []struct {
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
						CoverArt struct {
							ExtractedColors struct {
								ColorRaw struct {
									Hex string `json:"hex"`
								} `json:"colorRaw"`
							} `json:"extractedColors"`
							Sources []struct {
								URL    string `json:"url"`
								Width  int    `json:"width"`
								Height int    `json:"height"`
							} `json:"sources"`
						} `json:"coverArt"`
					} `json:"items"`
				} `json:"releases"`
			} `json:"items"`
		} `json:"albums"`
		PopularReleasesAlbums struct {
			Items []struct {
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
				CoverArt struct {
					ExtractedColors struct {
						ColorRaw struct {
							Hex string `json:"hex"`
						} `json:"colorRaw"`
					} `json:"extractedColors"`
					Sources []struct {
						URL    string `json:"url"`
						Width  int    `json:"width"`
						Height int    `json:"height"`
					} `json:"sources"`
				} `json:"coverArt"`
			} `json:"items"`
		} `json:"popularReleasesAlbums"`
		TopTracks struct {
			Items []struct {
				Track struct {
					Artists struct {
						Items []struct {
							URI     string `json:"uri"`
							Profile struct {
								Name string `json:"name"`
							} `json:"profile"`
						} `json:"items"`
					} `json:"artists"`
					AlbumOfTrack struct {
						Name     string `json:"name"`
						URI      string `json:"uri"`
						CoverArt struct {
							ExtractedColors struct {
								ColorRaw struct {
									Hex string `json:"hex"`
								} `json:"colorRaw"`
							} `json:"extractedColors"`
							Sources []struct {
								URL    string `json:"url"`
								Width  int    `json:"width"`
								Height int    `json:"height"`
							} `json:"sources"`
						} `json:"coverArt"`
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
				} `json:"track"`
			} `json:"items"`
		} `json:"topTracks"`
	} `json:"discography"`
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
						Sources []struct {
							Width  int    `json:"width"`
							Height int    `json:"height"`
							URL    string `json:"url"`
						} `json:"sources"`
					} `json:"avatarImage"`
				} `json:"visuals"`
			} `json:"items"`
		} `json:"relatedArtists"`
	} `json:"relatedContent"`
}

type TrackResponse struct {
	Data struct {
		TrackUnion struct {
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
				TotalCount int      `json:"totalCount"`
				Items      []Artist `json:"items"`
			} `json:"firstArtist"`
			OtherArtists struct {
				Items []Artist `json:"items"`
			} `json:"otherArtists"`
			AlbumOfTrack struct {
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
				CoverArt struct {
					ExtractedColors struct {
						ColorRaw struct {
							Hex string `json:"hex"`
						} `json:"colorRaw"`
					} `json:"extractedColors"`
					Sources []struct {
						URL    string `json:"url"`
						Width  int    `json:"width"`
						Height int    `json:"height"`
					} `json:"sources"`
				} `json:"coverArt"`
			} `json:"albumOfTrack"`
		} `json:"trackUnion"`
	} `json:"data"`
	Extensions struct {
	} `json:"extensions"`
}

type TrackPlayCount struct {
	TrackId   string
	PlayCount int
}
