package logger

import (
	"github.com/sirupsen/logrus"
)

var std = logrus.New()

var Info = std.Info
var Warn = std.Warn
var Error = std.Error

var Infof = std.Infof
var Warnf = std.Warnf

func init() {
	// std.SetFormatter(&logrus.JSONFormatter{})
}

func SetLevel(level string) {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = logrus.DebugLevel
	}
	std.SetLevel(lvl)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return std.WithFields(fields)
}
