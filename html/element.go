package html

import (
	"context"
	"html/template"

	"github.com/stanistan/veun"
)

//go:generate ./generate-elements

func El(name string, attrs Attrs, content veun.AsView) veun.AsView {
	return veun.Views{
		veun.Raw(openingTag(name, attrs)),
		content,
		veun.Raw(closingTag(name)),
	}
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
