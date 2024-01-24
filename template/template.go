package template

import (
	"context"
	"html/template"
	"io/fs"
)

type HTML = template.HTML

func MustParse(name, contents string) *template.Template {
	return template.Must(newTpl(name).Parse(contents))
}

func MustParseFS(f fs.FS, ps ...string) *template.Template {
	return template.Must(newTpl("ROOT").ParseFS(f, ps...))
}

func newTpl(name string) *template.Template {
	return Slots{}.addToTemplate(context.TODO(), template.New(name))
}
