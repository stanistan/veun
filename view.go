package veun

import (
	"context"
	"html/template"
)

type View struct {
	Tpl   *template.Template
	Slots Slots
	Data  any
}

func (v View) Template(ctx context.Context) (*template.Template, error) {
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

func MustParseTemplate(name, contents string) *template.Template {
	return template.Must(
		template.New(name).
			Funcs(template.FuncMap{"slot": slotFuncStub}).
			Parse(contents),
	)
}
