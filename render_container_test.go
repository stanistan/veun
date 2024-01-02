package veun_test

import (
	"context"
	"html/template"
	"testing"

	"github.com/alecthomas/assert/v2"

	. "github.com/stanistan/veun"
)

type ContainerView struct {
	Heading, Body AsView
}

var containerViewTpl = MustParseTemplate("containerView", `<div>
	<div class="heading">{{ slot "heading" }}</div>
	<div class="body">{{ slot "body" }}</div>
</div>`)

func tplWithRealSlotFunc(ctx context.Context, tpl *template.Template, slots map[string]AsView) *template.Template {
	return tpl.Funcs(template.FuncMap{
		"slot": func(name string) (template.HTML, error) {
			slot, ok := slots[name]
			if ok {
				return Render(ctx, slot)
			}
			return template.HTML(""), nil
		},
	})
}

func (v ContainerView) AsHTML(ctx context.Context) (template.HTML, error) {
	return BasicTemplate{
		Tpl: tplWithRealSlotFunc(ctx, containerViewTpl, map[string]AsView{
			"heading": v.Heading,
			"body":    v.Body,
		}),
	}.AsHTML(ctx)
}

func (v ContainerView) View(_ context.Context) (*View, error) {
	return V(v), nil
}

var childViewTemplate = template.Must(
	template.New("childView").Parse(`{{ . }}`),
)

type ChildView1 struct{}

func (v ChildView1) View(_ context.Context) (*View, error) {
	return V(Template{Tpl: childViewTemplate, Data: "HEADING"}), nil
}

type ChildView2 struct{}

func (v ChildView2) View(_ context.Context) (*View, error) {
	return V(Template{Tpl: childViewTemplate, Data: "BODY"}), nil
}

func TestRenderContainer(t *testing.T) {
	html, err := Render(context.Background(), &ContainerView{
		Heading: ChildView1{},
		Body:    ChildView2{},
	})
	assert.NoError(t, err)
	assert.Equal(t, template.HTML(`<div>
	<div class="heading">HEADING</div>
	<div class="body">BODY</div>
</div>`), html)
}
