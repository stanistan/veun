package veun

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
)

func Render(ctx context.Context, r AsRenderable) (template.HTML, error) {
	if r == nil {
		return template.HTML(""), nil
	}

	renderable, err := r.Renderable(ctx)
	if err != nil {
		return handleRenderError(ctx, err, r)
	}

	out, err := renderable.RenderToHTML(ctx)
	if err != nil {
		return handleRenderError(ctx, err, r)
	}

	return out, nil
}

func RenderToHTML(tpl *template.Template, data any) (template.HTML, error) {
	var empty template.HTML

	if tpl == nil {
		return empty, fmt.Errorf("missing template")
	}

	var bs bytes.Buffer
	if err := tpl.Execute(&bs, data); err != nil {
		return empty, fmt.Errorf("tpl.Execute(): %w", err)
	}

	return template.HTML(bs.String()), nil
}
