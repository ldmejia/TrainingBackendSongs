package services

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"searchsong/models"
	"time"
)

func SearchChart(artist, song string) ([]models.Song, error) {
	if artist == "" || song == "" {
		return nil, fmt.Errorf("artist y song no pueden estar vacíos")
	}

	baseURL := "http://api.chartlyrics.com/apiv1.asmx/SearchLyric"
	params := url.Values{}
	params.Set("artist", artist)
	params.Set("song", song)

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("error al hacer request a ChartLyrics: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error leyendo respuesta de ChartLyrics: %w", err)
	}

	if len(body) == 0 {
		return nil, fmt.Errorf("respuesta vacía de ChartLyrics")
	}

	var xmlResponse models.ChartLyricsSearchResponse
	if err := xml.Unmarshal(body, &xmlResponse); err != nil {
		return nil, fmt.Errorf("error parseando XML de ChartLyrics: %w", err)
	}

	var songs []models.Song
	for _, r := range xmlResponse.Results {
		songs = append(songs, ConvertChartLyricstoSong(r))
	}

	return songs, nil
}
