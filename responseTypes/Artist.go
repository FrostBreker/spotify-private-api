package responseTypes

type ArtistResponseType struct {
	Data struct {
		ArtistUnion struct {
			Typename    string `json:"__typename"`
			ID          string `json:"id"`
			URI         string `json:"uri"`
			Saved       bool   `json:"saved"`
			SharingInfo struct {
				ShareURL string `json:"shareUrl"`
				ShareID  string `json:"shareId"`
			} `json:"sharingInfo"`
			Profile struct {
				Name       string `json:"name"`
				Verified   bool   `json:"verified"`
				PinnedItem struct {
					Comment         string `json:"comment"`
					Type            string `json:"type"`
					BackgroundImage struct {
						Sources []struct {
							URL string `json:"url"`
						} `json:"sources"`
					} `json:"backgroundImage"`
					ItemV2 struct {
					} `json:"itemV2"`
					Item struct {
						URI    string `json:"uri"`
						Name   string `json:"name"`
						Images struct {
							Items []struct {
								Sources []struct {
									URL    string `json:"url"`
									Width  any    `json:"width"`
									Height any    `json:"height"`
								} `json:"sources"`
							} `json:"items"`
						} `json:"images"`
					} `json:"item"`
				} `json:"pinnedItem"`
				Biography struct {
					Type string `json:"type"`
					Text string `json:"text"`
				} `json:"biography"`
				ExternalLinks struct {
					Items []struct {
						Name string `json:"name"`
						URL  string `json:"url"`
					} `json:"items"`
				} `json:"externalLinks"`
				PlaylistsV2 struct {
					TotalCount int `json:"totalCount"`
					Items      []struct {
						Data struct {
							Typename    string `json:"__typename"`
							URI         string `json:"uri"`
							Name        string `json:"name"`
							Description string `json:"description"`
							OwnerV2     struct {
								Data struct {
									Typename string `json:"__typename"`
									Name     string `json:"name"`
								} `json:"data"`
							} `json:"ownerV2"`
							Images struct {
								Items []struct {
									Sources []struct {
										URL    string `json:"url"`
										Width  any    `json:"width"`
										Height any    `json:"height"`
									} `json:"sources"`
								} `json:"items"`
							} `json:"images"`
						} `json:"data"`
					} `json:"items"`
				} `json:"playlistsV2"`
			} `json:"profile"`
			Visuals struct {
				Gallery struct {
					Items []any `json:"items"`
				} `json:"gallery"`
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
				HeaderImage struct {
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
				} `json:"headerImage"`
			} `json:"visuals"`
			Discography struct {
				Latest struct {
					ID        string `json:"id"`
					URI       string `json:"uri"`
					Name      string `json:"name"`
					Type      string `json:"type"`
					Copyright struct {
						Items []struct {
							Type string `json:"type"`
							Text string `json:"text"`
						} `json:"items"`
					} `json:"copyright"`
					Date struct {
						Year      int    `json:"year"`
						Month     int    `json:"month"`
						Day       int    `json:"day"`
						Precision string `json:"precision"`
					} `json:"date"`
					CoverArt struct {
						Sources []struct {
							URL    string `json:"url"`
							Width  int    `json:"width"`
							Height int    `json:"height"`
						} `json:"sources"`
					} `json:"coverArt"`
					Tracks struct {
						TotalCount int `json:"totalCount"`
					} `json:"tracks"`
					Label       string `json:"label"`
					Playability struct {
						Playable bool   `json:"playable"`
						Reason   string `json:"reason"`
					} `json:"playability"`
					SharingInfo struct {
						ShareID  string `json:"shareId"`
						ShareURL string `json:"shareUrl"`
					} `json:"sharingInfo"`
				} `json:"latest"`
				PopularReleasesAlbums struct {
					TotalCount int `json:"totalCount"`
					Items      []struct {
						ID        string `json:"id"`
						URI       string `json:"uri"`
						Name      string `json:"name"`
						Type      string `json:"type"`
						Copyright struct {
							Items []struct {
								Type string `json:"type"`
								Text string `json:"text"`
							} `json:"items"`
						} `json:"copyright"`
						Date struct {
							Year      int    `json:"year"`
							Month     int    `json:"month"`
							Day       int    `json:"day"`
							Precision string `json:"precision"`
						} `json:"date"`
						CoverArt struct {
							Sources []struct {
								URL    string `json:"url"`
								Width  int    `json:"width"`
								Height int    `json:"height"`
							} `json:"sources"`
						} `json:"coverArt"`
						Tracks struct {
							TotalCount int `json:"totalCount"`
						} `json:"tracks"`
						Label       string `json:"label"`
						Playability struct {
							Playable bool   `json:"playable"`
							Reason   string `json:"reason"`
						} `json:"playability"`
						SharingInfo struct {
							ShareID  string `json:"shareId"`
							ShareURL string `json:"shareUrl"`
						} `json:"sharingInfo"`
					} `json:"items"`
				} `json:"popularReleasesAlbums"`
				Singles struct {
					TotalCount int `json:"totalCount"`
					Items      []struct {
						Releases struct {
							Items []struct {
								ID        string `json:"id"`
								URI       string `json:"uri"`
								Name      string `json:"name"`
								Type      string `json:"type"`
								Copyright struct {
									Items []struct {
										Type string `json:"type"`
										Text string `json:"text"`
									} `json:"items"`
								} `json:"copyright"`
								Date struct {
									Year      int    `json:"year"`
									Month     int    `json:"month"`
									Day       int    `json:"day"`
									Precision string `json:"precision"`
								} `json:"date"`
								CoverArt struct {
									Sources []struct {
										URL    string `json:"url"`
										Width  int    `json:"width"`
										Height int    `json:"height"`
									} `json:"sources"`
								} `json:"coverArt"`
								Tracks struct {
									TotalCount int `json:"totalCount"`
								} `json:"tracks"`
								Label       string `json:"label"`
								Playability struct {
									Playable bool   `json:"playable"`
									Reason   string `json:"reason"`
								} `json:"playability"`
								SharingInfo struct {
									ShareID  string `json:"shareId"`
									ShareURL string `json:"shareUrl"`
								} `json:"sharingInfo"`
							} `json:"items"`
						} `json:"releases"`
					} `json:"items"`
				} `json:"singles"`
				Albums struct {
					TotalCount int `json:"totalCount"`
					Items      []struct {
						Releases struct {
							Items []struct {
								ID        string `json:"id"`
								URI       string `json:"uri"`
								Name      string `json:"name"`
								Type      string `json:"type"`
								Copyright struct {
									Items []struct {
										Type string `json:"type"`
										Text string `json:"text"`
									} `json:"items"`
								} `json:"copyright"`
								Date struct {
									Year      int    `json:"year"`
									Month     int    `json:"month"`
									Day       int    `json:"day"`
									Precision string `json:"precision"`
								} `json:"date"`
								CoverArt struct {
									Sources []struct {
										URL    string `json:"url"`
										Width  int    `json:"width"`
										Height int    `json:"height"`
									} `json:"sources"`
								} `json:"coverArt"`
								Tracks struct {
									TotalCount int `json:"totalCount"`
								} `json:"tracks"`
								Label       string `json:"label"`
								Playability struct {
									Playable bool   `json:"playable"`
									Reason   string `json:"reason"`
								} `json:"playability"`
								SharingInfo struct {
									ShareID  string `json:"shareId"`
									ShareURL string `json:"shareUrl"`
								} `json:"sharingInfo"`
							} `json:"items"`
						} `json:"releases"`
					} `json:"items"`
				} `json:"albums"`
				Compilations struct {
					TotalCount int   `json:"totalCount"`
					Items      []any `json:"items"`
				} `json:"compilations"`
				TopTracks struct {
					Items []struct {
						UID   string `json:"uid"`
						Track struct {
							ID         string `json:"id"`
							URI        string `json:"uri"`
							Name       string `json:"name"`
							Playcount  string `json:"playcount"`
							DiscNumber int    `json:"discNumber"`
							Duration   struct {
								TotalMilliseconds int `json:"totalMilliseconds"`
							} `json:"duration"`
							Playability struct {
								Playable bool   `json:"playable"`
								Reason   string `json:"reason"`
							} `json:"playability"`
							ContentRating struct {
								Label string `json:"label"`
							} `json:"contentRating"`
							Artists struct {
								Items []struct {
									URI     string `json:"uri"`
									Profile struct {
										Name string `json:"name"`
									} `json:"profile"`
								} `json:"items"`
							} `json:"artists"`
							AlbumOfTrack struct {
								URI      string `json:"uri"`
								CoverArt struct {
									Sources []struct {
										URL string `json:"url"`
									} `json:"sources"`
								} `json:"coverArt"`
							} `json:"albumOfTrack"`
						} `json:"track"`
					} `json:"items"`
				} `json:"topTracks"`
			} `json:"discography"`
			Stats struct {
				Followers        int `json:"followers"`
				MonthlyListeners int `json:"monthlyListeners"`
				WorldRank        int `json:"worldRank"`
				TopCities        struct {
					Items []struct {
						NumberOfListeners int    `json:"numberOfListeners"`
						City              string `json:"city"`
						Country           string `json:"country"`
						Region            string `json:"region"`
					} `json:"items"`
				} `json:"topCities"`
			} `json:"stats"`
			RelatedContent struct {
				AppearsOn struct {
					TotalCount int `json:"totalCount"`
					Items      []struct {
						Releases struct {
							TotalCount int `json:"totalCount"`
							Items      []struct {
								URI     string `json:"uri"`
								ID      string `json:"id"`
								Name    string `json:"name"`
								Type    string `json:"type"`
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
								Date struct {
									Year int `json:"year"`
								} `json:"date"`
								SharingInfo struct {
									ShareID  string `json:"shareId"`
									ShareURL string `json:"shareUrl"`
								} `json:"sharingInfo"`
							} `json:"items"`
						} `json:"releases"`
					} `json:"items"`
				} `json:"appearsOn"`
				FeaturingV2 struct {
					TotalCount int `json:"totalCount"`
					Items      []struct {
						Data struct {
							Typename string `json:"__typename"`
							URI      string `json:"uri"`
							ID       string `json:"id"`
							OwnerV2  struct {
								Data struct {
									Typename string `json:"__typename"`
									Name     string `json:"name"`
								} `json:"data"`
							} `json:"ownerV2"`
							Name        string `json:"name"`
							Description string `json:"description"`
							Images      struct {
								TotalCount int `json:"totalCount"`
								Items      []struct {
									Sources []struct {
										URL    string `json:"url"`
										Width  any    `json:"width"`
										Height any    `json:"height"`
									} `json:"sources"`
								} `json:"items"`
							} `json:"images"`
						} `json:"data"`
					} `json:"items"`
				} `json:"featuringV2"`
				DiscoveredOnV2 struct {
					TotalCount int `json:"totalCount"`
					Items      []struct {
						Data struct {
							Typename string `json:"__typename"`
							URI      string `json:"uri"`
							ID       string `json:"id"`
							OwnerV2  struct {
								Data struct {
									Typename string `json:"__typename"`
									Name     string `json:"name"`
								} `json:"data"`
							} `json:"ownerV2"`
							Name        string `json:"name"`
							Description string `json:"description"`
							Images      struct {
								TotalCount int `json:"totalCount"`
								Items      []struct {
									Sources []struct {
										URL    string `json:"url"`
										Width  any    `json:"width"`
										Height any    `json:"height"`
									} `json:"sources"`
								} `json:"items"`
							} `json:"images"`
						} `json:"data"`
					} `json:"items"`
				} `json:"discoveredOnV2"`
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
									URL    string `json:"url"`
									Width  int    `json:"width"`
									Height int    `json:"height"`
								} `json:"sources"`
							} `json:"avatarImage"`
						} `json:"visuals"`
					} `json:"items"`
				} `json:"relatedArtists"`
			} `json:"relatedContent"`
			Goods struct {
				Events struct {
					UserLocation struct {
						Name string `json:"name"`
					} `json:"userLocation"`
					Concerts struct {
						TotalCount int   `json:"totalCount"`
						Items      []any `json:"items"`
						PagingInfo struct {
							Limit int `json:"limit"`
						} `json:"pagingInfo"`
					} `json:"concerts"`
				} `json:"events"`
				Merch struct {
					Items []any `json:"items"`
				} `json:"merch"`
			} `json:"goods"`
		} `json:"artistUnion"`
	} `json:"data"`
	Extensions struct {
	} `json:"extensions"`
}
