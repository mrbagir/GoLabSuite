package logger

import (
	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	logger *logrus.Entry
}

func NewLogrusLogger() Logger {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	return &LogrusLogger{
		logger: logrus.NewEntry(l),
	}
}

func (l *LogrusLogger) Info(msg string) {
	l.logger.Info(msg)
}

func (l *LogrusLogger) WithFields(fields map[string]any) Logger {
	return &LogrusLogger{
		logger: l.logger.WithFields(fields),
	}
}
