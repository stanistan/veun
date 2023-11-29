package veun

import (
	"bytes"
	"fmt"
	"html/template"
)

type RenderFunc func(r Renderable) (template.HTML, error)

type Renderable interface {
	Template() (*template.Template, error)
	TemplateData() (any, error)
}

func Render(r Renderable) (template.HTML, error) {
	var empty template.HTML

	tpl, err := r.Template()
	if err != nil {
		return empty, err
	}

	if tpl == nil {
		return empty, fmt.Errorf("missing template")
	}

	data, err := r.TemplateData()
	if err != nil {
		return empty, err
	}

	var bs bytes.Buffer
	if err := tpl.Execute(&bs, data); err != nil {
		return empty, err
	}

	return template.HTML(bs.String()), nil
}
