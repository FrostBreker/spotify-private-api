package models

// ArtistResponse represents the response from fetching an artist.
type ArtistResponse struct {
	Data struct {
		ArtistUnion ArtistUnion `json:"artistUnion"`
	} `json:"data"`
	Extensions struct{} `json:"extensions"`
}

// ArtistUnion contains the main artist data.
type ArtistUnion struct {
	Typename    string `json:"__typename"`
	ID          string `json:"id"`
	URI         string `json:"uri"`
	Saved       bool   `json:"saved"`
	SharingInfo struct {
		ShareURL string `json:"shareUrl"`
		ShareID  string `json:"shareId"`
	} `json:"sharingInfo"`
	Profile ArtistProfile `json:"profile"`
	Visuals ArtistVisuals `json:"visuals"`
	Stats   ArtistStats   `json:"stats"`
}

// ArtistProfile contains artist profile information.
type ArtistProfile struct {
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
		ItemV2 struct{} `json:"itemV2"`
		Item   struct {
			URI    string `json:"uri"`
			Name   string `json:"name"`
			Images struct {
				Items []struct {
					Sources []ImageSource `json:"sources"`
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
						Sources []ImageSource `json:"sources"`
					} `json:"items"`
				} `json:"images"`
			} `json:"data"`
		} `json:"items"`
	} `json:"playlistsV2"`
}

// ArtistVisuals contains artist visual assets.
type ArtistVisuals struct {
	Gallery struct {
		Items []interface{} `json:"items"`
	} `json:"gallery"`
	AvatarImage struct {
		Sources         []ImageSource `json:"sources"`
		ExtractedColors struct {
			ColorRaw struct {
				Hex string `json:"hex"`
			} `json:"colorRaw"`
		} `json:"extractedColors"`
	} `json:"avatarImage"`
	HeaderImage struct {
		Sources         []ImageSource `json:"sources"`
		ExtractedColors struct {
			ColorRaw struct {
				Hex string `json:"hex"`
			} `json:"colorRaw"`
		} `json:"extractedColors"`
	} `json:"headerImage"`
}

// ArtistStats contains artist statistics.
type ArtistStats struct {
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
}
