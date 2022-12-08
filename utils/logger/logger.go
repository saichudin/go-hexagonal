package logger

import (
	"os"

	"bitbucket.org/klopos/majoo-logger/logger"
)

var Logger logger.ILogger

func MajooLog() logger.ILogger {
	majooLog, err := logger.NewLogger(
		logger.WithAppName(os.Getenv("APP_NAME")),
		logger.WithLevel(logger.Level.DEBUG),
		logger.WithOutput(logger.Output.JSON),
	)
	if err != nil {
		panic(err)
	}
	return majooLog
}

func LoadLogger() {
	Logger = MajooLog()
}
