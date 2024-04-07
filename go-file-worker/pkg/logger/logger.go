package logger

import (
	"app/pkg/runmode"

	"github.com/sirupsen/logrus"
)

type LoggerService interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Warn(args ...interface{})
}

func getLevelByAppRunMode(runMode string) uint32 {
	switch runMode {
	case runmode.DEV, runmode.TEST:
		return uint32(logrus.DebugLevel)
	case runmode.PROD:
		return uint32(logrus.InfoLevel)
	default:
		return uint32(logrus.DebugLevel)
	}
}
