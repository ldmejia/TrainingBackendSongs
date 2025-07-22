package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"searchsong/models"
	"searchsong/services"
	"strings"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "query string 'q' es requerida", http.StatusBadRequest)
		return
	}

	parts := strings.Split(query, " ")
	artist := ""
	song := ""

	if len(parts) >= 2 {
		artist = parts[0]
		song = strings.Join(parts[1:], " ")
	} else {
		song = query
	}

	itunesResults, err := services.SearchFromiTunes(query)
	if err != nil {
		http.Error(w, "Error buscando en iTunes: "+err.Error(), http.StatusInternalServerError)
		return
	}
	var chartResults []models.Song
	if artist != "" && song != "" {
		chartResults, err = services.SearchChart(artist, song)
		if err != nil {
			chartResults = []models.Song{}
		}
	}

	allResults := append(itunesResults, chartResults...)

	for _, song := range allResults {
		
	if err := services.SaveSong(song); err != nil {
		log.Println("Error guardando canción:", err)
	}
}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allResults)
}
