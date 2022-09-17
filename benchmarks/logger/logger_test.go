package logger

import (
	"io"
	"testing"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/damianopetrungaro/workshop"
)

/**
go test --bench=. --benchmem -benchtime=2s -count=5
Benchmark_WorkshopLogger/simple_info_log-12             19823806               115.7 ns/op           240 B/op          4 allocs/op
Benchmark_WorkshopLogger/simple_info_log-12             18853231               122.8 ns/op           240 B/op          4 allocs/op
Benchmark_WorkshopLogger/simple_info_log-12             19446843               114.4 ns/op           240 B/op          4 allocs/op
Benchmark_WorkshopLogger/simple_info_log-12             21283146               113.6 ns/op           240 B/op          4 allocs/op
Benchmark_WorkshopLogger/simple_info_log-12             18045288               127.9 ns/op           240 B/op          4 allocs/op
Benchmark_WorkshopLogger/log_info_with_5_fields-12       3223809               754.7 ns/op          1331 B/op         25 allocs/op
Benchmark_WorkshopLogger/log_info_with_5_fields-12       3378568               708.6 ns/op          1331 B/op         25 allocs/op
Benchmark_WorkshopLogger/log_info_with_5_fields-12       3323348               787.1 ns/op          1331 B/op         25 allocs/op
Benchmark_WorkshopLogger/log_info_with_5_fields-12       3282172               751.3 ns/op          1331 B/op         25 allocs/op
Benchmark_WorkshopLogger/log_info_with_5_fields-12       3153771               781.1 ns/op          1332 B/op         25 allocs/op
Benchmark_LogrusLogger/simple_info_log-12                 956430              2137 ns/op             513 B/op         16 allocs/op
Benchmark_LogrusLogger/simple_info_log-12                1000000              2093 ns/op             513 B/op         16 allocs/op
Benchmark_LogrusLogger/simple_info_log-12                1000000              2124 ns/op             513 B/op         16 allocs/op
Benchmark_LogrusLogger/simple_info_log-12                1000000              2107 ns/op             513 B/op         16 allocs/op
Benchmark_LogrusLogger/simple_info_log-12                1000000              2082 ns/op             513 B/op         16 allocs/op
Benchmark_LogrusLogger/log_info_with_5_fields-12          409662              5762 ns/op            3360 B/op         43 allocs/op
Benchmark_LogrusLogger/log_info_with_5_fields-12          418306              5755 ns/op            3360 B/op         43 allocs/op
Benchmark_LogrusLogger/log_info_with_5_fields-12          418279              5759 ns/op            3360 B/op         43 allocs/op
Benchmark_LogrusLogger/log_info_with_5_fields-12          411640              5751 ns/op            3360 B/op         43 allocs/op
Benchmark_LogrusLogger/log_info_with_5_fields-12          418074              5979 ns/op            3360 B/op         43 allocs/op
Benchmark_ZapLogger/simple_info_log-12                  19592696               128.9 ns/op             0 B/op          0 allocs/op
Benchmark_ZapLogger/simple_info_log-12                  17563126               128.0 ns/op             0 B/op          0 allocs/op
Benchmark_ZapLogger/simple_info_log-12                  19791685               127.2 ns/op             0 B/op          0 allocs/op
Benchmark_ZapLogger/simple_info_log-12                  19708780               122.5 ns/op             0 B/op          0 allocs/op
Benchmark_ZapLogger/simple_info_log-12                  19020962               126.4 ns/op             0 B/op          0 allocs/op
Benchmark_ZapLogger/log_info_with_5_fields-12            1308580              1763 ns/op            6733 B/op         30 allocs/op
Benchmark_ZapLogger/log_info_with_5_fields-12            1320468              1884 ns/op            6733 B/op         30 allocs/op
Benchmark_ZapLogger/log_info_with_5_fields-12            1232829              2010 ns/op            6733 B/op         30 allocs/op
Benchmark_ZapLogger/log_info_with_5_fields-12            1271334              1840 ns/op            6733 B/op         30 allocs/op
Benchmark_ZapLogger/log_info_with_5_fields-12            1345098              1921 ns/op            6733 B/op         30 allocs/op
*/
func Benchmark_WorkshopLogger(b *testing.B) {
	b.Run("simple info log", func(b *testing.B) {
		logger := workshop.New(io.Discard)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info("An info event happened here")
			}
		})
	})

	b.Run("log info with 5 fields", func(b *testing.B) {
		logger := workshop.New(io.Discard)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.
					With(workshop.String("1", "value 1")).
					With(workshop.String("2", "value 2")).
					With(workshop.String("3", "value 3")).
					With(workshop.String("4", "value 4")).
					With(workshop.String("5", "value 5")).
					Info("An info event happened here")
			}
		})
	})
}

func Benchmark_LogrusLogger(b *testing.B) {
	b.Run("simple info log", func(b *testing.B) {
		logger := logrus.New()
		logger.Out = io.Discard
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info("An info event happened here")
			}
		})
	})

	b.Run("log info with 5 fields", func(b *testing.B) {
		logger := logrus.New()
		logger.Out = io.Discard
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.
					WithField("1", "value 1").
					WithField("2", "value 2").
					WithField("3", "value 3").
					WithField("4", "value 4").
					WithField("5", "value 5").
					Info("An info event happened here")
			}
		})
	})
}

func Benchmark_ZapLogger(b *testing.B) {
	b.Run("simple info log", func(b *testing.B) {
		core := zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()), &Discarder{}, zap.DebugLevel)
		logger := zap.New(core).WithOptions()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info("An info event happened here")
			}
		})
	})

	b.Run("log info with 5 fields", func(b *testing.B) {
		core := zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()), &Discarder{}, zap.DebugLevel)
		logger := zap.New(core).WithOptions()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.
					With(zap.String("1", "value 1")).
					With(zap.String("2", "value 2")).
					With(zap.String("3", "value 3")).
					With(zap.String("4", "value 4")).
					With(zap.String("5", "value 5")).
					Info("An info event happened here")
			}
		})
	})
}
