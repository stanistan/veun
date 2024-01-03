package veun

import (
	"context"
	"fmt"
	"html/template"
)

type ErrorHandler interface {
	ViewForError(ctx context.Context, err error) (AsView, error)
}

type ErrorHandlerFunc func(context.Context, error) (AsView, error)

func (f ErrorHandlerFunc) ViewForError(ctx context.Context, err error) (AsView, error) {
	return f(ctx, err)
}

func PassThroughErrorHandler() ErrorHandler {
	return ErrorHandlerFunc(func(_ context.Context, err error) (AsView, error) {
		return nil, err
	})
}

func MakeErrorHandler(from any) ErrorHandler {
	if eh, ok := from.(ErrorHandler); ok && eh != nil {
		return eh
	}

	return PassThroughErrorHandler()
}

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
		return empty, fmt.Errorf("renderError %T: %w", v, err)
	}

	return out, nil
}
