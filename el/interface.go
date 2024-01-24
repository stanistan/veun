package el

type el[T any] interface {
	In(parent *Element) *Element
	Class(name string) T
	Attrs(a Attrs) T
	Attr(name string, value string) T
}
