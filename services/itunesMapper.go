package services

import (
	"fmt"
	"searchsong/models"

	"github.com/google/uuid"
)

func convertiTunesToSong(track models.ITunesSong) models.Song {
	minutes := track.TrackTimeMillis / 60000
	seconds := (track.TrackTimeMillis % 60000) / 1000
	duration := fmt.Sprintf("%02d:%02d", minutes, seconds)

	price := fmt.Sprintf("GTQ %.2f", track.TrackPrice)

	return models.Song{
		ID:       uuid.New().String(),
		Name:     track.TrackName,
		Artist:   track.ArtistName,
		Duration: duration,
		Album:    track.CollectionName,
		Artwork:  track.ArtWorkUrl100,
		Price:    price,
		Origin:   "apple",
	}
}
