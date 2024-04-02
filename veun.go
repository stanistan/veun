// Package veun is a small library to enable composition based
// template rendering of go functions and types into HTML.
package veun

import (
	"context"
	"html/template"

	"github.com/stanistan/veun/internal/view"
)

type (
	View   = view.View
	AsView = view.AsView
	AsV    = view.AsView
)

type (
	Raw   = view.Raw
	Views = view.Views
)

// Render renders a view tree into HTML given a context.
//
//nolint:wrapcheck
func Render(ctx context.Context, v AsView) (template.HTML, error) {
	return view.Render(ctx, v)
}

// V is a factory function that transforms any of its
// inputs into a [View].
//
// If this is not view convertible, this call will succeed,
// but any call to [Render] will fail.
//
// This is by design to allow for error handling during composition.
func V(in any) *View {
	return view.V(in)
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
