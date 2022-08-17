package main

import (
	"log"

	"github.com/okankaraduman/golang-test-task/config"
	"github.com/okankaraduman/golang-test-task/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
