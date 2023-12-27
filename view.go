package veun

import (
	"context"
	"html/template"
)

type View struct {
	Tpl   *template.Template
	Slots Slots
	Data  any
}

// RenderToHTML implements Renderable for View.
func (v View) RenderToHTML(ctx context.Context) (template.HTML, error) {
	return TemplateRenderable{
		Tpl:  v.Slots.addToTemplate(ctx, v.Tpl),
		Data: v.Data,
	}.RenderToHTML(ctx)
}

// Renderable implements AsRenderable for View.
func (v View) Renderable(ctx context.Context) (Renderable, error) {
	return v, nil
}
