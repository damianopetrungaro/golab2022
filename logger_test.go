package workshop

import (
	"bytes"
	"testing"
)

func TestStdLogger_Debug(t *testing.T) {
	w := &bytes.Buffer{}
	logger := New(w)
	logger.Debug("This is a debug message")
	if got := w.String(); got != "DEBUG: This is a debug message" {
		t.Fatalf("could not match message: %s", got)
	}
}
func TestStdLogger_Info(t *testing.T) {
	w := &bytes.Buffer{}
	logger := New(w)
	logger.Info("This is a info message")
	if got := w.String(); got != "INFO: This is a info message" {
		t.Fatalf("could not match message: %s", got)
	}
}
func TestStdLogger_Warning(t *testing.T) {
	w := &bytes.Buffer{}
	logger := New(w)
	logger.Warning("This is a warning message")
	if got := w.String(); got != "WARNING: This is a warning message" {
		t.Fatalf("could not match message: %s", got)
	}
}
func TestStdLogger_Error(t *testing.T) {
	w := &bytes.Buffer{}
	logger := New(w)
	logger.Error("This is a error message")
	if got := w.String(); got != "ERROR: This is a error message" {
		t.Fatalf("could not match message: %s", got)
	}
}
func TestStdLogger_Fatal(t *testing.T) {
	w := &bytes.Buffer{}
	logger := New(w)
	logger.Fatal("This is a fatal message")
	if got := w.String(); got != "FATAL: This is a fatal message" {
		t.Fatalf("could not match message: %s", got)
	}
}
