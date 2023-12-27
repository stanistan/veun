package veun

import (
	"context"
	"html/template"
)

// Slots is a key value map of AsRenderable items.
type Slots map[string]AsRenderable

func (s Slots) renderSlot(ctx context.Context) func(string) (template.HTML, error) {
	return func(name string) (template.HTML, error) {
		if s == nil {
			return emptyHTML(), nil
		}

		slot, _ := s[name]
		if slot == nil {
			return emptyHTML(), nil
		}

		return Render(ctx, slot)
	}
}

func (s Slots) addToTemplate(ctx context.Context, t *template.Template) *template.Template {
	if t == nil {
		return nil
	}
	return t.Funcs(template.FuncMap{"slot": s.renderSlot(ctx)})
}
