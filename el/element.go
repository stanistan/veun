package el

import (
	"context"
	"html/template"
	"strings"

	"github.com/stanistan/veun"
)

//go:generate ./generate-elements elements.txt

// Element is a representation of an HTML Element that is also a veun.View.
type Element struct {
	tag     string
	attrs   Attrs
	content veun.AsView
}

var _ veun.AsView = &Element{}

func (e *Element) View(ctx context.Context) (*veun.View, error) {
	return veun.Views{
		veun.Raw(openingTag(e.tag, e.attrs)),
		e.content,
		veun.Raw(closingTag(e.tag)),
	}.View(ctx)
}

func (e *Element) Attrs(a Attrs) *Element {
	e.attrs = a
	return e
}

func (e *Element) Attr(name, value string) *Element {
	if e.attrs == nil {
		e.attrs = Attrs{}
	}
	e.attrs[name] = value
	return e
}

func (e *Element) Class(name string) *Element {
	return e.Attr("class", name)
}

func (e *Element) Content(cs ...veun.AsView) *Element {
	switch len(cs) {
	case 0:
		e.content = nil
	case 1:
		e.content = cs[0]
	default:
		e.content = veun.Views(cs)
	}

	return e
}

func (e *Element) InnerText(t string) *Element {
	return e.Content(Text(t))
}

// El is a constructor for an Element.
func El(name string, short bool) *Element {
	return &Element{tag: name}
}

func Text(in string) veun.AsView {
	return text(template.HTMLEscapeString(in))
}

type text string

func (t text) AsHTML(_ context.Context) (template.HTML, error) {
	return template.HTML(t), nil
}

func (t text) View(_ context.Context) (*veun.View, error) {
	return veun.V(t), nil
}

func openingTag(name string, a Attrs) string {
	var sb strings.Builder

	sb.WriteString("<")
	sb.WriteString(name)
	a.writeTo(&sb)
	sb.WriteString(">")

	return sb.String()
}

func closingTag(name string) string {
	return "</" + name + ">"
}
