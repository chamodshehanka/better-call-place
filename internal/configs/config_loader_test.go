package configs

import (
	"os"
	"testing"
)

func TestLoadEnvFileIfAvailable(t *testing.T) {
	// Set up a temporary .env file for testing
	envContent := "PORT=8080\nGOOGLE_PLACE_API_KEY=test_api_key"
	if err := os.WriteFile(".env", []byte(envContent), 0644); err != nil {
		t.Fatalf("Failed to create .env file: %v", err)
	}
	defer os.Remove(".env")

	loadEnvFileIfAvailable()

	if os.Getenv("PORT") != "8080" {
		t.Errorf("Expected PORT to be '8080', got '%s'", os.Getenv("PORT"))
	}
	if os.Getenv("GOOGLE_PLACE_API_KEY") != "test_api_key" {
		t.Errorf("Expected GOOGLE_PLACE_API_KEY to be 'test_api_key', got '%s'", os.Getenv("GOOGLE_PLACE_API_KEY"))
	}
}

func TestEnsureRequiredEnvsAreAvailable(t *testing.T) {
	os.Setenv("PORT", "8080")
	os.Setenv("GOOGLE_PLACE_API_KEY", "test_api_key")

	if err := ensureRequiredEnvsAreAvailable(); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	os.Unsetenv("PORT")
	if err := ensureRequiredEnvsAreAvailable(); err == nil {
		t.Errorf("Expected error for missing PORT, got nil")
	}
}

func TestGetConfig(t *testing.T) {
	os.Setenv("PORT", "8080")
	os.Setenv("GOOGLE_PLACE_API_KEY", "test_api_key")

	config := GetConfig()

	if config.Port != "8080" {
		t.Errorf("Expected Port to be '8080', got '%s'", config.Port)
	}
	if config.GooglePlaceAPIKey != "test_api_key" {
		t.Errorf("Expected GooglePlaceAPIKey to be 'test_api_key', got '%s'", config.GooglePlaceAPIKey)
	}
}
