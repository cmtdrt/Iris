package routing

import (
	"net/http"
	"net/url"
	"strings"

	"iris/src/config"
)

// MatchRoute returns the target URL for the incoming request, or nil if no route matches.
func MatchRoute(cfg config.Config, r *http.Request) *url.URL {
	path := r.URL.Path

	// Select the first route whose prefix matches the path.
	for _, route := range cfg.Routes {
		if strings.HasPrefix(path, route.Prefix) {
			targetURL, err := url.Parse(route.Target)
			if err != nil {
				// If the route configuration is invalid, skip it.
				continue
			}
			return targetURL
		}
	}

	return nil
}

