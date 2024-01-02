package veun

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
)

func Render(ctx context.Context, v AsView) (template.HTML, error) {
	return V(v).Render(ctx)
}

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
