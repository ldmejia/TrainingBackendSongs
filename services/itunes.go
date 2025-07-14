package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"searchsong/models"
	"time"
)

func SearchFromiTunes(query string) ([]models.Song, error){
	baseUrl := "https://itunes.apple.com/search"
	params := url.Values{}
	params.Set("term", query)
	params.Set("limit", "10")

	fullURL := fmt.Sprintf("%s?%s", baseUrl, params.Encode())

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("error al hacer la solicitud a iTunes %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer la respuesta de iTunes: %w", err)
	}

	var itunes models.ITunesSearchResponse 
	if err := json.Unmarshal(body, &itunes); err != nil {
		return nil, fmt.Errorf("error al parsear JSON de iTunes: %w", err)
	}

	var songs []models.Song
	for _, track := range itunes.Result {
		songs = append(songs, convertiTunesToSong(track))
	}

	return songs, nil
}