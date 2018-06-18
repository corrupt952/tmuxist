package main

import (
	"os"
	"log"
	"github.com/hashicorp/logutils"
)

// logging
type Logger struct{}

var logger *Logger = &Logger{}

func (*Logger) logging(level, msg string) {
	log.Printf("[" + level + "] " + msg)
}

func (l *Logger) debug(msg string) {
	l.logging("DEBUG", msg)
}

func (l *Logger) info(msg string) {
	l.logging("INFO", msg)
}

func (l *Logger) warn(msg string) {
	l.logging("WARN", msg)
}

func (l *Logger) err(msg string) {
	l.logging("ERROR", msg)
}

func initLogger() {
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel("WARN"),
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)
}
