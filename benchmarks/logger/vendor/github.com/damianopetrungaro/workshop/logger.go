package workshop

import (
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
	Debug(msg string)
	Info(msg string)
	Warning(msg string)
	Error(msg string)
	Fatal(msg string)
}

type StdLogger struct {
	w io.Writer
}

func New(w io.Writer) *StdLogger {
	return &StdLogger{w: w}
}

func (s *StdLogger) Debug(msg string) {
	_, _ = fmt.Fprintf(s.w, "%s: %s", Debug, msg)
}

func (s *StdLogger) Info(msg string) {
	_, _ = fmt.Fprintf(s.w, "%s: %s", Info, msg)
}

func (s *StdLogger) Warning(msg string) {
	_, _ = fmt.Fprintf(s.w, "%s: %s", Warning, msg)
}

func (s *StdLogger) Error(msg string) {
	_, _ = fmt.Fprintf(s.w, "%s: %s", Error, msg)
}

func (s *StdLogger) Fatal(msg string) {
	_, _ = fmt.Fprintf(s.w, "%s: %s", Fatal, msg)
}
