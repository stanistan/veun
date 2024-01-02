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

type AsView interface {
	View(ctx context.Context) (*View, error)
}

type errViewInvalid struct {
	Err error
}

func (e errViewInvalid) Error() string {
	return e.Err.Error()
}

func (e errViewInvalid) AsHTML(_ context.Context) (template.HTML, error) {
	return template.HTML(""), e
}

type renderable struct {
	r AsView
}

func (r renderable) AsHTML(ctx context.Context) (template.HTML, error) {
	var empty template.HTML

	v, err := r.r.View(ctx)
	if err != nil {
		return empty, fmt.Errorf("%T Renderable: %w", r.r, err)
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
	case AsView:
		return renderable{t}
	default:
		return errViewInvalid{
			Err: fmt.Errorf("invalid view %T", in),
		}
	}
}

func (r *View) View(ctx context.Context) (*View, error) {
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

func V(in any) *View {
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
	case AsView:
		return &View{
			r:  renderable{t},
			eh: PassThroughErrorHandler(),
		}
	}

	return &View{
		eh: PassThroughErrorHandler(),
		r:  errViewInvalid{fmt.Errorf("can't construct %T", in)},
	}
}

func (r *View) WithErrorHandler(eh ErrorHandler) *View {
	return &View{r: r.r, eh: eh}
}
