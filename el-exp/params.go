package el

// Param represents a parameter to a non-void HTML element,
// it can be a mutation to the attributes, or children.
//
// See: [P], [Div], etc.
type Param interface {
	applyToElement(e *element[nodeChildren])
}

// VoidParam represents a parameter to a void html element,
// such as [Br], [Hr], etc.
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
