package veun

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
)

func Render(ctx context.Context, r AsRenderable) (template.HTML, error) {
	renderable, err := r.Renderable(ctx)
	if err != nil {
		return handleRenderError(ctx, err, r)
	}

	out, err := render(ctx, renderable)
	if err != nil {
		return handleRenderError(ctx, err, r)
	}

	return out, nil
}

func render(ctx context.Context, r Renderable) (template.HTML, error) {
	var empty template.HTML

	tpl, err := r.Template(ctx)
	if err != nil {
		return empty, err
	}

	if tpl == nil {
		return empty, fmt.Errorf("missing template")
	}

	data, err := r.TemplateData(ctx)
	if err != nil {
		return empty, err
	}

	var bs bytes.Buffer
	if err := tpl.Execute(&bs, data); err != nil {
		return empty, err
	}

	return template.HTML(bs.String()), nil
}
