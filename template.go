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

func MustParseTemplate(name, contents string) *template.Template {
	return template.Must(newTemplate(name).Parse(contents))
}

func MustParseTemplateFS(f fs.FS, ps ...string) *template.Template {
	return template.Must(newTemplate("ROOT").ParseFS(f, ps...))
}

func newTemplate(name string) *template.Template {
	return Slots{}.addToTemplate(context.TODO(), template.New(name))
}

var ErrNilTemplate = errors.New("nil template")

// BasicTemplate encapsulates basic html templare rendering.
type BasicTemplate struct {
	Tpl  *template.Template
	Data any
}

// AsHTML transforms a BasicTemplate into html.
//
//nolint:gosec
func (v BasicTemplate) AsHTML(_ context.Context) (template.HTML, error) {
	var empty template.HTML

	if v.Tpl == nil {
		return empty, ErrNilTemplate
	}

	var bs bytes.Buffer
	if err := v.Tpl.Execute(&bs, v.Data); err != nil {
		return empty, fmt.Errorf("execute template: %w", err)
	}

	return template.HTML(bs.String()), nil
}

// Template encapsulates basic html template rendering, including Slots.
type Template struct {
	Tpl   *template.Template
	Slots Slots
	Data  any
}

func (v Template) AsHTML(ctx context.Context) (template.HTML, error) {
	out, err := BasicTemplate{Tpl: v.Slots.addToTemplate(ctx, v.Tpl), Data: v.Data}.AsHTML(ctx)
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

// Slots are a mapping from a string name to a "slot".
type Slots map[string]AsView

func (s Slots) renderSlot(ctx context.Context) func(string) (template.HTML, error) {
	return func(name string) (template.HTML, error) {
		out, err := Render(ctx, s[name])
		if err != nil {
			return out, fmt.Errorf("slot '%s': %w", name, err)
		}

		return out, nil
	}
}

func (s Slots) addToTemplate(ctx context.Context, t *template.Template) *template.Template {
	if t == nil {
		return nil
	}

	return t.Funcs(template.FuncMap{"slot": s.renderSlot(ctx)})
}

func tplName(t *template.Template) string {
	if t == nil {
		return "<nil>"
	}

	return t.Name()
}
