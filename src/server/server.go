package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"iris/src/config"
	"iris/src/logging"
	"iris/src/routing"
)

// Start starts the HTTP server for the API gateway.
func Start(cfg config.Config) error {
	addr := fmt.Sprintf(":%d", cfg.Port)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logging.LogRequestReceived(r.Method, r.URL.Path)

		// Health check endpoint.
		if r.URL.Path == "/health" {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
			return
		}

		targetURL := routing.MatchRoute(cfg, r)
		if targetURL == nil {
			logging.LogRequestNotRedirected(r.URL.Path)
			http.NotFound(w, r)
			return
		}

		logging.LogRequestRedirected(r.URL.Path, targetURL.String())

		proxy := httputil.NewSingleHostReverseProxy(targetURL)
		proxy.ServeHTTP(w, r)
	})

	log.Printf("Starting Iris API gateway on %s\n", addr)
	return http.ListenAndServe(addr, handler)
}
