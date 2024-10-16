package middlewares

import (
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Create a response writer to capture the status code
		ww := &responseWriter{w, http.StatusOK}

		// Process request
		next.ServeHTTP(ww, r)

		// Log request
		end := time.Now()
		latency := end.Sub(start)
		clientIP := r.RemoteAddr
		method := r.Method
		path := r.URL.Path
		statusCode := ww.statusCode

		logEntry := log.Info()
		if statusCode >= 400 {
			logEntry = log.Error()
		}

		logEntry.
			Int("status", statusCode).
			Dur("latency", latency).
			Str("client_ip", clientIP).
			Str("method", method).
			Str("path", path).
			Msg("request details")
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
