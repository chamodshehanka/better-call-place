package services

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestFetchPlaceSuggestions(t *testing.T) {
	// Set up a mock server
	mockResponse := PlacesResponse{
		Predictions: []PlaceSuggestion{
			{Description: "Pizza Place 1"},
			{Description: "Pizza Place 2"},
		},
	}
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer mockServer.Close()

	// Override the GooglePlacesAPI constant for testing
	//GooglePlacesAPI = mockServer.URL

	// Set up environment variable for API key
	os.Setenv("GOOGLE_PLACE_API_KEY", "test_api_key")
	defer os.Unsetenv("GOOGLE_PLACE_API_KEY")

	// Call the function
	suggestions, err := FetchPlaceSuggestions(mockServer.URL, "pizza")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check the results
	if len(suggestions) != 2 {
		t.Errorf("Expected 2 suggestions, got %d", len(suggestions))
	}
	if suggestions[0].Description != "Pizza Place 1" {
		t.Errorf("Expected 'Pizza Place 1', got '%s'", suggestions[0].Description)
	}
	if suggestions[1].Description != "Pizza Place 2" {
		t.Errorf("Expected 'Pizza Place 2', got '%s'", suggestions[1].Description)
	}
}
