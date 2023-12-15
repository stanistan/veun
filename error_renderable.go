package veun

import (
	"context"
	"html/template"
)

type ErrorRenderable interface {
	ErrorRenderable(ctx context.Context, err error) (AsRenderable, error)
}

type ErrorRenderableFunc func(context.Context, error) (AsRenderable, error)

func (f ErrorRenderableFunc) ErrorRenderable(ctx context.Context, err error) (AsRenderable, error) {
	return f(ctx, err)
}

func handleRenderError(ctx context.Context, err error, with any) (template.HTML, error) {
	var empty template.HTML

	if with == nil {
		return empty, err
	}

	errRenderable, ok := with.(ErrorRenderable)
	if !ok {
		return empty, err
	}

	r, err := errRenderable.ErrorRenderable(ctx, err)
	if err != nil {
		return empty, err
	}

	if r == nil {
		return empty, nil
	}

	return Render(ctx, r)
}
