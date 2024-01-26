package el

import (
	"context"
	"html/template"

	"github.com/stanistan/veun"
)

// Text represents any text that should be HTML escaped.
type Text string

func (t Text) applyToElement(e *element) {
	e.children = append(e.children, t)
}

// AsHTML implements [veun.HTMLRenderable] for [Text].
func (t Text) AsHTML(_ context.Context) (template.HTML, error) {
	return template.HTML(template.HTMLEscapeString(string(t))), nil //nolint:gosec
}

// View implements [veun.AsView] for [Text].
func (t Text) View(_ context.Context) (*veun.View, error) {
	return veun.V(t), nil
}
