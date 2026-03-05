package server

import (
	"fmt"
	"log"
	"net/http"

	"iris/src/config"
	"iris/src/logging"
	"iris/src/routing"
)

// Start starts the HTTP server for the API gateway.
func Start(cfg config.Config) error {
	addr := fmt.Sprintf(":%d", cfg.Port)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/", gatewayHandler(cfg))

	log.Printf("Starting Iris API gateway on %s\n", addr)
	return http.ListenAndServe(addr, mux)
}

// gatewayHandler returns the main handler that routes and proxies incoming requests.
func gatewayHandler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logging.LogRequestReceived(r.Method, r.URL.Path)

		matched := routing.MatchRoute(cfg, r)
		if matched == nil {
			logging.LogRequestNotRedirected(r.URL.Path)
			http.NotFound(w, r)
			return
		}

		if len(matched.Route.Methods) > 0 && !methodAllowed(matched.Route.Methods, r.Method) {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		logging.LogRequestRedirected(r.URL.Path, matched.TargetURL.String())

		proxy := buildProxy(matched, r)
		proxy.ServeHTTP(w, r)
	}
}
