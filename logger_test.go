package workshop

import (
	"bytes"
	"context"
	"testing"
)

func TestStdLogger_With(t *testing.T) {
	ctx := context.Background()
	w := &bytes.Buffer{}
	logger := New(w)
	loggerWithFields := logger.With(String("name", "workshop"), Int("year", 2022))
	loggerWithFields.Debug(ctx, "This is a debug message with fields")
	if got := w.String(); got != `{"level":"DEBUG","message":"This is a debug message with fields","fields":[{"name":"workshop"},{"year":2022}]}` {
		t.Fatalf("could not match message: %s", got)
	}
	w.Reset()
	logger.Debug(ctx, "This is a debug message without fields")
	if got := w.String(); got != `{"level":"DEBUG","message":"This is a debug message without fields","fields":[]}` {
		t.Fatalf("could not match message: %s", got)
	}
}
func TestStdLogger_Debug(t *testing.T) {
	ctx := context.Background()
	w := &bytes.Buffer{}
	logger := New(w)
	logger.Debug(ctx, "This is a debug message")
	if got := w.String(); got != `{"level":"DEBUG","message":"This is a debug message","fields":[]}` {
		t.Fatalf("could not match message: %s", got)
	}
}
func TestStdLogger_Info(t *testing.T) {
	ctx := context.Background()
	w := &bytes.Buffer{}
	logger := New(w)
	logger.Info(ctx, "This is a info message")
	if got := w.String(); got != `{"level":"INFO","message":"This is a info message","fields":[]}` {
		t.Fatalf("could not match message: %s", got)
	}
}
func TestStdLogger_Warning(t *testing.T) {
	ctx := context.Background()
	w := &bytes.Buffer{}
	logger := New(w)
	logger.Warning(ctx, "This is a warning message")
	if got := w.String(); got != `{"level":"WARNING","message":"This is a warning message","fields":[]}` {
		t.Fatalf("could not match message: %s", got)
	}
}
func TestStdLogger_Error(t *testing.T) {
	ctx := context.Background()
	w := &bytes.Buffer{}
	logger := New(w)
	logger.Error(ctx, "This is a error message")
	if got := w.String(); got != `{"level":"ERROR","message":"This is a error message","fields":[]}` {
		t.Fatalf("could not match message: %s", got)
	}
}
func TestStdLogger_Fatal(t *testing.T) {
	ctx := context.Background()
	w := &bytes.Buffer{}
	logger := New(w)
	logger.Fatal(ctx, "This is a fatal message")
	if got := w.String(); got != `{"level":"FATAL","message":"This is a fatal message","fields":[]}` {
		t.Fatalf("could not match message: %s", got)
	}
}
