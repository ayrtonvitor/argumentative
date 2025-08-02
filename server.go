package main

import (
	"log"
	"net/http"
	"os"
)

func (cfg *apiConfig) startHttpServer() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("ERROR: PORT environment variable is not set")
	}
	mux := http.NewServeMux()
	registerEndpoints(mux, cfg)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("INFO: Serving on port %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
