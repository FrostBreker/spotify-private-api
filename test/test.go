package main

import (
	"fmt"

	spotifyprivateapi "github.com/FrostBreker/spotify-private-api"
)

func main() {
	// Create a new client object with the default options.
	client := spotifyprivateapi.NewClient(spotifyprivateapi.Client{
		Debug: true,
	})

	track, err := client.FetchTrack("3n3Ppam7vgaVa1iaRUc9Lp")
	if err != nil {
		fmt.Println(err)
	}
	playCount, err := client.FetchTrackPlayCount("3n3Ppam7vgaVa1iaRUc9Lp")
	if err != nil {
		fmt.Println(err)
	}
	album, err := client.FetchAlbum(track.Data.TrackUnion.AlbumOfTrack.ID)
	if err != nil {
		fmt.Println(err)
	}

	artist, err := client.FetchArtist(track.Data.TrackUnion.FirstArtist.Items[0].ID)
	if err != nil {
		fmt.Println(err)
	}

	search, err := client.Search("kerc")
	if err != nil {
		fmt.Println(err)
	}

	top, err := client.SearchTopResult("kerc")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(track.Data.TrackUnion.Name)
	fmt.Println(playCount)
	fmt.Println(album.Data.AlbumUnion.Name)
	fmt.Println(artist.Data.ArtistUnion.Visuals.AvatarImage.Sources[0].URL)
	fmt.Println(search.Data.SearchV2.Albums.Items)
	fmt.Println("--------------------------------------")
	fmt.Println(top.TopResults.Featured)
	fmt.Println("--------------------------------------")
	fmt.Println(top.TopResults.ItemsV2)
}
