package logger

import (
	"log"
	"sync"

	"github.com/caarlos0/env/v10"
	"go.uber.org/zap"
)

type logger struct {
	IsProd  bool   `env:"IS_PROD" envDefault:"false"`
	LogPath string `env:"LOG_PATH" envDefault:"./.tmp/log.txt"`
}

var (
	zapLogger *zap.Logger
	once      sync.Once
)

// NewLogger creates a new zap logger.
//
// No parameters.
// Returns *zap.Logger.
func NewLogger() *zap.Logger {
	once.Do(func() {
		var logger logger
		var config zap.Config

		err := env.Parse(&logger)
		if err != nil {
			log.Println(err.Error())
		}

		if logger.IsProd {
			config = zap.NewProductionConfig()
		} else {
			config = zap.NewDevelopmentConfig()
		}

		config.OutputPaths = []string{"stderr", logger.LogPath}
		zapLogger, err = config.Build()
		if err != nil {
			log.Println(err.Error())
		}
	})

	return zapLogger
}
