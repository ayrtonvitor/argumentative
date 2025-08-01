package main

import "net/http"

func (cfg *apiConfig) handleThesisCreation(w http.ResponseWriter, r *http.Request) {
	cfg.DB.CreateThesis(r.Context(), "A wild thesis appears")
	w.WriteHeader(http.StatusOK)
}
