package main

import (
	"github.com/aleksanderpalamar/jobs/config"
	"github.com/aleksanderpalamar/jobs/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Failed to initialize config: %v", err)
		return
	}

	// Initialize router
	router.Initialize()
}
