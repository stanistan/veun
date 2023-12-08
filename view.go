package veun

import (
	"context"
	"fmt"
	"html/template"
	"io/fs"
)

type View struct {
	Tpl   *template.Template
	Slots Slots
	Data  any
}

func (v View) Template(ctx context.Context) (*template.Template, error) {
	if v.Tpl == nil {
		return nil, fmt.Errorf("template missing")
	}

	return v.Slots.addToTemplate(ctx, v.Tpl), nil
}

func (v View) TemplateData(_ context.Context) (any, error) {
	return v.Data, nil
}

func (v View) Renderable(_ context.Context) (Renderable, error) {
	return v, nil
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
