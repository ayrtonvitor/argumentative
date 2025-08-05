package main

import (
	"encoding/json"
	"net/http"

	"github.com/ayrtonvitor/argumentative/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handleSourceCreation(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Content    string    `json:"content"`
		ArgumentId uuid.UUID `json:"argument_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not decode request", err)
		return
	}

	createSourceParams := database.CreateArgumentSourcesParams{
		Content:    params.Content,
		ArgumentID: params.ArgumentId,
	}
	created, err := cfg.queries.CreateArgumentSources(r.Context(), createSourceParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not add source", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, created)
}
