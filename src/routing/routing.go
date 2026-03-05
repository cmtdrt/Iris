package routing

import (
	"net/http"
	"net/url"
	"strings"

	"iris/src/config"
)

// MatchedRoute holds both the matching route and its parsed target URL.
type MatchedRoute struct {
	Route     *config.Route
	TargetURL *url.URL
}

// MatchRoute returns the matching route and target URL for the incoming request, or nil if no route matches.
func MatchRoute(cfg config.Config, r *http.Request) *MatchedRoute {
	path := r.URL.Path

	// Select the first route whose prefix matches the path.
	for i := range cfg.Routes {
		route := &cfg.Routes[i]
		if strings.HasPrefix(path, route.Prefix) {
			targetURL, err := url.Parse(route.Target)
			if err != nil {
				// If the route configuration is invalid, skip it.
				continue
			}

			return &MatchedRoute{
				Route:     route,
				TargetURL: targetURL,
			}
		}
	}

	return nil
}

