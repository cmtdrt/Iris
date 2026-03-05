package main

import (
	"fmt"
	"log"

	"iris/src/config"
	"iris/src/server"
)

const CONFIG_FILE_NAME = "iris-config.json"

func main() {
	cfg, err := config.LoadConfig(CONFIG_FILE_NAME)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	fmt.Println(cfg.String())

	if err := server.Start(cfg); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
