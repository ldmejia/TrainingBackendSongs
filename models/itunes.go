package models 

type ITunesSearchResponse struct {
	Result []ITunesSong `json:"results"`
}

type ITunesSong struct {
	TrackID int `json:"trackId"`
	TrackName string `json:"trackName"`
	ArtistName string `json:"artistName"`
	CollectionName string `json:"collectionName"`
	TrackTimeMillis int `json:"trackTimeMillis"`
	ArtWorkUrl100 string `json:"artworkUrl100"`
	TrackPrice float64 `json:"trackPrice"`
	Currency string `json:"currency"`
}