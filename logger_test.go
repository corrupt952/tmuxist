package main

import (
	"bytes"
	"log"
	"testing"
)

func TestLevelFilter(t *testing.T) {
	l := Logger{}
	buf := new(bytes.Buffer)
	filter := l.levelFilter(buf)

	logger := log.New(filter, "", 0)
	logger.Print("[WARN] 1")
	logger.Println("[ERROR] 2")
	logger.Println("[DEBUG] 3")
	logger.Println("[WARN] 4")

	actual := buf.String()
	expected := "[WARN] 1\n[ERROR] 2\n[WARN] 4\n"
	AssertEquals(t, actual, expected)
}

func TestLogging(t *testing.T) {
}

func TestDebug(t *testing.T) {
}

func TestInfo(t *testing.T) {
}

func TestWarn(t *testing.T) {
}

func TestErr(t *testing.T) {
}
