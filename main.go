package main

import (
	"log"

	"github.com/hyuti/pocketbase-template/config"
	"github.com/hyuti/pocketbase-template/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
