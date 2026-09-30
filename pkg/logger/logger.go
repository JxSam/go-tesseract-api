package logger

import (
	"fmt"
	"log/slog"
	"os"
)

type Logger struct {
	log *slog.Logger
}

func New() *Logger {
	return &Logger{
		log: slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
	}
}

func (l *Logger) ErrorRequest(msg string, endpoint string, payload any, err error) {
	l.log.Error(
		fmt.Sprintf("%s request failed", msg),
		slog.Any("error", err),
		slog.String("endpoint", endpoint),
		slog.Any("payload", payload),
	)
}

func (l *Logger) ErrorFunc(filename string, function string, err error) {
	l.log.Error(
		"func failed",
		slog.Any("error", err),
		slog.String("filename", filename),
		slog.String("function", function),
	)
}
