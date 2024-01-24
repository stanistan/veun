// Package template includes [html/template] related
// functions for the veun library.
package template

import (
	"context"
	"html/template"
	"io/fs"
)

// HTML is a type alias for [template.HTML].
type HTML = template.HTML

// MustParse will panic if it cannot parse the string contents
// of the given template.
func MustParse(name, contents string) *template.Template {
	return template.Must(newTpl(name).Parse(contents))
}

// MustParseFS will panic if it cannot create/parse the
// file system given to it.
func MustParseFS(f fs.FS, ps ...string) *template.Template {
	return template.Must(newTpl("ROOT").ParseFS(f, ps...))
}

func newTpl(name string) *template.Template {
	return Slots{}.addToTemplate(context.TODO(), template.New(name))
}
