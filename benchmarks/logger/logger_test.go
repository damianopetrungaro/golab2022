package logger

import (
	"io"
	"testing"

	"github.com/damianopetrungaro/workshop"
)

/**
go test --bench=. --benchmem -benchtime=2s -count=5
Benchmark_WorkshopLogger/simple_info_log-12             41561128                50.11 ns/op           32 B/op          2 allocs/op
Benchmark_WorkshopLogger/simple_info_log-12             46491171                49.42 ns/op           32 B/op          2 allocs/op
Benchmark_WorkshopLogger/simple_info_log-12             48416540                50.55 ns/op           32 B/op          2 allocs/op
Benchmark_WorkshopLogger/simple_info_log-12             48467454                51.60 ns/op           32 B/op          2 allocs/op
Benchmark_WorkshopLogger/simple_info_log-12             37740775                62.50 ns/op           32 B/op          2 allocs/op
Benchmark_WorkshopLogger/log_info_with_5_fields-12       4094596               545.5 ns/op           969 B/op         21 allocs/op
Benchmark_WorkshopLogger/log_info_with_5_fields-12       4486573               551.0 ns/op           969 B/op         21 allocs/op
Benchmark_WorkshopLogger/log_info_with_5_fields-12       4312716               556.1 ns/op           969 B/op         21 allocs/op
Benchmark_WorkshopLogger/log_info_with_5_fields-12       4406878               544.0 ns/op           969 B/op         21 allocs/op
Benchmark_WorkshopLogger/log_info_with_5_fields-12       4361113               559.7 ns/op           969 B/op         21 allocs/op

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
