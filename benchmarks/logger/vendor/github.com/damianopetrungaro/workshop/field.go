package workshop

import (
	"strconv"
)

type (
	Fields []Field
	Field  struct {
		k string
		v any
	}
)

func (f Field) MarshalJSON() ([]byte, error) {
	s := `{"` + f.k + `":`
	switch t := f.v.(type) {
	case string:
		s = s + `"` + t + `"`
	case int:
		s = s + strconv.Itoa(t)
	default:
		s = s + `"--error--"`
	}
	s = s + `}`
	return []byte(s), nil
}

func String(k, v string) Field {
	return Field{k: k, v: v}
}

func Int(k string, v int) Field {
	return Field{k: k, v: v}
}
