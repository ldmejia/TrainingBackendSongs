package main

import (
	"log"
	"net/http"
	"searchsong/handlers"
)

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("/search", handlers.SearchHandler)

	log.Println("Servidor en 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}