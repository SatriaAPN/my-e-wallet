package utils

import (
	log "github.com/sirupsen/logrus"
)

type LoggerfDataProvider interface {
	GetFields() map[string]interface{}
	GetInfo() string
}

type Logger interface {
	Errorf(ldp LoggerfDataProvider)
	// Fatalf(format string, args ...interface{})
	// Fatal(args ...interface{})
	Infof(ldp LoggerfDataProvider)
	// Info( args ...interface{})
	// Warnf(format string, args ...interface{})
	// Debugf(format string, args ...interface{})
	// Debug(args ...interface{})
}

var logUnit Logger

func GetLogger() Logger {
	if logUnit == nil {
		logUnit = newLogger()
	}

	return logUnit
}

func newLogger() Logger {
	return &loggerWrapper{}
}

type loggerWrapper struct {
}

func (l *loggerWrapper) Errorf(ldp LoggerfDataProvider) {
	log.WithFields(ldp.GetFields()).Error(ldp.GetInfo())
}

func (l *loggerWrapper) Infof(ldp LoggerfDataProvider) {
	log.WithFields(ldp.GetFields()).Info(ldp.GetInfo())
}
