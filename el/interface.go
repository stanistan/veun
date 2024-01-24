package el

type el[T any] interface {
	In(*Element) *Element
	Class(string) T
	Attrs(Attrs) T
	Attr(string, string) T
}
