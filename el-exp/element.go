package el

import (
	"context"
	"html/template"

	"github.com/stanistan/veun"
)

type element[T elementKind] struct {
	tag   tag
	inner T
}

func (e element[T]) View(_ context.Context) (*veun.View, error) {
	return veun.V(e), nil
}

func (e element[T]) AsHTML(ctx context.Context) (template.HTML, error) {
	return e.inner.AsHTML(ctx, e.tag)
}

func (e *element[T]) attrs(fn func(Attrs)) {
	e.tag.applyAttrs(fn)
}

func newElement[T elementKind](t string) element[T] {
	return element[T]{tag: tag{name: t}}
}

func newElementWithChildren(t string, ps []Param) element[nodeChildren] {
	e := newElement[nodeChildren](t)
	for _, p := range ps {
		p.applyToElement(&e)
	}

	return e
}

func newVoidElement(t string, ps []VoidParam) element[void] {
	e := newElement[void](t)
	for _, p := range ps {
		p.applyToVoidElement(&e)
	}

	return e
}

// Content is a group of [veun.AsView] it can also be applied to a
// non-void HTML element, such as [Div].
type Content []veun.AsView

func (v Content) applyToElement(e *element[nodeChildren]) {
	e.inner = append(e.inner, v...)
}
