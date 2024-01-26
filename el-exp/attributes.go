package el

import (
	"html/template"
	"strings"
)

// Attrs represents an attribute map.
//
// Attrs can be applied to both void and non-void elements.
type Attrs map[string]string

func (a Attrs) render() string {
	var w strings.Builder

	for k, v := range a {
		_, _ = w.WriteString(" ")
		template.HTMLEscape(&w, []byte(k))
		_, _ = w.WriteString(`="`)
		template.HTMLEscape(&w, []byte(v))
		_, _ = w.WriteString(`"`)
	}

	return w.String()
}

func (a Attrs) applyToElement(e *element) {
	if e.attrs == nil {
		e.attrs = Attrs{}
	}
	for k, v := range a {
		e.attrs[k] = v
	}
}

func (a Attrs) applyToVoidElement(e *voidElement) {
	if e.attrs == nil {
		e.attrs = Attrs{}
	}
	for k, v := range a {
		e.attrs[k] = v
	}
}

// Attr is a single key value attribute.
//
// Attr can be applied to both void and non-void elements.
type Attr [2]string

func (a Attr) applyToElement(e *element) {
	if e.attrs == nil {
		e.attrs = Attrs{}
	}
	e.attrs[a[0]] = a[1]
}

func (a Attr) applyToVoidElement(e *voidElement) {
	if e.attrs == nil {
		e.attrs = Attrs{}
	}
	e.attrs[a[0]] = a[1]
}
