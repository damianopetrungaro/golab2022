package fields

import (
	"fmt"
	"testing"
)

func BenchmarkWith(b *testing.B) {
	b.Run("slice", func(b *testing.B) {
		slice := NewSlice()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			k := fmt.Sprintf("key_%d", i)
			slice = slice.With(k, "value")
		}
	})
	b.Run("map", func(b *testing.B) {
		_map := NewMap()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			k := fmt.Sprintf("key_%d", i)
			_map = _map.With(k, "value")
		}
	})
}
