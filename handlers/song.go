package handlers

import (
	"encoding/json"
	"net/http"
	"searchsong/services"
)

func GetSavedSongsHandler(w http.ResponseWriter, r *http.Request) {
	songs, err := services.GetAllSongs()
	if err != nil {
		http.Error(w, "Error obteniendo canciones: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(songs)
}
