package veun

import (
	"context"
	"html/template"
)

type ErrorHandler interface {
	ViewForError(ctx context.Context, err error) (AsR, error)
}

type ErrorHandlerFunc func(context.Context, error) (AsR, error)

func (f ErrorHandlerFunc) ViewForError(ctx context.Context, err error) (AsR, error) {
	return f(ctx, err)
}

func PassThroughErrorHandler() ErrorHandler {
	return ErrorHandlerFunc(func(_ context.Context, err error) (AsR, error) {
		return nil, err
	})
}

func MakeErrorHandler(from any) ErrorHandler {
	if eh, ok := from.(ErrorHandler); ok && eh != nil {
		return eh
	}

	return PassThroughErrorHandler()
}

func handleRenderError(ctx context.Context, err error, with any) (template.HTML, error) {
	var empty template.HTML

	v, err := MakeErrorHandler(with).ViewForError(ctx, err)
	if err != nil {
		return empty, err
	}

	return Render(ctx, v)
}
