package server

import (
	"net/http"
	"net/http/httputil"
	"strings"

	"iris/src/config"
	"iris/src/routing"
)

// Creates a reverse proxy for the matched route, using Rewrite to handle
// path rewriting, header forwarding and header injection.
func buildProxy(matched *routing.MatchedRoute, originalReq *http.Request) *httputil.ReverseProxy {
	route := matched.Route

	proxy := &httputil.ReverseProxy{
		Rewrite: func(pr *httputil.ProxyRequest) {
			pr.SetURL(matched.TargetURL)

			rewritePath(pr.Out, route)
			applyHeaders(pr.Out, originalReq, route)
		},
	}

	return proxy
}

// Replaces the route prefix with rewritePrefix when configured.
func rewritePath(req *http.Request, route *config.Route) {
	if route.RewritePrefix == "" {
		return
	}
	path := req.URL.Path
	if strings.HasPrefix(path, route.Prefix) {
		req.URL.Path = route.RewritePrefix + strings.TrimPrefix(path, route.Prefix)
	}
}

// Sets up the outgoing request headers based on route configuration.
func applyHeaders(req *http.Request, originalReq *http.Request, route *config.Route) {
	if len(route.ForwardHeaders) > 0 {
		headers := http.Header{}
		for _, h := range route.ForwardHeaders {
			if values, ok := originalReq.Header[h]; ok {
				for _, v := range values {
					headers.Add(h, v)
				}
			}
		}
		req.Header = headers
	}

	for k, v := range route.AddHeaders {
		req.Header.Set(k, v)
	}
}

// Checks whether the given HTTP method is allowed by the route.
func methodAllowed(allowed []string, method string) bool {
	for _, m := range allowed {
		if strings.EqualFold(m, method) {
			return true
		}
	}
	return false
}
