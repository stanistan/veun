package veun

import (
	"context"
	"html/template"
)

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

// AsViews transforms a slice of T (implementing AsView) into Views.
func AsViews[T AsView](ts []T) Views {
	vs := make(Views, len(ts))
	for idx, v := range ts {
		vs[idx] = v
	}

	return vs
}

// MapToViews transforms a slice of T (any) with a function into Views.
func MapToViews[T any, V AsView](ts []T, f func(T) V) Views {
	vs := make(Views, len(ts))
	for idx, v := range ts {
		vs[idx] = f(v)
	}

	return vs
}
