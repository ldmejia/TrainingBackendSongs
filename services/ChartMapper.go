package services

import (
	"searchsong/models"
	"github.com/google/uuid"
)

func ConvertChartLyricstoSong(result models.ChartLyricsSearchResult) models.Song {

	var defaultValue string = "N/A"
	id := uuid.New().String()

	name := result.Song 
	if name == "" {
		name = defaultValue
	}

	artist := result.Artist
	if artist == "" {
		artist = defaultValue 
	}

	album := result.Album
	if album == "" {
		album = defaultValue
	}

	return models.Song{
		ID: id, 
		Name: name, 
		Artist: artist,
		Duration: defaultValue,
		Album: album,
		Artwork: defaultValue,
		Price: defaultValue,
		Origin: "chartlyrics",
	}
}