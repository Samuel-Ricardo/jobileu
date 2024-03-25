package main

import "github.com/Samuel-Ricardo/jobileu/config"

var logger *config.Logger

func main() {
	logger = config.GetLogger("Main")

	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}
}
