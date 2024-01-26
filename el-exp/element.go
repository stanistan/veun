package el

import (
	"context"
	"html/template"

	"github.com/stanistan/veun"
)

// Param represents a parameter to a non-void HTML element.
type Param interface {
	applyToElement(e *element)
}

type element struct {
	tag      string
	attrs    Attrs
	children []veun.AsView
}

func (e element) apply(params []Param) element {
	for _, param := range params {
		param.applyToElement(&e)
	}

	return e
}

func (e element) View(_ context.Context) (*veun.View, error) {
	return veun.V(e), nil
}

//nolint:gosec
func (e element) AsHTML(ctx context.Context) (template.HTML, error) {
	content, err := veun.Render(ctx, veun.Views(e.children))
	if err != nil {
		return content, err
	}

	return template.HTML("<"+e.tag+e.attrs.render()+">") +
		content +
		template.HTML("</"+e.tag+">"), nil
}

// Content is a group of [veun.AsView] it can also be applied to a
// non-void HTML element, such as [Div].
type Content []veun.AsView

func (v Content) applyToElement(e *element) {
	e.children = append(e.children, v...)
}
