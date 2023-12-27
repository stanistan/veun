package veun

import (
	"context"
	"html/template"
	"io/fs"
)

// MustParseTemplate constructs a template given a name and
// string contents. It injects stub a "slot" function
// implementation.
//
// This will panic if it is unable to parse the template.
func MustParseTemplate(name, contents string) *template.Template {
	return template.Must(newTemplate(name).Parse(contents))
}

// MustParseTemplateFS constructs a template given a filesystem.
//
// This will panic if it is unable to parse the fs.
func MustParseTemplateFS(f fs.FS, ps ...string) *template.Template {
	return template.Must(newTemplate("ROOT").ParseFS(f, ps...))
}

func newTemplate(name string) *template.Template {
	return Slots{}.addToTemplate(context.TODO(), template.New(name))
}
