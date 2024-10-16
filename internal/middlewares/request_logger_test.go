package middlewares

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogger(t *testing.T) {
	// Set up a buffer to capture log output
	var logBuffer strings.Builder
	log.Logger = zerolog.New(&logBuffer).With().Timestamp().Logger()

	// Create a test handler to wrap with the Logger middleware
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Wrap the test handler with the Logger middleware
	wrappedHandler := Logger(testHandler)

	// Create a test request
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	req.RemoteAddr = "127.0.0.1:12345"

	// Create a test response recorder
	rr := httptest.NewRecorder()

	// Serve the request
	wrappedHandler.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the log output
	logOutput := logBuffer.String()
	if !strings.Contains(logOutput, "request details") {
		t.Errorf("log output does not contain expected message: %v", logOutput)
	}
	if !strings.Contains(logOutput, "status") {
		t.Errorf("log output does not contain status: %v", logOutput)
	}
	if !strings.Contains(logOutput, "latency") {
		t.Errorf("log output does not contain latency: %v", logOutput)
	}
	if !strings.Contains(logOutput, "client_ip") {
		t.Errorf("log output does not contain client_ip: %v", logOutput)
	}
	if !strings.Contains(logOutput, "method") {
		t.Errorf("log output does not contain method: %v", logOutput)
	}
	if !strings.Contains(logOutput, "path") {
		t.Errorf("log output does not contain path: %v", logOutput)
	}
}
