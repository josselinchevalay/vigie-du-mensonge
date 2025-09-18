package logger

import (
	"log/slog"
	"os"
)

var globalLogger = initLogger()

func initLogger() Logger {
	return newSlogAdapter(
		slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})))
}

type Field struct {
	Key   string
	Value any
}

type Logger interface {
	Debug(string, ...Field)
	Info(string, ...Field)
	Error(string, ...Field)
}

func Err(err error) Field {
	return Field{
		Key:   "error",
		Value: err,
	}
}

func Any(key string, value any) Field {
	return Field{
		Key:   key,
		Value: value,
	}
}

func Debug(s string, fields ...Field) {
	globalLogger.Debug(s, fields...)
}

func Info(s string, fields ...Field) {
	globalLogger.Info(s, fields...)
}

func Error(s string, fields ...Field) {
	globalLogger.Error(s, fields...)
}
