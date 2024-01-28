package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

type Logger struct {
	log zerolog.Logger
}

func NewLogger() *Logger {
	return &Logger{
		log: zerolog.New(os.Stdout),
	}
}

func (l *Logger) Debug(message string) {
	log.WithLevel(zerolog.DebugLevel).Str("message", message).Send()
}

func (l *Logger) Info(message string) {
	log.WithLevel(zerolog.InfoLevel).Str("message", message).Send()
}

func (l *Logger) Warn(message string) {
	log.WithLevel(zerolog.WarnLevel).Str("message", message).Send()
}

func (l *Logger) Error(message string) {
	log.WithLevel(zerolog.ErrorLevel).Str("message", message).Send()
}

func (l *Logger) Fatal(message string) {
	log.WithLevel(zerolog.FatalLevel).Str("message", message).Send()
	panic(message)
}
