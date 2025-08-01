package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/ayrtonvitor/argumentative/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("ERROR: .env unreadable: %v", err)
	}

	dbQueries := getDbQueries()

	cfg := &apiConfig{dbQueries}

	mux := http.NewServeMux()
	registerEndpoints(mux, cfg)

	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Serving on port 8080")
	log.Fatal(srv.ListenAndServe())
}

func getDbQueries() *database.Queries {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatalln("ERROR: DATABASE_URL environment variable is not set")
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("ERROR: Unnable to cannect to database: %v", err)
	}
	log.Println("INFO: Connected to the database")
	return database.New(db)
}

func registerEndpoints(mux *http.ServeMux, cfg *apiConfig) {
	mux.HandleFunc("POST /api/thesis", cfg.handleThesisCreation)
}
