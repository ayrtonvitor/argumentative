package main

import (
	"database/sql"
	"log"
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

	cfg := apiConfig{dbQueries}
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
