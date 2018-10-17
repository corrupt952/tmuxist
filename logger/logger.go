package logger

import (
	"io"
	"log"

	"github.com/hashicorp/logutils"
)

const (
	// LogLevelDebug represents identifier for debug level
	LogLevelDebug = "DEBUG"
	// LogLevelWarn represents identifier for warning level
	LogLevelWarn = "WARN"
	// LogLevelErr represents identifier for error level
	LogLevelErr = "ERROR"
)

var l *Logger

// Logger represents logging.
type Logger struct{}

// Setup setups log level and output stream.
func Setup(w io.Writer) {
	if l == nil {
		l = &Logger{}
	}

	filter := l.levelFilter(w, LogLevelWarn)
	log.SetOutput(filter)
}

func (*Logger) levelFilter(w io.Writer, mLvl string) *logutils.LevelFilter {
	return &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{LogLevelDebug, LogLevelWarn, LogLevelErr},
		MinLevel: logutils.LogLevel(mLvl),
		Writer:   w,
	}
}

func (*Logger) logging(level, msg string) {
	log.Printf("[" + level + "] " + msg)
}

// Debug outputs message at debug level.
func Debug(msg string) {
	l.logging(LogLevelDebug, msg)
}

// Warn outputs message at warning level.
func Warn(msg string) {
	l.logging(LogLevelWarn, msg)
}

// Err outputs message at error level.
func Err(msg string) {
	l.logging(LogLevelErr, msg)
}
