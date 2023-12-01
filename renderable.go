package veun

import (
	"context"
	"html/template"
)

type Renderable interface {
	Template(ctx context.Context) (*template.Template, error)
	TemplateData(ctx context.Context) (any, error)
}

type AsRenderable interface {
	Renderable(ctx context.Context) (Renderable, error)
}

type RenderableFunc func(context.Context) (Renderable, error)

func (f RenderableFunc) Renderable(ctx context.Context) (Renderable, error) {
	return f(ctx)
}
