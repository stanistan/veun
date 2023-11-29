package veun_test

import (
	"html/template"
	"testing"

	"github.com/alecthomas/assert/v2"

	. "github.com/stanistan/veun"
)

func slotFuncStub(name string) (template.HTML, error) {
	return template.HTML(""), nil
}

type ContainerView struct {
	Heading Renderable
	Body    Renderable
}

func mustParseTemplate(name, contents string) *template.Template {
	return template.Must(
		template.New(name).
			Funcs(template.FuncMap{"slot": slotFuncStub}).
			Parse(contents),
	)
}

var containerViewTpl = mustParseTemplate("containerView", `<div>
	<div class="heading">{{ slot "heading" }}</div>
	<div class="body">{{ slot "body" }}</div>
</div>`)

func (v ContainerView) Template() (*template.Template, error) {
	return containerViewTpl.Funcs(template.FuncMap{
		"slot": func(name string) (template.HTML, error) {
			switch name {
			case "heading":
				return Render(v.Heading)
			case "body":
				return Render(v.Body)
			default:
				return template.HTML(""), nil
			}
		},
	}), nil
}

func (v ContainerView) TemplateData() (any, error) {
	return nil, nil
}

var childViewTemplate = template.Must(
	template.New("childView").Parse(`{{ . }}`),
)

type ChildView1 struct{}

func (v ChildView1) Template() (*template.Template, error) {
	return childViewTemplate, nil
}

func (v ChildView1) TemplateData() (any, error) {
	return "HEADING", nil
}

type ChildView2 struct{}

func (v ChildView2) Template() (*template.Template, error) {
	return childViewTemplate, nil
}

func (v ChildView2) TemplateData() (any, error) {
	return "BODY", nil
}

func TestRenderContainer(t *testing.T) {
	html, err := Render(&ContainerView{
		Heading: ChildView1{},
		Body:    ChildView2{},
	})
	assert.NoError(t, err)
	assert.Equal(t, template.HTML(`<div>
	<div class="heading">HEADING</div>
	<div class="body">BODY</div>
</div>`), html)
}
