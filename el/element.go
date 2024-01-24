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

// View constructs a [*veun.View] from an Element.
func (e *Element) View(ctx context.Context) (*veun.View, error) {
	return veun.Views{
		veun.Raw(openingTag(e.tag, e.attrs)),
		e.content,
		veun.Raw(closingTag(e.tag)),
	}.View(ctx)
}

// Attrs sets the attributes for the element.
func (e *Element) Attrs(a Attrs) *Element {
	e.attrs = a

	return e
}

// Attr sets a single attribute on the element.
func (e *Element) Attr(name, value string) *Element {
	if e.attrs == nil {
		e.attrs = Attrs{}
	}

	e.attrs[name] = value

	return e
}

// Class sets the class attribute.
func (e *Element) Class(name string) *Element {
	return e.Attr("class", name)
}

// Content sets the inner content of the element.
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

// InnerText sets the content to be html escaped text.
func (e *Element) InnerText(t string) *Element {
	return e.Content(Text(t))
}

// In encloses the current element in a parent, returning
// the parent.
func (e *Element) In(parent *Element) *Element {
	return parent.Content(e)
}

// Text creates a HTML escaped text view.
func Text(in string) veun.AsView { //nolint:ireturn
	return text(template.HTMLEscapeString(in))
}

type text string

//nolint:gosec
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
