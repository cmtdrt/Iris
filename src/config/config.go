package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type CORSConfig struct {
	AllowedOrigins   []string `json:"allowedOrigins,omitempty"`
	AllowedMethods   []string `json:"allowedMethods,omitempty"`
	AllowedHeaders   []string `json:"allowedHeaders,omitempty"`
	ExposedHeaders   []string `json:"exposedHeaders,omitempty"`
	AllowCredentials bool     `json:"allowCredentials,omitempty"`
}

type Route struct {
	Prefix         string            `json:"prefix"`
	Target         string            `json:"target"`
	Methods        []string          `json:"methods,omitempty"`
	ForwardHeaders []string          `json:"forwardHeaders,omitempty"`
	AddHeaders     map[string]string `json:"addHeaders,omitempty"`
	TimeoutMs      int               `json:"timeoutMs,omitempty"`
	RewritePrefix  string            `json:"rewritePrefix,omitempty"`
	CORS           *CORSConfig       `json:"corsConfig,omitempty"`
}

type Config struct {
	Port   int     `json:"port"`
	Routes []Route `json:"routes"`
}

func LoadConfig(filename string) (Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

// Magnificient toString
func (c Config) String() string {
	var b strings.Builder

	b.WriteString("Config:\n")
	b.WriteString(fmt.Sprintf("  Port: %d\n", c.Port))
	b.WriteString("  Routes:\n")

	for _, r := range c.Routes {
		b.WriteString("    - Prefix: ")
		b.WriteString(r.Prefix)
		b.WriteString("\n")
		b.WriteString("      Target: ")
		b.WriteString(r.Target)
		b.WriteString("\n")
	}

	return b.String()
}
