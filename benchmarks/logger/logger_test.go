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
Benchmark_WorkshopLogger/simple_info_log-12             19323600               114.3 ns/op           240 B/op          4 allocs/op
Benchmark_WorkshopLogger/simple_info_log-12             22000051               110.4 ns/op           240 B/op          4 allocs/op
Benchmark_WorkshopLogger/simple_info_log-12             21835646               112.4 ns/op           240 B/op          4 allocs/op
Benchmark_WorkshopLogger/simple_info_log-12             20922128               119.7 ns/op           240 B/op          4 allocs/op
Benchmark_WorkshopLogger/simple_info_log-12             18388087               135.8 ns/op           240 B/op          4 allocs/op
Benchmark_WorkshopLogger/log_info_with_5_fields-12       1372326              1757 ns/op            4084 B/op         52 allocs/op
Benchmark_WorkshopLogger/log_info_with_5_fields-12       1280415              1830 ns/op            4084 B/op         52 allocs/op
Benchmark_WorkshopLogger/log_info_with_5_fields-12       1000000              3286 ns/op            4083 B/op         52 allocs/op
Benchmark_WorkshopLogger/log_info_with_5_fields-12       1032081              1960 ns/op            4084 B/op         52 allocs/op
Benchmark_WorkshopLogger/log_info_with_5_fields-12       1346217              1849 ns/op            4085 B/op         52 allocs/op
Benchmark_LogrusLogger/simple_info_log-12                1000000              2169 ns/op             513 B/op         16 allocs/op
Benchmark_LogrusLogger/simple_info_log-12                1000000              2079 ns/op             513 B/op         16 allocs/op
Benchmark_LogrusLogger/simple_info_log-12                1000000              2085 ns/op             513 B/op         16 allocs/op
Benchmark_LogrusLogger/simple_info_log-12                1000000              2070 ns/op             513 B/op         16 allocs/op
Benchmark_LogrusLogger/simple_info_log-12                1000000              2088 ns/op             513 B/op         16 allocs/op
Benchmark_LogrusLogger/log_info_with_5_fields-12          407560              5747 ns/op            3360 B/op         43 allocs/op
Benchmark_LogrusLogger/log_info_with_5_fields-12          415886              5742 ns/op            3360 B/op         43 allocs/op
Benchmark_LogrusLogger/log_info_with_5_fields-12          415995              5682 ns/op            3360 B/op         43 allocs/op
Benchmark_LogrusLogger/log_info_with_5_fields-12          410802              5889 ns/op            3360 B/op         43 allocs/op
Benchmark_LogrusLogger/log_info_with_5_fields-12          418998              5743 ns/op            3360 B/op         43 allocs/op
Benchmark_ZapLogger/simple_info_log-12                  20071482               138.2 ns/op             0 B/op          0 allocs/op
Benchmark_ZapLogger/simple_info_log-12                  20010974               129.6 ns/op             0 B/op          0 allocs/op
Benchmark_ZapLogger/simple_info_log-12                  18857326               136.5 ns/op             0 B/op          0 allocs/op
Benchmark_ZapLogger/simple_info_log-12                  18673839               126.5 ns/op             0 B/op          0 allocs/op
Benchmark_ZapLogger/simple_info_log-12                  18939207               131.7 ns/op             0 B/op          0 allocs/op
Benchmark_ZapLogger/log_info_with_5_fields-12            1235990              1933 ns/op            6733 B/op         30 allocs/op
Benchmark_ZapLogger/log_info_with_5_fields-12            1215138              1908 ns/op            6733 B/op         30 allocs/op
Benchmark_ZapLogger/log_info_with_5_fields-12            1290912              1825 ns/op            6733 B/op         30 allocs/op
Benchmark_ZapLogger/log_info_with_5_fields-12            1332651              1824 ns/op            6733 B/op         30 allocs/op
Benchmark_ZapLogger/log_info_with_5_fields-12            1279083              2049 ns/op            6733 B/op         30 allocs/op

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
					With("1", "value 1").
					With("2", "value 2").
					With("3", "value 3").
					With("4", "value 4").
					With("5", "value 5").
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
