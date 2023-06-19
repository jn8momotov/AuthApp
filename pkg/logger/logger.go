package logger

import "github.com/sirupsen/logrus"

func Debug(msg ...interface{}) {
	logrus.Debug(msg)
}

func Error(msg ...interface{}) {
	logrus.Error(msg)
}

func Info(msg ...interface{}) {
	logrus.Info(msg)
}

func Warn(msg ...interface{}) {
	logrus.Warn(msg)
}
