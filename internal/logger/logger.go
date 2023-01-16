package logger

import (
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
	if len(args) == 0 {
		zap.L().Info(message)
	} else {
		zap.L().Sugar().Infof(message, args)
	}
}

func Debug(message string, args ...interface{}) {
	if len(args) == 0 {
		zap.L().Debug(message)
	} else {
		zap.L().Sugar().Debugf(message, args)
	}
}

func Fatal(message string, args ...interface{}) {
	if len(args) == 0 {
		zap.L().Fatal(message)
	} else {
		zap.L().Sugar().Fatalf(message, args)
	}
}
