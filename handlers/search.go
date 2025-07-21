package handlers

import (
	"encoding/json"
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

	// Llamar a iTunes siempre con el query completo
	itunesResults, err := services.SearchFromiTunes(query)
	if err != nil {
		http.Error(w, "Error buscando en iTunes: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Llamar a ChartLyrics solo si artist y song NO están vacíos
	var chartResults []models.Song
	if artist != "" && song != "" {
		chartResults, err = services.SearchChart(artist, song)
		if err != nil {
			// Solo logueamos el error, no detenemos todo el flujo
			chartResults = []models.Song{}
		}
	}

	// Consolidar resultados
	allResults := append(itunesResults, chartResults...)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allResults)
}
