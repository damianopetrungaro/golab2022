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
Benchmark_WorkshopLogger/simple_info_log-12             19857986               115.1 ns/op           240 B/op          4 allocs/op
Benchmark_WorkshopLogger/simple_info_log-12             21409348               112.4 ns/op           240 B/op          4 allocs/op
Benchmark_WorkshopLogger/simple_info_log-12             21454494               111.5 ns/op           240 B/op          4 allocs/op
Benchmark_WorkshopLogger/simple_info_log-12             21650558               118.5 ns/op           240 B/op          4 allocs/op
Benchmark_WorkshopLogger/simple_info_log-12             20154054               125.1 ns/op           240 B/op          4 allocs/op
Benchmark_WorkshopLogger/log_info_with_5_fields-12       1344625              1742 ns/op            4084 B/op         52 allocs/op
Benchmark_WorkshopLogger/log_info_with_5_fields-12       1426080              1720 ns/op            4084 B/op         52 allocs/op
Benchmark_WorkshopLogger/log_info_with_5_fields-12       1355678              1741 ns/op            4084 B/op         52 allocs/op
Benchmark_WorkshopLogger/log_info_with_5_fields-12       1323265              1768 ns/op            4085 B/op         52 allocs/op
Benchmark_WorkshopLogger/log_info_with_5_fields-12       1389840              1795 ns/op            4084 B/op         52 allocs/op
Benchmark_LogrusLogger/simple_info_log-12                 973387              2144 ns/op             513 B/op         16 allocs/op
Benchmark_LogrusLogger/simple_info_log-12                1000000              2071 ns/op             513 B/op         16 allocs/op
Benchmark_LogrusLogger/simple_info_log-12                1000000              2064 ns/op             513 B/op         16 allocs/op
Benchmark_LogrusLogger/simple_info_log-12                1000000              2075 ns/op             513 B/op         16 allocs/op
Benchmark_LogrusLogger/simple_info_log-12                1000000              2070 ns/op             513 B/op         16 allocs/op
Benchmark_LogrusLogger/log_info_with_5_fields-12          416242              5785 ns/op            3360 B/op         43 allocs/op
Benchmark_LogrusLogger/log_info_with_5_fields-12          423064              5746 ns/op            3360 B/op         43 allocs/op
Benchmark_LogrusLogger/log_info_with_5_fields-12          416894              5683 ns/op            3360 B/op         43 allocs/op
Benchmark_LogrusLogger/log_info_with_5_fields-12          427000              5722 ns/op            3360 B/op         43 allocs/op
Benchmark_LogrusLogger/log_info_with_5_fields-12          422842              5718 ns/op            3359 B/op         43 allocs/op
Benchmark_ZapLogger/simple_info_log-12                  22104961               113.6 ns/op             0 B/op          0 allocs/op
Benchmark_ZapLogger/simple_info_log-12                  21326161               113.4 ns/op             0 B/op          0 allocs/op
Benchmark_ZapLogger/simple_info_log-12                  21195076               115.0 ns/op             0 B/op          0 allocs/op
Benchmark_ZapLogger/simple_info_log-12                  21099259               116.0 ns/op             0 B/op          0 allocs/op
Benchmark_ZapLogger/simple_info_log-12                  20416849               117.2 ns/op             0 B/op          0 allocs/op
Benchmark_ZapLogger/log_info_with_5_fields-12            1319438              1890 ns/op            6733 B/op         30 allocs/op
Benchmark_ZapLogger/log_info_with_5_fields-12            1346614              1750 ns/op            6733 B/op         30 allocs/op
Benchmark_ZapLogger/log_info_with_5_fields-12            1354759              1782 ns/op            6733 B/op         30 allocs/op
Benchmark_ZapLogger/log_info_with_5_fields-12            1316520              1776 ns/op            6733 B/op         30 allocs/op
Benchmark_ZapLogger/log_info_with_5_fields-12            1345928              1763 ns/op            6733 B/op         30 allocs/op
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
