package veun

import (
	"context"
	"html/template"
)

type ErrorRenderable interface {
	// ErrorRenderable can return bubble the error
	// back up, which will continue to fail the render
	// the same as it did before.
	//
	// It can also return nil for Renderable,
	// which will ignore the error entirely.
	//
	// Otherwise we will attempt to render next one.
	ErrorRenderable(ctx context.Context, err error) (AsRenderable, error)
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
