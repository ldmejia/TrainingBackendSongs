package models

import "encoding/xml"

type ChartLyricsSearchResponse struct {
	XMLNAME xml.Name `xml:"ArrayOfSearchLyricResult"`
	Results []ChartLyricsSearchResult `xml:"SearchLyric"`
}

type ChartLyricsSearchResult struct {
	SongID string `xml:"TrackId"`
	Artist string `xml:"Artist"`
	Song string `xml:"Song"`
	Album string `xml:"Album"`
}

