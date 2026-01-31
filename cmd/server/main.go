package main

import (
	"candipack-pdf/configs"
	"candipack-pdf/internal/handlers"
	"candipack-pdf/internal/middleware"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config := configs.Load()
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("POST /resume", handlers.HandleResume)
	mux.HandleFunc("POST /cover-letter", handlers.HandleCoverLetter)
	mux.HandleFunc("GET /templates", handlers.HandleTemplates)
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	// Middleware
	handler := middleware.CORS(middleware.APIKey(mux, config.APIKey))

	log.Printf("Server running on :%d", config.Port)
	log.Fatal(http.ListenAndServe(":"+fmt.Sprint(config.Port), handler))
}