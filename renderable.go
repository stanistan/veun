package veun

import (
	"context"
	"html/template"
)

// Renderable represents any struct that can be rendered
// in the Render function.
type Renderable interface {
	// Template provides the template object / parsed and compiled,
	// that Render will execute given a context.
	Template(ctx context.Context) (*template.Template, error)
	// TemplateData provides the data to the template given a context.
	TemplateData(ctx context.Context) (any, error)
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
