package template

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	tt "text/template"

	"github.com/stanistan/veun/internal/view"
)

// T is an alias for the Template type.
type T = Template

// Template represents a template with slots
// that can be rendered as HTML.
//
// Template fulfills the [veun.HTMLRenderable] interface.
type Template struct {
	Tpl   *template.Template
	Slots Slots
	Data  any
}

var _ view.HTMLRenderable = Template{}

// AsHTML fulfills [veun.AsRenderable] for [Template].
//
// It will attempt to extract the template error from
// the golang library to make the error stack a bit
// more parseable when there is a failure.
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
