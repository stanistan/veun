package view

import (
	"context"
	"errors"
	"fmt"
	"html/template"
)

// ErrorHandler represents something that can handle an error.
type ErrorHandler interface {
	ViewForError(ctx context.Context, err error) (AsView, error)
}

type ErrorHandlerFunc func(context.Context, error) (AsView, error)

func (f ErrorHandlerFunc) ViewForError(ctx context.Context, err error) (AsView, error) {
	return f(ctx, err)
}

// HTMLRenderable represents anything that can be rendered to [template.HTML].
type HTMLRenderable interface {
	AsHTML(ctx context.Context) (template.HTML, error)
}

// AsView is anything that can be represented as a [*View].
type AsView interface {
	View(ctx context.Context) (*View, error)
}

// View composes an HTMLRenderable with an ErrorHandler.
type View struct {
	r  HTMLRenderable
	eh ErrorHandler
}

// View conforms View to AsView.
func (r *View) View(_ context.Context) (*View, error) {
	return r, nil
}

// WithErrorHandler creates a new View with the error handler.
func (r *View) WithErrorHandler(eh ErrorHandler) *View {
	return &View{r: r.r, eh: eh}
}

func (r *View) render(ctx context.Context) (template.HTML, error) {
	var empty template.HTML

	if r == nil {
		return empty, nil
	}

	if r.r == nil {
		return empty, nil
	}

	out, err := r.r.AsHTML(ctx)
	if err != nil {
		return RenderError(ctx, r.eh, err)
	}

	return out, nil
}

type viewInvalidError struct {
	Err error
}

func (e viewInvalidError) Error() string {
	return e.Err.Error()
}

func (e viewInvalidError) AsHTML(_ context.Context) (template.HTML, error) {
	return template.HTML(""), e
}

type renderable struct {
	r AsView
}

func (r renderable) AsHTML(ctx context.Context) (template.HTML, error) {
	var empty template.HTML

	if r.r == nil {
		return empty, nil
	}

	v, err := r.r.View(ctx)
	if err != nil {
		return empty, fmt.Errorf("View(%T): %w", r.r, err)
	}

	out, err := v.render(ctx)
	if err != nil {
		return out, fmt.Errorf("Render(%T): %w", r.r, err)
	}

	return out, nil
}

// V is a factory function that transforms any of its
// inputs into a [View].
//
// If this is not view convertible, this call will succeed,
// but any call to [Render] will fail.
//
// This is by design to allow for error handling during composition.
func V(in any) *View {
	if in == nil {
		return nil
	}

	switch t := in.(type) {
	case *View:
		return t
	case template.HTML:
		return &View{r: Raw(t)}
	case HTMLRenderable:
		return &View{r: t}
	case AsView:
		return &View{r: renderable{t}}
	}

	return &View{
		r: viewInvalidError{
			Err: fmt.Errorf("invalid input %T: %w", in, errInvalidVParam),
		},
	}
}

var errInvalidVParam = errors.New("can't consturct View")

// RenderError renders an error given an ErrorHandler.
func RenderError(ctx context.Context, h ErrorHandler, err error) (template.HTML, error) {
	var empty template.HTML

	if h == nil {
		return empty, err
	}

	v, err := h.ViewForError(ctx, err)
	if err != nil {
		return empty, fmt.Errorf("RenderError %T: %w", h, err)
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

// Raw represents a string that is directly renderable to HTML.
type Raw string

// View creates a *View.
func (r Raw) View(_ context.Context) (*View, error) { return V(r), nil }

// AsHTML transforms a Raw into template.HTML.
//
//nolint:gosec
func (r Raw) AsHTML(_ context.Context) (template.HTML, error) {
	return template.HTML(r), nil
}

// Views is a collection of AsView.
type Views []AsView

// View creates a *View, which is also an AsView.
func (vs Views) View(_ context.Context) (*View, error) {
	return V(vs), nil
}

// AsHTML renders each view into html concatenating the views.
func (vs Views) AsHTML(ctx context.Context) (template.HTML, error) {
	var out template.HTML

	for _, v := range vs {
		html, err := Render(ctx, v)
		if err != nil {
			return template.HTML(""), err
		}

		out += html
	}

	return out, nil
}

// Render renders a view tree into HTML given a context.
func Render(ctx context.Context, v AsView) (template.HTML, error) {
	return V(v).render(ctx)
}
