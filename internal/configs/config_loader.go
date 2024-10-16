package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
)

var config Config

var requiredEnvs = [...]string{
	// Server
	"PORT",

	// Google Place API
	"GOOGLE_PLACE_API_KEY",
}

func init() {
	loadEnvFileIfAvailable()

	if err := ensureRequiredEnvsAreAvailable(); err != nil {
		log.Fatal().Msgf("Error loading environment variables: %v", err)
	}

	config = Config{
		Port:              getEnv("PORT"),
		GooglePlaceAPIKey: getEnv("GOOGLE_PLACE_API_KEY"),
	}

	log.Info().Msg("Config loaded successfully!")
}

func GetConfig() *Config {
	return &config
}

func ensureRequiredEnvsAreAvailable() error {
	for _, env := range requiredEnvs {
		if getEnv(env) == "" {
			return fmt.Errorf("fatal: required environment variable '%s' not found", env)
		}
	}
	return nil
}

func loadEnvFileIfAvailable() {
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Fatal().Msgf("Error loading .env file: %v", err)
		}
	} else if !os.IsNotExist(err) {
		log.Fatal().Msgf("Error checking .env file: %v", err)
	}
}

func getEnv(key string) string {
	return os.Getenv(key)
}

func getIntegerEnv(key string) int {
	envStr := getEnv(key)

	intEnv, err := strconv.Atoi(envStr)
	if err != nil {
		log.Err(err).Msgf("Error converting environment variable '%s' to integer: %v", key, err)
		return 0
	}

	return intEnv
}
