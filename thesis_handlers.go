package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ayrtonvitor/argumentative/internal/database"
	"github.com/google/uuid"
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

func (cfg *apiConfig) handleGetThesisById(w http.ResponseWriter, r *http.Request) {
	const queryParamName = "thesis_id"
	const idListSep = ","
	rawThesisIds := r.URL.Query().Get(queryParamName)
	if rawThesisIds == "" {
		respondWithError(w, http.StatusBadRequest, "Thesis id must be non-empty", nil)
		return
	}

	thesisIds := make([]uuid.UUID, 0)
	for _, rawThesisId := range strings.Split(rawThesisIds, idListSep) {
		thesisId, err := uuid.Parse(rawThesisId)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Thesis id is malformed", err)
			return
		}
		thesisIds = append(thesisIds, thesisId)
	}

	created, err := cfg.queries.GetThesisById(r.Context(), thesisIds)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not get thesis", err)
		return
	}

	respondWithJSON(w, http.StatusOK, created)
}
