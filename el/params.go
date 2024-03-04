package el

import "github.com/stanistan/veun"

// Param represents a parameter to a non-void HTML element,
// it can be a mutation to the attributes, or children.
//
// Every defined element is also a [Param], meaning they
// can be added as a child to another element.
//
// See: [P], [Div], etc.
type Param interface {
	applyToElement(e *element[nodeChildren])
}

// VoidParam represents a parameter to a void html element,
// such as [Br], [Hr], etc.
//
// In practice what this means this is they are mutations
// on the attributes of the element.
//
// See [Attributes] and [WithAttr].
//
// All types that implement [VoidParam] also implement [Param].
type VoidParam interface {
	applyToVoidElement(e *element[void])
}

// Fragment holds parameters to transform a non-void HTML element.
//
// This is similar to an element type such as [Div], but
// applies the parameters to its parent.
type Fragment []Param

func (f Fragment) applyToElement(e *element[nodeChildren]) {
	for _, p := range f {
		p.applyToElement(e)
	}
}

// VoidFragment holds parameters to transform a void HTML element.
//
// This is similar to an element type such as [Br], but
// applies the parameters to its parent.
type VoidFragment []VoidParam

func (v VoidFragment) applyToVoidElement(e *element[void]) {
	for _, p := range v {
		p.applyToVoidElement(e)
	}
}

// Content is a group of [veun.AsView] it can also be applied to a
// non-void HTML element, such as [Div].
type Content []veun.AsView

func (v Content) applyToElement(e *element[nodeChildren]) {
	e.inner = append(e.inner, v...)
}
