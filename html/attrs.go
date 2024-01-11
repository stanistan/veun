package html

import (
	"html/template"
	"io"
)

var (
	space       = []byte(" ")
	doubleQuote = []byte(`"`)
	equalsQuote = []byte(`="`)
)

type Attrs map[string]string

func (a Attrs) writeTo(w io.Writer) {
	for k, v := range a {
		_, _ = w.Write(space)
		template.HTMLEscape(w, []byte(k))
		_, _ = w.Write(equalsQuote)
		template.HTMLEscape(w, []byte(v))
		_, _ = w.Write(doubleQuote)
	}
}
