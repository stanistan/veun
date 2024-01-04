package veun

import (
	"context"
	"fmt"
	"html/template"
)

// ErrorHandler represents something that can handle an error.
type ErrorHandler interface {
	ViewForError(ctx context.Context, err error) (AsView, error)
}

// ErrorHandlerFunc is a function representation of ErrorHandler.
type ErrorHandlerFunc func(context.Context, error) (AsView, error)

// ViewForError implements ErrorHandler.
func (f ErrorHandlerFunc) ViewForError(ctx context.Context, err error) (AsView, error) {
	return f(ctx, err)
}

// RenderError renders an error given an ErrorHandler.
func RenderError(ctx context.Context, h ErrorHandler, err error) (template.HTML, error) {
	var empty template.HTML

	if h == nil {
		return empty, err
	}

	v, err := h.ViewForError(ctx, err)
	if err != nil {
		return empty, err
	}

	if v == nil {
		return empty, nil
	}

	out, err := Render(ctx, v)
	if err != nil {
		return empty, fmt.Errorf("RenderError %T: %w", v, err)
	}

	return out, nil
}
