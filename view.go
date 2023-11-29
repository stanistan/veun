package veun

import "html/template"

type View struct {
	Tpl   *template.Template
	Slots map[string]Renderable
	Data  any
}

func (v View) Template() (*template.Template, error) {
	return tplWithRealSlotFunc(v.Tpl, v.Slots), nil
}

func (v View) TemplateData() (any, error) {
	return v.Data, nil
}

func tplWithRealSlotFunc(tpl *template.Template, slots map[string]Renderable) *template.Template {
	return tpl.Funcs(template.FuncMap{
		"slot": func(name string) (template.HTML, error) {
			slot, ok := slots[name]
			if ok {
				return Render(slot)
			}
			return template.HTML(""), nil
		},
	})
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
