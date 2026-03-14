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

// newTpl registers an empty slot func on the template so that
// {{ slot "name" }} calls are valid at parse time. The context
// passed here is intentionally context.TODO() because the slot
// map is empty and the func will never be invoked during parsing;
// the real per-request context is wired in via Slots.addToTemplate
// at render time.
func newTpl(name string) *template.Template {
	return Slots{}.addToTemplate(context.TODO(), template.New(name))
}
