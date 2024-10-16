package controllers

import (
	"encoding/json"
	"github.com/chamodshehanka/better-call-place/internal/services"
	"github.com/rs/zerolog/log"
	"net/http"
)

func PlaceSuggestionsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "Query parameter is required", http.StatusBadRequest)
		return
	}

	suggestions, err := services.FetchPlaceSuggestions(query)
	if err != nil {
		log.Err(err).Msg("Error fetching place suggestions")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(suggestions)
	if err != nil {
		return
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Better Call Place!"))
}

func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Better Call Place API is OK"))
}
