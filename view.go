package veun

import "html/template"

type View struct {
	Tpl   *template.Template
	Slots Slots
	Data  any
}

func (v View) Template() (*template.Template, error) {
	return v.Slots.addToTemplate(v.Tpl), nil
}

func (v View) TemplateData() (any, error) {
	return v.Data, nil
}

func (v View) Renderable() (Renderable, error) {
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
