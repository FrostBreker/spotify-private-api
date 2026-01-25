// Package main provides an example of using the spotify-private-api package.
package main

import (
	"context"
	"fmt"
	"log"

	spotify "github.com/FrostBreker/spotify-private-api"
)

func main() {
	client := spotify.NewClient(spotify.WithDebug(true))
	ctx := context.Background()

	// Fetch a track
	fmt.Println("\n=== TRACK ===")
	track, err := client.FetchTrack(ctx, "4cOdK2wGLETKBW3PvgPWqT")
	if err != nil {
		log.Printf("Error fetching track: %v", err)
	} else {
		fmt.Printf("Name: %s\n", track.Data.TrackUnion.Name)
		fmt.Printf("URI: %s\n", track.Data.TrackUnion.URI)
		fmt.Printf("Duration: %d ms\n", track.Data.TrackUnion.Duration.TotalMilliseconds)
		if len(track.Data.TrackUnion.FirstArtist.Items) > 0 {
			fmt.Printf("Artist: %s\n", track.Data.TrackUnion.FirstArtist.Items[0].Profile.Name)
		}
		fmt.Printf("Album: %s\n", track.Data.TrackUnion.AlbumOfTrack.Name)
	}

	// Fetch an artist
	fmt.Println("\n=== ARTIST ===")
	artist, err := client.FetchArtist(ctx, "4tZwfgrHOc3mvqYlEYSvVi")
	if err != nil {
		log.Printf("Error fetching artist: %v", err)
	} else {
		fmt.Printf("Name: %s\n", artist.Data.ArtistUnion.Profile.Name)
		fmt.Printf("URI: %s\n", artist.Data.ArtistUnion.URI)
		fmt.Printf("Followers: %d\n", artist.Data.ArtistUnion.Stats.Followers)
		fmt.Printf("Monthly Listeners: %d\n", artist.Data.ArtistUnion.Stats.MonthlyListeners)
	}

	// Fetch an album
	fmt.Println("\n=== ALBUM ===")
	album, err := client.FetchAlbum(ctx, "2noRn2Aes5aoNVsU6iWThc")
	if err != nil {
		log.Printf("Error fetching album: %v", err)
	} else {
		fmt.Printf("Name: %s\n", album.Data.AlbumUnion.Name)
		fmt.Printf("URI: %s\n", album.Data.AlbumUnion.URI)
		if len(album.Data.AlbumUnion.Artists.Items) > 0 {
			fmt.Printf("Artist: %s\n", album.Data.AlbumUnion.Artists.Items[0].Profile.Name)
		}
		fmt.Printf("Total Tracks: %d\n", album.Data.AlbumUnion.Tracks.TotalCount)
	}
}
