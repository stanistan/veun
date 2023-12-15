package veun

import (
	"context"
	"html/template"
)

// Renderable represents anything that can be rendered
// to HTML.
type Renderable interface {
	RenderToHTML(ctx context.Context) (template.HTML, error)
}

type AsRenderable interface {
	// Renderable produces a Renderable struct given a context.
	Renderable(ctx context.Context) (Renderable, error)
}

// RenderableFunc is a function that conforms to the Renderable interface.
type RenderableFunc func(context.Context) (Renderable, error)

// Renderable implements Renderable for RenderableFunc.
func (f RenderableFunc) Renderable(ctx context.Context) (Renderable, error) {
	return f(ctx)
}
