package logging

import "log"

// LogRequestReceived logs the basic information about an incoming HTTP request.
func LogRequestReceived(method, path string) {
	log.Printf("Received request: %s - %s", method, path)
}

// LogRequestNotRedirected logs when a request did not match any route.
func LogRequestNotRedirected(path string) {
	log.Printf("Request not redirected: %s", path)
}

// LogRequestRedirected logs when a request is routed to a target backend.
func LogRequestRedirected(fromPath, target string) {
	log.Printf("Request redirected: %s -> %s", fromPath, target)
}

