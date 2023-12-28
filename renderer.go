package veun

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
)

func RenderToHTML(ctx context.Context, r HTMLRenderable, errHandler any) (template.HTML, error) {
	var empty template.HTML

	if r == nil {
		return empty, nil
	}

	out, err := r.AsHTML(ctx)
	if err != nil {
		return handleRenderError(ctx, err, errHandler)
	}

	return out, nil
}

func Render(ctx context.Context, v AsRenderable) (template.HTML, error) {
	var empty template.HTML

	if v == nil {
		return empty, nil
	}

	r, err := v.Renderable(ctx)
	if err != nil {
		return handleRenderError(ctx, err, v)
	}

	out, err := RenderToHTML(ctx, r, v)
	if err != nil {
		return empty, err
	}

	return out, nil
}

type TemplateRenderable struct {
	Tpl  *template.Template
	Data any
}

func (v TemplateRenderable) RenderToHTML(_ context.Context) (template.HTML, error) {
	var empty template.HTML

	if v.Tpl == nil {
		return empty, fmt.Errorf("missing template")
	}

	var bs bytes.Buffer
	if err := v.Tpl.Execute(&bs, v.Data); err != nil {
		return empty, fmt.Errorf("tpl.Execute(): %w", err)
	}

	return template.HTML(bs.String()), nil
}
