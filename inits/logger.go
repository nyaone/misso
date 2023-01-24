package inits

import (
	"fmt"
	"go.uber.org/zap"
	"misso/config"
	"misso/global"
)

func Logger() error {
	var err error

	var logger *zap.Logger

	// Prepare logger
	if config.Config.System.Debug {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		return fmt.Errorf("failed to initialize logger: %v", err)
	}

	// Flush logs
	defer logger.Sync() // Unable to handle errors here

	// Sugar it
	global.Logger = logger.Sugar()

	return nil
}
