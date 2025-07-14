package models

type Song struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Artist   string `json:"artist"`
	Duration string `json:"duration"`
	Album    string `json:"album"`
	Artwork  string `json:"artwork"`
	Price    string `json:"price"`
	Origin   string `json:"origin"`
}
