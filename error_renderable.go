package veun

import (
	"html/template"
)

type ErrorHandler interface {
	ViewForError(e *Error) (Renderable, error)
}

type ErrorHandlerFunc func(*Error) (Renderable, error)

func (f ErrorHandlerFunc) ViewForError(e *Error) (Renderable, error) {
	return f(e)
}

func PassThroughErrorHandler() ErrorHandler {
	return ErrorHandlerFunc(func(e *Error) (Renderable, error) {
		return nil, e.Err
	})
}

func MakeErrorHandler(in any) ErrorHandler {
	errR, ok := in.(ErrorHandler)
	if !ok || in == nil {
		return PassThroughErrorHandler()
	} else {
		return errR
	}
}

func RenderError(e *Error, with any) (template.HTML, error) {
	v, err := MakeErrorHandler(with).ViewForError(e)
	if err != nil {
		return emptyHTML(), err
	}

	out, err := RenderToHTML(e.Context(), v, nil)
	if err != nil {
		return emptyHTML(), err
	}

	return out, nil
}
