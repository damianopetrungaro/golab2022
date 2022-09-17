package workshop

import (
	"encoding/json"
	"fmt"
	"io"
)

type Level string

const (
	Debug   Level = "DEBUG"
	Info    Level = "INFO"
	Warning Level = "WARNING"
	Error   Level = "ERROR"
	Fatal   Level = "FATAL"
)

type (
	Fields []Field
	Field  struct {
		K string
		V any
	}
)

func (f Field) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{f.K: f.V})
}

type Logger interface {
	With(k string, v any) Logger
	Debug(msg string)
	Info(msg string)
	Warning(msg string)
	Error(msg string)
	Fatal(msg string)
}

type StdLogger struct {
	w      io.Writer
	fields Fields
}

func New(w io.Writer) *StdLogger {
	return &StdLogger{w: w, fields: Fields{}}
}

func (s StdLogger) With(k string, v any) Logger {
	return &StdLogger{w: s.w, fields: append(s.fields, Field{K: k, V: v})}
}

func (s StdLogger) Debug(msg string) {
	s.write(Debug, msg)
}

func (s StdLogger) Info(msg string) {
	s.write(Info, msg)
}

func (s StdLogger) Warning(msg string) {
	s.write(Warning, msg)
}

func (s StdLogger) Error(msg string) {
	s.write(Error, msg)
}

func (s StdLogger) Fatal(msg string) {
	s.write(Fatal, msg)
}

func (s StdLogger) write(lvl Level, msg string) {

	type log struct {
		Level   string `json:"level"`
		Message string `json:"message"`
		Fields  Fields `json:"fields"`
	}
	data, err := json.Marshal(log{
		Level:   string(lvl),
		Message: msg,
		Fields:  s.fields,
	})
	if err != nil {
		panic(err)
	}

	_, _ = fmt.Fprint(s.w, string(data))
}
