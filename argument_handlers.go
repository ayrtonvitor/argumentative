package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/ayrtonvitor/argumentative/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handleArgumentCreation(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Brief       string    `json:"brief"`
		Description string    `json:"description"`
		ThesisId    uuid.UUID `json:"thesisId"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not create argument", err)
		return
	}

	tx, err := cfg.db.Begin()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not create argument", err)
		return
	}
	defer tx.Rollback()

	qtx := cfg.queries.WithTx(tx)
	createArgumentParams := database.CreateArgumentParams{
		Brief:       params.Brief,
		Description: sql.NullString{params.Description, true},
	}
	created, err := qtx.CreateArgument(r.Context(), createArgumentParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not create argument", err)
		return
	}

	insertJointThesisParams := database.InsertTheisJoinArgumentParams{params.ThesisId, created.ID}
	err = qtx.InsertTheisJoinArgument(r.Context(), insertJointThesisParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not create argument", err)
		return
	}
	err = tx.Commit()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not create argument", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, created)
}
