package config

import (
	"encoding/json"
	"os"
)

type Route struct {
	Prefix string `json:"prefix"`
	Target string `json:"target"`
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
