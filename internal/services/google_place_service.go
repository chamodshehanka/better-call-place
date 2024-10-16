package services

import (
	"encoding/json"
	"fmt"
	"github.com/chamodshehanka/better-call-place/internal/configs"
	"io/ioutil"
	"net/http"
)

const googlePlacesAPI = "https://maps.googleapis.com/maps/api/place/autocomplete/json"

type PlaceSuggestion struct {
	Description string `json:"description"`
}

type PlacesResponse struct {
	Predictions []PlaceSuggestion `json:"predictions"`
}

func FetchPlaceSuggestions(query string) ([]PlaceSuggestion, error) {
	apiKey := configs.GetConfig().GooglePlaceAPIKey
	url := fmt.Sprintf("%s?input=%s&key=%s", googlePlacesAPI, query, apiKey)

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
