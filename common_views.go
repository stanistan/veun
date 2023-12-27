package veun

import (
	"context"
	"html/template"
)

// Raw represents a string view.
type Raw string

// Renderable implements AsRenderable for Raw.
func (r Raw) Renderable(_ context.Context) (Renderable, error) { return r, nil }

// RenderToHTML implements Renderable for Raw.
func (r Raw) RenderToHTML(_ context.Context) (template.HTML, error) { return template.HTML(r), nil }

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

type RViews []Renderable

func (vs RViews) Renderable(_ context.Context) (Renderable, error) {
	return vs, nil
}

func (vs RViews) RenderToHTML(ctx context.Context) (template.HTML, error) {
	var out template.HTML

	for _, r := range vs {
		html, err := r.RenderToHTML(ctx)
		if err != nil {
			return template.HTML(""), err
		}

		out += html
	}

	return out, nil
}
