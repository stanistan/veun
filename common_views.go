package veun

import (
	"context"
	"html/template"
)

type Raw string

func (r Raw) Renderable(_ context.Context) (Renderable, error) { return r, nil }

func (r Raw) RenderToHTML(_ context.Context) (template.HTML, error) {
	return template.HTML(r), nil
}

type Views []AsRenderable

func (vs Views) Renderable(ctx context.Context) (Renderable, error) {
	return vs, nil
}

func (vs Views) RenderToHTML(ctx context.Context) (template.HTML, error) {
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
