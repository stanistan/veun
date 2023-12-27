package veun

import (
	"html/template"
)

type ErrorRenderable interface {
	ErrorRenderable(e *Error) (AsRenderable, error)
}

type ErrorRenderableFunc func(*Error) (AsRenderable, error)

func (f ErrorRenderableFunc) ErrorRenderable(e *Error) (AsRenderable, error) {
	return f(e)
}

func PassThroughErrorRenderable() ErrorRenderable {
	return ErrorRenderableFunc(func(e *Error) (AsRenderable, error) {
		return nil, e.Err
	})
}

func MakeErrorRenderable(in any) ErrorRenderable {
	errR, ok := in.(ErrorRenderable)
	if !ok || in == nil {
		return PassThroughErrorRenderable()
	} else {
		return errR
	}
}

func RenderError(e *Error, with any) (template.HTML, error) {
	v, err := MakeErrorRenderable(with).ErrorRenderable(e)
	if err != nil {
		return emptyHTML(), err
	}

	out, err := Render(e.Context(), v)
	if err != nil {
		return emptyHTML(), err
	}

	return out, nil
}
