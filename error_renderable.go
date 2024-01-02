package veun

import (
	"context"
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
