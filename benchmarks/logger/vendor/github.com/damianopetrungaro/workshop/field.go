package workshop

import (
	"bytes"
	"strconv"
)

type (
	Fields []Field
	Field  struct {
		k string
		v any
	}
)

func (fs Fields) Append(buf *bytes.Buffer) {
	buf.WriteString(`[`)
	for i, f := range fs {
		i++
		buf.WriteString(`{"`)
		buf.WriteString(f.k)
		buf.WriteString(`":`)
		switch t := f.v.(type) {
		case string:
			buf.WriteString(`"`)
			buf.WriteString(t)
			buf.WriteString(`"`)
		case int:
			buf.WriteString(strconv.Itoa(t))
		default:
			buf.WriteString(`"--error--"`)
		}
		buf.WriteString(`}`)
		if i != len(fs) {
			buf.WriteString(`,`)
		}
	}
	buf.WriteString(`]`)
}

func String(k, v string) Field {
	return Field{k: k, v: v}
}

func Int(k string, v int) Field {
	return Field{k: k, v: v}
}
