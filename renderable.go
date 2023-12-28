package veun

import (
	"context"
	"html/template"
)

// HTMLRenderable represents anything that can be rendered to HTML.
type HTMLRenderable interface {
	AsHTML(ctx context.Context) (template.HTML, error)
}

type AsRenderable interface {
	// Renderable produces a Renderable struct given a context.
	Renderable(ctx context.Context) (HTMLRenderable, error)
}

// RenderableFunc is a function that conforms to the Renderable interface.
type RenderableFunc func(context.Context) (HTMLRenderable, error)

// Renderable implements Renderable for RenderableFunc.
func (f RenderableFunc) Renderable(ctx context.Context) (HTMLRenderable, error) {
	return f(ctx)
}
