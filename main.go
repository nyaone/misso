package main

import (
	"log"
	"misso/global"
	"misso/inits"
)

func main() {
	log.Println("System starting...")

	// Initialize config
	if err := inits.Config(); err != nil {
		log.Fatalln("Failed to load config:", err)
	}

	// Initialize logger
	if err := inits.Logger(); err != nil {
		log.Fatalln("Failed to load logger:", err)
	}

	global.Logger.Info("Logger initialized, switch to here.")

	// Initialize redis
	if err := inits.Redis(); err != nil {
		global.Logger.Fatalf("Failed to load redis: %v", err)
	}

	// Initialize Hydra
	if err := inits.Hydra(); err != nil {
		global.Logger.Fatalf("Failed to load hydra: %v", err)
	}

	// Initializing Gin
	engine := inits.WebEngine()

	global.Logger.Info("Initialization complete.")

	// Start
	global.Logger.Info("Service starting...")
	if err := engine.Run(); err != nil {
		global.Logger.Fatalf("Failed to start service: %v", err)
	}

}
