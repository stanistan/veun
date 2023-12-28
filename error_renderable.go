package veun

import (
	"context"
	"html/template"
)

type ErrorHandler interface {
	ViewForError(ctx context.Context, err error) (AsRenderable, error)
}

type ErrorHandlerFunc func(context.Context, error) (AsRenderable, error)

func (f ErrorHandlerFunc) ViewForError(ctx context.Context, err error) (AsRenderable, error) {
	return f(ctx, err)
}

func handleRenderError(ctx context.Context, err error, with any) (template.HTML, error) {
	var empty template.HTML

	if with == nil {
		return empty, err
	}

	errRenderable, ok := with.(ErrorHandler)
	if !ok {
		return empty, err
	}

	r, err := errRenderable.ViewForError(ctx, err)
	if err != nil {
		return empty, err
	}

	if r == nil {
		return empty, nil
	}

	return Render(ctx, r)
}
