package logger

import (
	"io"
	"testing"

	"github.com/damianopetrungaro/workshop"
)

/**
go test --bench=. --benchmem -benchtime=2s -count=5
Benchmark_WorkshopLogger/simple_info_log-12     69667310                28.95 ns/op           16 B/op          1 allocs/op
Benchmark_WorkshopLogger/simple_info_log-12     79329039                29.85 ns/op           16 B/op          1 allocs/op
Benchmark_WorkshopLogger/simple_info_log-12     76109995                30.87 ns/op           16 B/op          1 allocs/op
Benchmark_WorkshopLogger/simple_info_log-12     68782282                34.16 ns/op           16 B/op          1 allocs/op
Benchmark_WorkshopLogger/simple_info_log-12     67763760                35.02 ns/op           16 B/op          1 allocs/op
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
}
