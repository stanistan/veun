// Package el contains types representing HTML elements.
//
// This is a separate implementation than github.com/stanistan/veun/el where that
// package is based on function composition and this one is based on struct
// composition.
//
// Most of the contents in this package are generated.
package el

//go:generate ./generate-elements elements.txt elements.generated.go
