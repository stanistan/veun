package veun

import (
	"context"
	"fmt"
	"html/template"
)

// HTMLRenderable represents anything that can be rendered to HTML.
type HTMLRenderable interface {
	AsHTML(ctx context.Context) (template.HTML, error)
}

type AsR interface {
	Renderable(ctx context.Context) (*Renderable, error)
}

type errorHTMLRenderable struct {
	err error
}

func (e errorHTMLRenderable) AsHTML(_ context.Context) (template.HTML, error) {
	return template.HTML(""), e.err
}

type asHTML struct {
	r AsR
}

func (r asHTML) AsHTML(ctx context.Context) (template.HTML, error) {
	var empty template.HTML

	v, err := r.r.Renderable(ctx)
	if err != nil {
		return empty, err
	}

	return v.Render(ctx)
}

type Renderable struct {
	r  HTMLRenderable
	eh ErrorHandler
}

func makeHTMLRenderable(in any) HTMLRenderable {
	if in == nil {
		return Raw("")
	}

	switch t := in.(type) {
	case template.HTML:
		return Raw(t)
	case HTMLRenderable:
		return t
	case AsR:
		return asHTML{t}
	default:
		return errorHTMLRenderable{
			err: fmt.Errorf("invalid renderable %T", in),
		}
	}
}

func (r *Renderable) Renderable(ctx context.Context) (*Renderable, error) {
	return r, nil
}

func (r *Renderable) Render(ctx context.Context) (template.HTML, error) {
	var empty template.HTML

	if r == nil {
		return empty, nil
	}

	if r.r == nil {
		return empty, nil
	}

	out, err := r.r.AsHTML(ctx)
	if err == nil {
		return out, nil
	}

	// error handling
	v, err := r.eh.ViewForError(ctx, err)
	if err != nil {
		return empty, err
	} else if v == nil {
		return empty, nil
	} else {
		return Render(ctx, v)
	}
}

func R(in any) *Renderable {

	if in == nil {
		return nil
	}

	switch t := in.(type) {
	case *Renderable:
		return t
	case template.HTML:
		return &Renderable{
			r:  Raw(t),
			eh: PassThroughErrorHandler(),
		}
	case HTMLRenderable:
		return &Renderable{
			r:  t,
			eh: PassThroughErrorHandler(),
		}
	case AsR:
		return &Renderable{
			r:  asHTML{t},
			eh: PassThroughErrorHandler(),
		}
	}

	return &Renderable{
		eh: PassThroughErrorHandler(),
		r:  errorHTMLRenderable{fmt.Errorf("can't construct %T", in)},
	}
}

func (r *Renderable) WithErrorHandler(eh ErrorHandler) *Renderable {
	return &Renderable{r: r.r, eh: eh}
}
