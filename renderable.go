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
	Renderable(ctx context.Context) (*View, error)
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
		return empty, fmt.Errorf("Renderable(): %w", err)
	}

	out, err := v.Render(ctx)
	if err != nil {
		return out, fmt.Errorf("Render: %w", err)
	}

	return out, nil
}

type View struct {
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

func (r *View) Renderable(ctx context.Context) (*View, error) {
	return r, nil
}

func (r *View) Render(ctx context.Context) (template.HTML, error) {
	var empty template.HTML

	if r == nil {
		return empty, nil
	}

	if r.r == nil {
		return empty, nil
	}

	out, err := r.r.AsHTML(ctx)
	if err != nil {
		return renderError(ctx, r.eh, err)
	}

	return out, nil
}

func R(in any) *View {
	if in == nil {
		return nil
	}

	switch t := in.(type) {
	case *View:
		return t
	case template.HTML:
		return &View{
			r:  Raw(t),
			eh: PassThroughErrorHandler(),
		}
	case HTMLRenderable:
		return &View{
			r:  t,
			eh: PassThroughErrorHandler(),
		}
	case AsR:
		return &View{
			r:  asHTML{t},
			eh: PassThroughErrorHandler(),
		}
	}

	return &View{
		eh: PassThroughErrorHandler(),
		r:  errorHTMLRenderable{fmt.Errorf("can't construct %T", in)},
	}
}

func (r *View) WithErrorHandler(eh ErrorHandler) *View {
	return &View{r: r.r, eh: eh}
}
