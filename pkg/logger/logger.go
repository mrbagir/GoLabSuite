package logger

type Logger interface {
	Info(msg string)
	WithFields(fields map[string]any) Logger
}

func SetLogger(l Logger) Logger {
	return l
}
