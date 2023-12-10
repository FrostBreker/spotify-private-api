package responseTypes

import "time"

type SearchResponseType struct {
	Data struct {
		SearchV2 struct {
			Albums struct {
				TotalCount int `json:"totalCount"`
				Items      []struct {
					Data struct {
						Typename string `json:"__typename"`
						URI      string `json:"uri"`
						Name     string `json:"name"`
						Artists  struct {
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
							ExtractedColors struct {
								ColorDark struct {
									Hex        string `json:"hex"`
									IsFallback bool   `json:"isFallback"`
								} `json:"colorDark"`
							} `json:"extractedColors"`
						} `json:"coverArt"`
						Date struct {
							Year int `json:"year"`
						} `json:"date"`
					} `json:"data"`
				} `json:"items"`
			} `json:"albums"`
			Artists struct {
				TotalCount int `json:"totalCount"`
				Items      []struct {
					Data struct {
						Typename string `json:"__typename"`
						URI      string `json:"uri"`
						Profile  struct {
							Name     string `json:"name"`
							Verified bool   `json:"verified"`
						} `json:"profile"`
						Visuals struct {
							AvatarImage struct {
								Sources []struct {
									URL    string `json:"url"`
									Width  int    `json:"width"`
									Height int    `json:"height"`
								} `json:"sources"`
								ExtractedColors struct {
									ColorDark struct {
										Hex        string `json:"hex"`
										IsFallback bool   `json:"isFallback"`
									} `json:"colorDark"`
								} `json:"extractedColors"`
							} `json:"avatarImage"`
						} `json:"visuals"`
					} `json:"data"`
				} `json:"items"`
			} `json:"artists"`
			Episodes struct {
				TotalCount int `json:"totalCount"`
				Items      []struct {
					Data struct {
						Typename string `json:"__typename"`
						URI      string `json:"uri"`
						Name     string `json:"name"`
						CoverArt struct {
							Sources []struct {
								URL    string `json:"url"`
								Width  int    `json:"width"`
								Height int    `json:"height"`
							} `json:"sources"`
							ExtractedColors struct {
								ColorDark struct {
									Hex        string `json:"hex"`
									IsFallback bool   `json:"isFallback"`
								} `json:"colorDark"`
							} `json:"extractedColors"`
						} `json:"coverArt"`
						Duration struct {
							TotalMilliseconds int `json:"totalMilliseconds"`
						} `json:"duration"`
						ReleaseDate struct {
							IsoString time.Time `json:"isoString"`
							Precision string    `json:"precision"`
						} `json:"releaseDate"`
						PlayedState struct {
							PlayPositionMilliseconds int    `json:"playPositionMilliseconds"`
							State                    string `json:"state"`
						} `json:"playedState"`
						MediaTypes []string `json:"mediaTypes"`
						PodcastV2  struct {
							Data struct {
								Typename string `json:"__typename"`
								URI      string `json:"uri"`
								Name     string `json:"name"`
								CoverArt struct {
									Sources []struct {
										URL    string `json:"url"`
										Width  int    `json:"width"`
										Height int    `json:"height"`
									} `json:"sources"`
								} `json:"coverArt"`
								MediaType string `json:"mediaType"`
								Publisher struct {
									Name string `json:"name"`
								} `json:"publisher"`
							} `json:"data"`
						} `json:"podcastV2"`
						Description   string `json:"description"`
						ContentRating struct {
							Label string `json:"label"`
						} `json:"contentRating"`
					} `json:"data"`
				} `json:"items"`
			} `json:"episodes"`
			Genres struct {
				TotalCount int           `json:"totalCount"`
				Items      []interface{} `json:"items"`
			} `json:"genres"`
			Playlists struct {
				TotalCount int `json:"totalCount"`
				Items      []struct {
					Data struct {
						Typename    string `json:"__typename"`
						URI         string `json:"uri"`
						Name        string `json:"name"`
						Description string `json:"description"`
						Images      struct {
							Items []struct {
								Sources []struct {
									URL    string      `json:"url"`
									Width  interface{} `json:"width"`
									Height interface{} `json:"height"`
								} `json:"sources"`
								ExtractedColors struct {
									ColorDark struct {
										Hex        string `json:"hex"`
										IsFallback bool   `json:"isFallback"`
									} `json:"colorDark"`
								} `json:"extractedColors"`
							} `json:"items"`
						} `json:"images"`
						Format     string        `json:"format"`
						Attributes []interface{} `json:"attributes"`
						OwnerV2    struct {
							Data struct {
								Typename string `json:"__typename"`
								Name     string `json:"name"`
								URI      string `json:"uri"`
								Username string `json:"username"`
								Avatar   struct {
									Sources []struct {
										URL    string `json:"url"`
										Width  int    `json:"width"`
										Height int    `json:"height"`
									} `json:"sources"`
								} `json:"avatar"`
							} `json:"data"`
						} `json:"ownerV2"`
					} `json:"data"`
				} `json:"items"`
			} `json:"playlists"`
			Podcasts struct {
				TotalCount int `json:"totalCount"`
				Items      []struct {
					Data struct {
						Typename string `json:"__typename"`
						URI      string `json:"uri"`
						Name     string `json:"name"`
						CoverArt struct {
							Sources []struct {
								URL    string `json:"url"`
								Width  int    `json:"width"`
								Height int    `json:"height"`
							} `json:"sources"`
							ExtractedColors struct {
								ColorDark struct {
									Hex        string `json:"hex"`
									IsFallback bool   `json:"isFallback"`
								} `json:"colorDark"`
							} `json:"extractedColors"`
						} `json:"coverArt"`
						Publisher struct {
							Name string `json:"name"`
						} `json:"publisher"`
						MediaType string `json:"mediaType"`
						Topics    struct {
							Items []interface{} `json:"items"`
						} `json:"topics"`
					} `json:"data"`
				} `json:"items"`
			} `json:"podcasts"`
			TracksV2 struct {
				TotalCount int `json:"totalCount"`
				Items      []struct {
					MatchedFields []string `json:"matchedFields"`
					Item          struct {
						Data struct {
							Typename     string `json:"__typename"`
							URI          string `json:"uri"`
							ID           string `json:"id"`
							Name         string `json:"name"`
							AlbumOfTrack struct {
								URI      string `json:"uri"`
								Name     string `json:"name"`
								CoverArt struct {
									Sources []struct {
										URL    string `json:"url"`
										Width  int    `json:"width"`
										Height int    `json:"height"`
									} `json:"sources"`
									ExtractedColors struct {
										ColorDark struct {
											Hex        string `json:"hex"`
											IsFallback bool   `json:"isFallback"`
										} `json:"colorDark"`
									} `json:"extractedColors"`
								} `json:"coverArt"`
								ID string `json:"id"`
							} `json:"albumOfTrack"`
							Artists struct {
								Items []struct {
									URI     string `json:"uri"`
									Profile struct {
										Name string `json:"name"`
									} `json:"profile"`
								} `json:"items"`
							} `json:"artists"`
							ContentRating struct {
								Label string `json:"label"`
							} `json:"contentRating"`
							Duration struct {
								TotalMilliseconds int `json:"totalMilliseconds"`
							} `json:"duration"`
							Playability struct {
								Playable bool `json:"playable"`
							} `json:"playability"`
							Associations struct {
								AssociatedVideos struct {
									TotalCount int `json:"totalCount"`
								} `json:"associatedVideos"`
							} `json:"associations"`
						} `json:"data"`
					} `json:"item"`
				} `json:"items"`
			} `json:"tracksV2"`
			Users struct {
				TotalCount int `json:"totalCount"`
				Items      []struct {
					Data struct {
						Typename    string `json:"__typename"`
						URI         string `json:"uri"`
						ID          string `json:"id"`
						DisplayName string `json:"displayName"`
						Username    string `json:"username"`
						Avatar      struct {
							Sources []struct {
								URL    string `json:"url"`
								Width  int    `json:"width"`
								Height int    `json:"height"`
							} `json:"sources"`
							ExtractedColors struct {
								ColorDark struct {
									Hex        string `json:"hex"`
									IsFallback bool   `json:"isFallback"`
								} `json:"colorDark"`
							} `json:"extractedColors"`
						} `json:"avatar"`
					} `json:"data"`
				} `json:"items"`
			} `json:"users"`
			TopResults struct {
				ItemsV2 []struct {
					MatchedFields []interface{} `json:"matchedFields"`
					Item          struct {
						Data struct {
							Typename string `json:"__typename"`
							URI      string `json:"uri"`
							ID       string `json:"id"`
							Name     string `json:"name"`
							Profile  struct {
								Name     string `json:"name"`
								Verified bool   `json:"verified"`
							} `json:"profile"`
							Visuals struct {
								AvatarImage struct {
									Sources []struct {
										URL    string `json:"url"`
										Width  int    `json:"width"`
										Height int    `json:"height"`
									} `json:"sources"`
									ExtractedColors struct {
										ColorDark struct {
											Hex        string `json:"hex"`
											IsFallback bool   `json:"isFallback"`
										} `json:"colorDark"`
									} `json:"extractedColors"`
								} `json:"avatarImage"`
							} `json:"visuals"`
							AlbumOfTrack struct {
								URI      string `json:"uri"`
								Name     string `json:"name"`
								CoverArt struct {
									Sources []struct {
										URL    string `json:"url"`
										Width  int    `json:"width"`
										Height int    `json:"height"`
									} `json:"sources"`
									ExtractedColors struct {
										ColorDark struct {
											Hex        string `json:"hex"`
											IsFallback bool   `json:"isFallback"`
										} `json:"colorDark"`
									} `json:"extractedColors"`
								} `json:"coverArt"`
								ID string `json:"id"`
							} `json:"albumOfTrack"`
							Artists struct {
								Items []struct {
									URI     string `json:"uri"`
									Profile struct {
										Name string `json:"name"`
									} `json:"profile"`
								} `json:"items"`
							} `json:"artists"`
							ContentRating struct {
								Label string `json:"label"`
							} `json:"contentRating"`
							Duration struct {
								TotalMilliseconds int `json:"totalMilliseconds"`
							} `json:"duration"`
							Playability struct {
								Playable bool `json:"playable"`
							} `json:"playability"`
							Associations struct {
								AssociatedVideos struct {
									TotalCount int `json:"totalCount"`
								} `json:"associatedVideos"`
							} `json:"associations"`
							CoverArt struct {
								Sources []struct {
									URL    string `json:"url"`
									Width  int    `json:"width"`
									Height int    `json:"height"`
								} `json:"sources"`
								ExtractedColors struct {
									ColorDark struct {
										Hex        string `json:"hex"`
										IsFallback bool   `json:"isFallback"`
									} `json:"colorDark"`
								} `json:"extractedColors"`
							} `json:"coverArt"`
							Date struct {
								Year int `json:"year"`
							} `json:"date"`
						} `json:"data"`
					} `json:"item"`
				} `json:"itemsV2"`
				Featured []struct {
					Data struct {
						Typename    string `json:"__typename"`
						URI         string `json:"uri"`
						Name        string `json:"name"`
						Description string `json:"description"`
						Images      struct {
							Items []struct {
								Sources []struct {
									URL    string      `json:"url"`
									Width  interface{} `json:"width"`
									Height interface{} `json:"height"`
								} `json:"sources"`
								ExtractedColors struct {
									ColorDark struct {
										Hex        string `json:"hex"`
										IsFallback bool   `json:"isFallback"`
									} `json:"colorDark"`
								} `json:"extractedColors"`
							} `json:"items"`
						} `json:"images"`
						Format     string `json:"format"`
						Attributes []struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"attributes"`
						OwnerV2 struct {
							Data struct {
								Typename string `json:"__typename"`
								Name     string `json:"name"`
								URI      string `json:"uri"`
								Username string `json:"username"`
								Avatar   struct {
									Sources []struct {
										URL    string `json:"url"`
										Width  int    `json:"width"`
										Height int    `json:"height"`
									} `json:"sources"`
								} `json:"avatar"`
							} `json:"data"`
						} `json:"ownerV2"`
					} `json:"data"`
				} `json:"featured"`
			} `json:"topResults"`
		} `json:"searchV2"`
	} `json:"data"`
	Extensions struct {
		RequestIds struct {
			SearchV2 struct {
				SearchAPI string `json:"search-api"`
			} `json:"/searchV2"`
			SearchV2TopResults struct {
				SearchAPI string `json:"search-api"`
			} `json:"/searchV2/topResults"`
		} `json:"requestIds"`
	} `json:"extensions"`
}

type SearchTopResultsResponseType struct {
	TopResults struct {
		ItemsV2 []struct {
			MatchedFields []interface{} `json:"matchedFields"`
			Item          struct {
				Data struct {
					Typename string `json:"__typename"`
					URI      string `json:"uri"`
					ID       string `json:"id"`
					Name     string `json:"name"`
					Profile  struct {
						Name     string `json:"name"`
						Verified bool   `json:"verified"`
					} `json:"profile"`
					Visuals struct {
						AvatarImage struct {
							Sources []struct {
								URL    string `json:"url"`
								Width  int    `json:"width"`
								Height int    `json:"height"`
							} `json:"sources"`
							ExtractedColors struct {
								ColorDark struct {
									Hex        string `json:"hex"`
									IsFallback bool   `json:"isFallback"`
								} `json:"colorDark"`
							} `json:"extractedColors"`
						} `json:"avatarImage"`
					} `json:"visuals"`
					AlbumOfTrack struct {
						URI      string `json:"uri"`
						Name     string `json:"name"`
						CoverArt struct {
							Sources []struct {
								URL    string `json:"url"`
								Width  int    `json:"width"`
								Height int    `json:"height"`
							} `json:"sources"`
							ExtractedColors struct {
								ColorDark struct {
									Hex        string `json:"hex"`
									IsFallback bool   `json:"isFallback"`
								} `json:"colorDark"`
							} `json:"extractedColors"`
						} `json:"coverArt"`
						ID string `json:"id"`
					} `json:"albumOfTrack"`
					Artists struct {
						Items []struct {
							URI     string `json:"uri"`
							Profile struct {
								Name string `json:"name"`
							} `json:"profile"`
						} `json:"items"`
					} `json:"artists"`
					ContentRating struct {
						Label string `json:"label"`
					} `json:"contentRating"`
					Duration struct {
						TotalMilliseconds int `json:"totalMilliseconds"`
					} `json:"duration"`
					Playability struct {
						Playable bool `json:"playable"`
					} `json:"playability"`
					Associations struct {
						AssociatedVideos struct {
							TotalCount int `json:"totalCount"`
						} `json:"associatedVideos"`
					} `json:"associations"`
					CoverArt struct {
						Sources []struct {
							URL    string `json:"url"`
							Width  int    `json:"width"`
							Height int    `json:"height"`
						} `json:"sources"`
						ExtractedColors struct {
							ColorDark struct {
								Hex        string `json:"hex"`
								IsFallback bool   `json:"isFallback"`
							} `json:"colorDark"`
						} `json:"extractedColors"`
					} `json:"coverArt"`
					Date struct {
						Year int `json:"year"`
					} `json:"date"`
				} `json:"data"`
			} `json:"item"`
		} `json:"itemsV2"`
		Featured []struct {
			Data struct {
				Typename    string `json:"__typename"`
				URI         string `json:"uri"`
				Name        string `json:"name"`
				Description string `json:"description"`
				Images      struct {
					Items []struct {
						Sources []struct {
							URL    string      `json:"url"`
							Width  interface{} `json:"width"`
							Height interface{} `json:"height"`
						} `json:"sources"`
						ExtractedColors struct {
							ColorDark struct {
								Hex        string `json:"hex"`
								IsFallback bool   `json:"isFallback"`
							} `json:"colorDark"`
						} `json:"extractedColors"`
					} `json:"items"`
				} `json:"images"`
				Format     string `json:"format"`
				Attributes []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"attributes"`
				OwnerV2 struct {
					Data struct {
						Typename string `json:"__typename"`
						Name     string `json:"name"`
						URI      string `json:"uri"`
						Username string `json:"username"`
						Avatar   struct {
							Sources []struct {
								URL    string `json:"url"`
								Width  int    `json:"width"`
								Height int    `json:"height"`
							} `json:"sources"`
						} `json:"avatar"`
					} `json:"data"`
				} `json:"ownerV2"`
			} `json:"data"`
		} `json:"featured"`
	} `json:"topResults"`
	Extensions struct {
		RequestIds struct {
			SearchV2 struct {
				SearchAPI string `json:"search-api"`
			} `json:"/searchV2"`
			SearchV2TopResults struct {
				SearchAPI string `json:"search-api"`
			} `json:"/searchV2/topResults"`
		} `json:"requestIds"`
	} `json:"extensions"`
}
