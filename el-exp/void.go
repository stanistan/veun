package el

import (
	"context"
	"html/template"

	"github.com/stanistan/veun"
)

// VoidParam represents a parameter to void html element,
// such as [Br], [Hr], etc.
type VoidParam interface {
	applyToVoidElement(e *voidElement)
}

type voidElement struct {
	tag   string
	attrs Attrs
}

func (e voidElement) apply(params []VoidParam) voidElement {
	for _, param := range params {
		param.applyToVoidElement(&e)
	}

	return e
}

func (e voidElement) View(_ context.Context) (*veun.View, error) {
	return veun.V(e), nil
}

//nolint:gosec
func (e voidElement) AsHTML(_ context.Context) (template.HTML, error) {
	return template.HTML("<" + e.tag + e.attrs.render() + ">"), nil
}
