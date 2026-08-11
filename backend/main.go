// Command backend serves the sandbox API that the React frontend calls on
// startup. It deliberately depends on nothing outside the standard library, so
// `go build` works on a fresh machine without downloading any modules.
package main

import (
	"encoding/json"
	"errors"
	"log"
	"net"
	"net/http"
	"time"
)

// addr is the address the frontend expects the API on; see app/page.tsx.
const addr = "0.0.0.0:8000"

type helloResponse struct {
	Message string `json:"message"`
}

// handleHello returns the greeting the frontend renders as its page title.
func handleHello(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(helloResponse{Message: "Welcome"}); err != nil {
		log.Printf("write hello response: %v", err)
	}
}

// withCORS allows the frontend, served from a different port in development, to
// call the API from the browser.
func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/hello", handleHello)
	return mux
}

func main() {
	// Bind before logging, so a port already in use reports the real error
	// instead of a "listening" line the server never lived up to.
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to bind %s: %v", addr, err)
	}

	srv := &http.Server{
		Handler:           withCORS(newMux()),
		ReadHeaderTimeout: 10 * time.Second,
	}

	log.Printf("Backend listening on http://%s", addr)
	if err := srv.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("server error: %v", err)
	}
}
