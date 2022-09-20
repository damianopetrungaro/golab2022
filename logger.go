package workshop

import (
	"bytes"
	"context"
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
	Debug(ctx context.Context, msg string)
	Info(ctx context.Context, msg string)
	Warning(ctx context.Context, msg string)
	Error(ctx context.Context, msg string)
	Fatal(ctx context.Context, msg string)
}

type StdLogger struct {
	w          io.Writer
	fields     Fields
	decorators []Decorator
}

func New(w io.Writer, decorators ...Decorator) *StdLogger {
	return &StdLogger{w: w, fields: Fields{}, decorators: decorators}
}

func (s StdLogger) With(fields ...Field) Logger {
	return &StdLogger{w: s.w, decorators: s.decorators, fields: append(s.fields, fields...)}
}

func (s StdLogger) Debug(ctx context.Context, msg string) {
	s.write(ctx, Debug, msg)
}

func (s StdLogger) Info(ctx context.Context, msg string) {
	s.write(ctx, Info, msg)
}

func (s StdLogger) Warning(ctx context.Context, msg string) {
	s.write(ctx, Warning, msg)
}

func (s StdLogger) Error(ctx context.Context, msg string) {
	s.write(ctx, Error, msg)
}

func (s StdLogger) Fatal(ctx context.Context, msg string) {
	s.write(ctx, Fatal, msg)
}

func (s StdLogger) write(ctx context.Context, lvl Level, msg string) {
	fs := s.fields
	for _, dec := range s.decorators {
		fs = append(fs, dec(ctx)...)
	}

	buf := &bytes.Buffer{}
	buf.WriteString(fmt.Sprintf(`{"level":"%s","message":"%s","fields":`, string(lvl), msg))
	fs.Append(buf)
	buf.WriteString(fmt.Sprintf(`}`))

	if _, err := buf.WriteTo(s.w); err != nil {
		panic(err)
	}
}
