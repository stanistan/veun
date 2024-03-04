// Package el contains types representing HTML elements.
//
// Each element is represented by a type containing
// combinations to create it.
//
// It allows to write our HTML tree in a very familiar pattern:
//
//	view := el.Div{
//		el.Class("foo", "bar"),
//		el.P{
//			el.Text("hi!"),
//		},
//	})
//
// Most of the contents in this package are generated.
package el

//go:generate ./generate-elements elements.txt elements.generated.go
