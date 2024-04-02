package template

import (
	"context"
	"fmt"
	"html/template"

	"github.com/stanistan/veun/internal/view"
)

// Slots are a mapping from a string name to a "slot".
type Slots map[string]view.AsView

func (s Slots) renderSlot(ctx context.Context) func(string) (HTML, error) {
	return func(name string) (HTML, error) {
		out, err := view.Render(ctx, s[name])
		if err != nil {
			return out, fmt.Errorf("slot '%s': %w", name, err)
		}

		return out, nil
	}
}

func (s Slots) addToTemplate(ctx context.Context, t *template.Template) *template.Template {
	if t == nil {
		return nil
	}

	return t.Funcs(template.FuncMap{"slot": s.renderSlot(ctx)})
}

func tplName(t *template.Template) string {
	if t == nil {
		return "<nil>"
	}

	return t.Name()
}
