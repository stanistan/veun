package veun

import (
	"context"
	"fmt"
	"html/template"
	"io/fs"
)

type Template struct {
	Tpl   *template.Template
	Slots Slots
	Data  any
}

func (v Template) AsHTML(ctx context.Context) (template.HTML, error) {
	out, err := TemplateRenderable{
		Tpl:  v.Slots.addToTemplate(ctx, v.Tpl),
		Data: v.Data,
	}.AsHTML(ctx)

	if err != nil {
		return out, fmt.Errorf("TemplateRenderable.AsHTML: %w", err)
	}

	return out, nil
}

func slotFuncStub(name string) (template.HTML, error) {
	return template.HTML(""), nil
}

func newTemplate(name string) *template.Template {
	return template.New(name).Funcs(template.FuncMap{
		"slot": slotFuncStub,
	})
}

func MustParseTemplate(name, contents string) *template.Template {
	return template.Must(newTemplate(name).Parse(contents))
}

func MustParseTemplateFS(f fs.FS, ps ...string) *template.Template {
	return template.Must(newTemplate("ROOT").ParseFS(f, ps...))
}
