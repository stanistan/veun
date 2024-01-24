package template

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	tt "text/template"
)

type T = Template

type Template struct {
	Tpl   *template.Template
	Slots Slots
	Data  any
}

func (v Template) AsHTML(ctx context.Context) (HTML, error) {
	out, err := v.template(ctx).AsHTML(ctx)
	if err != nil {
		var tErr tt.ExecError
		if errors.As(err, &tErr) {
			if unwrapped := errors.Unwrap(tErr.Err); unwrapped != nil {
				err = unwrapped
			}
		}

		return out, fmt.Errorf("tpl '%s': %w", tplName(v.Tpl), err)
	}

	return out, nil
}

func (v Template) template(ctx context.Context) HTMLTemplate {
	return HTMLTemplate{
		Tpl:  v.Slots.addToTemplate(ctx, v.Tpl),
		Data: v.Data,
	}
}
