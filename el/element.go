package el

import (
	"context"
	"html/template"

	"github.com/stanistan/veun/internal/view"
)

type element[T elementKind] struct {
	tag   tag
	inner T
}

func (e element[T]) View(_ context.Context) (*view.View, error) {
	return view.V(e), nil
}

//nolint:wrapcheck
func (e element[T]) AsHTML(ctx context.Context) (template.HTML, error) {
	return e.inner.AsHTML(ctx, e.tag)
}

func (e *element[T]) attrs(fn func(Attrs)) {
	e.tag.applyAttrs(fn)
}

func newElementWithChildren(t string, ps []Param) element[nodeChildren] {
	e := element[nodeChildren]{tag: tag{name: t}, inner: nodeChildren{}}
	for _, p := range ps {
		p.applyToElement(&e)
	}

	return e
}

func newVoidElement(t string, ps []VoidParam) element[void] {
	e := element[void]{tag: tag{name: t}, inner: void{}}
	for _, p := range ps {
		p.applyToVoidElement(&e)
	}

	return e
}

func MapFragment[T any, U Param, E ~[]T](
	in E, fn func(T, int) U,
) Fragment {
	var out Fragment
	for idx, v := range in {
		out = append(out, fn(v, idx))
	}

	return out
}
