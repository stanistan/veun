package el

import (
	"context"
	"html/template"

	"github.com/stanistan/veun"
)

type elementKind interface {
	AsHTML(ctx context.Context, tag tag) (template.HTML, error)
}

type nodeChildren []veun.AsView

func (e nodeChildren) AsHTML(ctx context.Context, tag tag) (template.HTML, error) {
	content, err := veun.Render(ctx, veun.Views(e))
	if err != nil {
		return content, err
	}

	return tag.opening() + content + tag.closing(), nil
}

type void struct{}

func (v void) AsHTML(_ context.Context, tag tag) (template.HTML, error) {
	return tag.opening(), nil
}

type tag struct {
	name  string
	attrs Attrs
}

func (t tag) opening() template.HTML {
	return template.HTML("<" + t.name + t.attrs.render() + ">") //nolint:gosec
}

func (t tag) closing() template.HTML {
	return template.HTML("</" + t.name + ">") //nolint:gosec
}

func (t *tag) applyAttrs(fn func(Attrs)) {
	if t.attrs == nil {
		t.attrs = Attrs{}
	}

	fn(t.attrs)
}
