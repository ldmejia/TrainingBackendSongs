package main

import (
	"log"
	"net/http"
	"searchsong/handlers"
	"searchsong/middleware"
	"searchsong/services"
)

func main() {
	services.InitDB()

	mux := http.NewServeMux()

	mux.HandleFunc("/login", handlers.LoginHandler)

	mux.HandleFunc("/search", middleware.JWTMiddleware(handlers.SearchHandler))

	mux.HandleFunc("/songs", middleware.JWTMiddleware(handlers.GetSavedSongsHandler))

	log.Println("Servidor corriendo en :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
