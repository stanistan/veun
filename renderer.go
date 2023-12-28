package veun

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
)

func RenderToHTML(ctx context.Context, r HTMLRenderable, errHandler any) (template.HTML, error) {
	return (&Renderable{r: r, eh: MakeErrorHandler(errHandler)}).Render(ctx)
}

func Render(ctx context.Context, v AsR) (template.HTML, error) {
	return R(v).Render(ctx)
}

type TemplateRenderable struct {
	Tpl  *template.Template
	Data any
}

func (v TemplateRenderable) AsHTML(_ context.Context) (template.HTML, error) {
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
