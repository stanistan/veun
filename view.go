package veun

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	tt "text/template"
)

type Template struct {
	Tpl   *template.Template
	Slots Slots
	Data  any
}

func (v Template) AsHTML(ctx context.Context) (template.HTML, error) {
	out, err := BasicTemplate{
		Tpl:  v.Slots.addToTemplate(ctx, v.Tpl),
		Data: v.Data,
	}.AsHTML(ctx)

	if err != nil {
		var tErr tt.ExecError
		if errors.As(err, &tErr) {
			err = errors.Unwrap(tErr.Err)
		}

		return out, fmt.Errorf("in template '%s': %w", v.Tpl.Name(), err)
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
