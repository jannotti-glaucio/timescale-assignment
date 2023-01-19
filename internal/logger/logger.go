package logger

import (
	"fmt"
	"sync"

	"go.uber.org/zap"
)

var (
	once sync.Once
)

func Init() {
	once.Do(func() {
		singleton, _ := zap.NewDevelopment()
		zap.ReplaceGlobals(singleton)
	})
}

func Clean() error {
	return zap.L().Sync()
}

func Info(message string, args ...interface{}) {
	if len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}

	zap.L().Info(message)
}

func Debug(message string, args ...interface{}) {
	if len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}

	zap.L().Debug(message)
}

func Fatal(message string, args ...interface{}) {
	if len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}

	zap.L().Fatal(message)
}
