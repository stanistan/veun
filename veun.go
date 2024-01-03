package veun

import (
	"context"
	"fmt"
	"html/template"
)

// AsV is an alias for the AsView interface.
type AsV = AsView

// HTMLRenderable represents anything that can be rendered to HTML.
type HTMLRenderable interface {
	AsHTML(ctx context.Context) (template.HTML, error)
}

// AsView is anything that can be represented as a View.
type AsView interface {
	View(ctx context.Context) (*View, error)
}

// Render takes a context and something that can become a View
// and renders it.
func Render(ctx context.Context, v AsView) (template.HTML, error) {
	return (&View{
		r:  renderable{v},
		eh: PassThroughErrorHandler(),
	}).render(ctx)
}

// V is a factory function that transforms any of its
// inputs into a [[View]].
//
// If this is not view convertible, this call will succeed,
// but any call go [[Render]] this will fail. This is by
// design to allow for error handling during composition.
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
