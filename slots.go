package veun

import "html/template"

type Slots map[string]AsRenderable

func (s Slots) renderSlot(name string) (template.HTML, error) {
	slot, ok := s[name]
	if ok {
		return Render(slot)
	}

	var empty template.HTML
	return empty, nil
}

func (s Slots) addToTemplate(t *template.Template) *template.Template {
	return t.Funcs(template.FuncMap{"slot": s.renderSlot})
}
