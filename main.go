package main

import (
	"log"
	"net/http"
	"searchsong/handlers"
	"searchsong/middleware"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", handlers.LoginHandler)

	mux.HandleFunc("/search", middleware.JWTMiddleware(handlers.SearchHandler))

	log.Println("Servidor corriendo en :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
