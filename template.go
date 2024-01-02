package veun

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	tt "text/template"
)

type BasicTemplate struct {
	Tpl  *template.Template
	Data any
}

func (v BasicTemplate) AsHTML(_ context.Context) (template.HTML, error) {
	var empty template.HTML

	if v.Tpl == nil {
		return empty, fmt.Errorf("nil template")
	}

	var bs bytes.Buffer
	if err := v.Tpl.Execute(&bs, v.Data); err != nil {
		return empty, fmt.Errorf("execute template: %w", err)
	}

	return template.HTML(bs.String()), nil
}

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

type Slots map[string]AsView

func (s Slots) renderSlot(ctx context.Context) func(string) (template.HTML, error) {
	return func(name string) (template.HTML, error) {

		slot, ok := s[name]
		if ok && slot != nil {

			out, err := Render(ctx, slot)
			if err != nil {
				var tplError tt.ExecError
				if errors.As(err, &tplError) {
					return out, tplError.Unwrap()
				}

				return out, fmt.Errorf("slot '%s': %w", name, err)
			}

			return out, nil
		}

		var empty template.HTML
		return empty, nil
	}
}

func (s Slots) addToTemplate(ctx context.Context, t *template.Template) *template.Template {
	if t == nil {
		return nil
	}

	return t.Funcs(template.FuncMap{"slot": s.renderSlot(ctx)})
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
