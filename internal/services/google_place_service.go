package services

import (
	"encoding/json"
	"fmt"
	"github.com/chamodshehanka/better-call-place/internal/configs"
	"io/ioutil"
	"net/http"
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Circle struct {
	Center Location `json:"center"`
	Radius float64  `json:"radius"`
}

type LocationBias struct {
	Circle Circle `json:"circle"`
}

type SuggestionRequest struct {
	Input        string       `json:"input"`
	LocationBias LocationBias `json:"locationBias"`
}

type PlaceSuggestion struct {
	Description string `json:"description"`
}

type PlacesResponse struct {
	Predictions []PlaceSuggestion `json:"predictions"`
}

func FetchPlaceSuggestions(apiURL, query string) ([]PlaceSuggestion, error) {
	apiKey := configs.GetConfig().GooglePlaceAPIKey
	url := fmt.Sprintf("%s?input=%s&key=%s", apiURL, query, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var placesResponse PlacesResponse
	if err := json.Unmarshal(body, &placesResponse); err != nil {
		return nil, err
	}

	return placesResponse.Predictions, nil
}
