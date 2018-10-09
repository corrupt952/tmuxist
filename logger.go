package main

import (
	"io"
	"log"

	"github.com/hashicorp/logutils"
)

const (
	LogLevelDebug = "DEBUG"
	LogLevelWarn  = "WARN"
	LogLevelErr   = "ERROR"
)

var logger *Logger = &Logger{}

type Logger struct{}

func (l *Logger) setup(w io.Writer) {
	filter := l.levelFilter(w)
	log.SetOutput(filter)
}

func (*Logger) levelFilter(w io.Writer) *logutils.LevelFilter {
	return &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{LogLevelDebug, LogLevelWarn, LogLevelErr},
		MinLevel: logutils.LogLevel(LogLevelWarn),
		Writer:   w,
	}
}

func (*Logger) logging(level, msg string) {
	log.Printf("[" + level + "] " + msg)
}

func (l *Logger) debug(msg string) {
	l.logging(LogLevelDebug, msg)
}

func (l *Logger) warn(msg string) {
	l.logging(LogLevelWarn, msg)
}

func (l *Logger) err(msg string) {
	l.logging(LogLevelErr, msg)
}
