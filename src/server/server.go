package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"iris/src/config"
	"iris/src/routing"
)

// Start starts the HTTP server for the API gateway.
func Start(cfg config.Config) error {
	addr := fmt.Sprintf(":%d", cfg.Port)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Health check endpoint.
		if r.URL.Path == "/health" {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
			return
		}

		targetURL := routing.MatchRoute(cfg, r)
		if targetURL == nil {
			http.NotFound(w, r)
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(targetURL)
		proxy.ServeHTTP(w, r)
	})

	log.Printf("Starting Iris API gateway on %s\n", addr)
	return http.ListenAndServe(addr, handler)
}
