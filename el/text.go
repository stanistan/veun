package el

import (
	"context"
	"html/template"

	"github.com/stanistan/veun"
)

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
