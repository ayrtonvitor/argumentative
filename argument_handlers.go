package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handleArgumentCreation(w http.ResponseWriter, r *http.Request) {
	type parameterns struct {
		Brief       string    `json:"brief"`
		Description string    `json:"description"`
		ThesisId    uuid.UUID `json:"thesisId"`
		Sources     []string  `json:"sources"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameterns{}
	err := decoder.Decode(&decoder)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not create argument", err)
		return
	}

	created, err := cfg.queries
}
