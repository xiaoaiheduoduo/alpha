package httpclient

import (
	"go.uber.org/zap"
)

type Logger struct {
	sugar *zap.SugaredLogger
}

func NewLogger(sugar *zap.SugaredLogger) *Logger {
	return &Logger{
		sugar: sugar,
	}
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.sugar.Errorf(format, v)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.sugar.Warnf(format, v)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.sugar.Debugf(format, v)
}
