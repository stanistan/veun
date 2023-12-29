package veun

import (
	"context"
	"html/template"
)

type Raw string

func (r Raw) Renderable(_ context.Context) (*View, error) { return R(r), nil }

func (r Raw) AsHTML(_ context.Context) (template.HTML, error) {
	return template.HTML(r), nil
}

type Views []AsR

func (vs Views) Renderable(_ context.Context) (*View, error) {
	return R(vs), nil
}

func (vs Views) AsHTML(ctx context.Context) (template.HTML, error) {
	var out template.HTML
	for _, v := range vs {
		html, err := Render(ctx, v)
		if err != nil {
			return template.HTML(""), err
		}

		out += html
	}

	return out, nil
}
