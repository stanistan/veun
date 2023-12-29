package veun

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	tt "text/template"
)

type Slots map[string]AsR

func (s Slots) renderSlot(ctx context.Context) func(string) (template.HTML, error) {
	return func(name string) (template.HTML, error) {

		slot, ok := s[name]
		if ok && slot != nil {

			out, err := Render(ctx, slot)
			if err != nil {
				var tplError tt.ExecError
				if errors.As(err, &tplError) {
					return out, tplError.Unwrap()
				}

				return out, fmt.Errorf("named '%s': %w", name, err)
			}

			return out, nil
		}

		var empty template.HTML
		return empty, nil
	}
}

func (s Slots) addToTemplate(ctx context.Context, t *template.Template) *template.Template {
	if t == nil {
		return nil
	}

	return t.Funcs(template.FuncMap{"slot": s.renderSlot(ctx)})
}
