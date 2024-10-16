package main

import (
	"fmt"
	"github.com/chamodshehanka/better-call-place/internal/configs"
	"github.com/chamodshehanka/better-call-place/internal/middlewares"
	"github.com/chamodshehanka/better-call-place/routes"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()
	
	routes.RegisterRoutes()

	// Wrap the ServeMux with the Logger middleware
	handler := middlewares.Logger(mux)
	port := configs.GetConfig().Port

	log.Info().Msgf("Starting server at port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), handler); err != nil {
		log.Fatal().Err(err).Msg("Server failed to start")
	}
}
