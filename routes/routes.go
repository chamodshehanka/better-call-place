package routes

import (
	"github.com/chamodshehanka/better-call-place/controllers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/", controllers.HomeHandler)
	http.HandleFunc("/healthz", controllers.HealthzHandler)
	http.HandleFunc("/places", controllers.PlaceSuggestionsHandler)
}
