package el

import (
	"context"
	"html/template"

	"github.com/stanistan/veun/internal/view"
)

// Text represents any text that should be HTML escaped.
type Text string

func (t Text) applyToElement(e *element[nodeChildren]) {
	e.inner = append(e.inner, t)
}

// AsHTML implements [view.HTMLRenderable] for [Text].
func (t Text) AsHTML(_ context.Context) (template.HTML, error) {
	return template.HTML(template.HTMLEscapeString(string(t))), nil //nolint:gosec
}

// View implements [view.AsView] for [Text].
func (t Text) View(_ context.Context) (*view.View, error) {
	return view.V(t), nil
}
