package veun

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
)

type TemplateRenderable struct {
	Tpl  *template.Template
	Data any
}

func (v TemplateRenderable) RenderToHTML(_ context.Context) (template.HTML, error) {
	if v.Tpl == nil {
		return emptyHTML(), fmt.Errorf("missing template")
	}

	var buf bytes.Buffer
	if err := v.Tpl.Execute(&buf, v.Data); err != nil {
		return emptyHTML(), fmt.Errorf("tpl.Execute(): %w", err)
	}

	return template.HTML(buf.String()), nil
}
