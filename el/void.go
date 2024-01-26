package el

import (
	"context"

	"github.com/stanistan/veun"
)

// VoidElement is a representation of an empty/void HTML Element
// that is also a veun.View.
type VoidElement struct {
	tag   string
	attrs Attrs
}

var (
	_ veun.AsView      = &VoidElement{}
	_ el[*VoidElement] = &VoidElement{}
)

// View constructs a [*veun.View] from a VoidElement.
func (e *VoidElement) View(ctx context.Context) (*veun.View, error) {
	return veun.Raw(openingTag(e.tag, e.attrs)).View(ctx)
}

// Attrs sets the attributes for the element.
func (e *VoidElement) Attrs(a Attrs) *VoidElement {
	e.attrs = a

	return e
}

// Attr sets a single attribute on the element.
func (e *VoidElement) Attr(name, value string) *VoidElement {
	if e.attrs == nil {
		e.attrs = Attrs{}
	}

	e.attrs[name] = value

	return e
}

// In encloses the current element in a parent, returning
// the parent.
func (e *VoidElement) In(parent *Element) *Element {
	return parent.Content(e)
}

// Class sets the class attribute.
func (e *VoidElement) Class(name string) *VoidElement {
	return e.Attr("class", name)
}
