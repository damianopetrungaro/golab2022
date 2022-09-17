package workshop

import (
	"bytes"
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

type Logger interface {
	With(fields ...Field) Logger
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

func (s StdLogger) With(fields ...Field) Logger {
	return &StdLogger{w: s.w, fields: append(s.fields, fields...)}
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

	buf := &bytes.Buffer{}
	buf.WriteString(fmt.Sprintf(`{"level":"%s","message":"%s","fields":`, string(lvl), msg))
	s.fields.Append(buf)
	buf.WriteString(fmt.Sprintf(`}`))

	if _, err := buf.WriteTo(s.w); err != nil {
		panic(err)
	}
}
