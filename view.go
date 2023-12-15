package veun

import (
	"context"
	"html/template"
	"io/fs"
)

type View struct {
	Tpl   *template.Template
	Slots Slots
	Data  any
}

func (v View) RenderToHTML(ctx context.Context) (template.HTML, error) {
	tpl := v.Tpl
	if v.Tpl != nil {
		tpl = v.Slots.addToTemplate(ctx, v.Tpl)
	}

	return RenderToHTML(tpl, v.Data)
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
