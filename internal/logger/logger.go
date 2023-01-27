package logger

import (
	"fmt"

	"github.com/jannotti-glaucio/timescale-assignment/internal/excepts"
	"go.uber.org/zap"
)

func Init() error {
	singleton, err := zap.NewDevelopment()
	if err != nil {
		return excepts.ThrowException(excepts.ErrorConfiguringLogger, "Error configuring logger")
	}

	zap.ReplaceGlobals(singleton)
	return nil
}

func Clean() error {
	return zap.L().Sync()
}

func Info(message string, args ...interface{}) {
	message = formatMessage(message, args...)
	zap.L().Info(message)
}

func Debug(message string, args ...interface{}) {
	message = formatMessage(message, args...)
	zap.L().Debug(message)
}

func Fatal(message string, args ...interface{}) {
	message = formatMessage(message, args...)
	zap.L().Fatal(message)
}

func formatMessage(message string, args ...interface{}) string {
	if len(args) > 0 {
		return fmt.Sprintf(message, args...)
	} else {
		return message
	}
}
