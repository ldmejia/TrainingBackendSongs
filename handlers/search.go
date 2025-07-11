package handlers

import (
	"encoding/json"
	"net/http"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	if query == "" {
		http.Error(w, "query string 'q' es requerida", http.StatusBadRequest)
		return
	}

	result := map[string]interface{}{
		"message": "Búsqueda recibida correctamente",
		"query":   query,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
