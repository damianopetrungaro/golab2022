package workshop

import (
	"bytes"
	"testing"
)

func TestStdLogger_With(t *testing.T) {
	w := &bytes.Buffer{}
	logger := New(w)
	loggerWithFields := logger.With("key_1", "value_1").With("key_2", "value_2")
	loggerWithFields.Debug("This is a debug message with fields")
	if got := w.String(); got != `{"level":"DEBUG","message":"This is a debug message with fields","fields":[{"key_1":"value_1"},{"key_2":"value_2"}]}` {
		t.Fatalf("could not match message: %s", got)
	}
	w.Reset()
	logger.Debug("This is a debug message without fields")
	if got := w.String(); got != `{"level":"DEBUG","message":"This is a debug message without fields","fields":[]}` {
		t.Fatalf("could not match message: %s", got)
	}
}
func TestStdLogger_Debug(t *testing.T) {
	w := &bytes.Buffer{}
	logger := New(w)
	logger.Debug("This is a debug message")
	if got := w.String(); got != `{"level":"DEBUG","message":"This is a debug message","fields":[]}` {
		t.Fatalf("could not match message: %s", got)
	}
}
func TestStdLogger_Info(t *testing.T) {
	w := &bytes.Buffer{}
	logger := New(w)
	logger.Info("This is a info message")
	if got := w.String(); got != `{"level":"INFO","message":"This is a info message","fields":[]}` {
		t.Fatalf("could not match message: %s", got)
	}
}
func TestStdLogger_Warning(t *testing.T) {
	w := &bytes.Buffer{}
	logger := New(w)
	logger.Warning("This is a warning message")
	if got := w.String(); got != `{"level":"WARNING","message":"This is a warning message","fields":[]}` {
		t.Fatalf("could not match message: %s", got)
	}
}
func TestStdLogger_Error(t *testing.T) {
	w := &bytes.Buffer{}
	logger := New(w)
	logger.Error("This is a error message")
	if got := w.String(); got != `{"level":"ERROR","message":"This is a error message","fields":[]}` {
		t.Fatalf("could not match message: %s", got)
	}
}
func TestStdLogger_Fatal(t *testing.T) {
	w := &bytes.Buffer{}
	logger := New(w)
	logger.Fatal("This is a fatal message")
	if got := w.String(); got != `{"level":"FATAL","message":"This is a fatal message","fields":[]}` {
		t.Fatalf("could not match message: %s", got)
	}
}
