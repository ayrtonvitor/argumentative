package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/ayrtonvitor/argumentative/internal/database"
)

func (cfg *apiConfig) handleThesisCreation(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not decode request", err)
		return
	}

	created, err := cfg.queries.CreateThesis(r.Context(), database.CreateThesisParams{
		Title:       params.Title,
		Description: sql.NullString{String: params.Description, Valid: true},
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Colud not create thesis", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, created)
}
