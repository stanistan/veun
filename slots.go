package veun

import (
	"context"
	"html/template"
)

type Slots map[string]AsRenderable

func (s Slots) renderSlot(ctx context.Context) func(string) (template.HTML, error) {
	return func(name string) (template.HTML, error) {
		slot, ok := s[name]
		if ok && slot != nil {
			return Render(ctx, slot)
		}

		var empty template.HTML
		return empty, nil
	}
}

func (s Slots) addToTemplate(ctx context.Context, t *template.Template) *template.Template {
	return t.Funcs(template.FuncMap{"slot": s.renderSlot(ctx)})
}
