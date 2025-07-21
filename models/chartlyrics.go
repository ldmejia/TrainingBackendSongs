package models

import "encoding/xml"

type ChartLyricsSearchResponse struct {
	XMLName xml.Name                   `xml:"ArrayOfSearchLyricResult"`
	Results []ChartLyricsSearchResult `xml:"SearchLyricResult"`
}

type ChartLyricsSearchResult struct {
	TrackId string `xml:"TrackId"`
	Artist  string `xml:"Artist"`
	Song    string `xml:"Song"`
	Album   string `xml:"Album,omitempty"` 
}

