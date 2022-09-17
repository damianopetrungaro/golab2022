package fields

import (
	"sync"
)

type field struct {
	k string
	v any
}

type Slice struct {
	data []field
}

func NewSlice() Slice {
	return Slice{}
}

func (s Slice) With(k string, v any) Slice {
	return Slice{data: append(s.data, field{k: k, v: v})}
}

type Map struct {
	data map[string]any
	mu   sync.Mutex
}

func NewMap() *Map {
	return &Map{data: map[string]any{}}
}

func (m *Map) With(k string, v any) *Map {
	m.mu.Lock()
	defer m.mu.Unlock()

	data := make(map[string]any, len(m.data)+1)
	for k, v := range m.data {
		data[k] = v
	}

	data[k] = v
	return &Map{data: data}
}
