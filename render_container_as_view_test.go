package veun_test

import (
	"html/template"
	"testing"

	"github.com/alecthomas/assert/v2"

	. "github.com/stanistan/veun"
)

type ContainerView2 struct {
	Heading AsRenderable
	Body    AsRenderable
}

func (v ContainerView2) Renderable() (Renderable, error) {
	return View{
		Tpl:   containerViewTpl,
		Slots: Slots{"heading": v.Heading, "body": v.Body},
	}, nil
}

func TestRenderContainerAsView(t *testing.T) {
	html, err := Render(ContainerView2{
		Heading: ChildView1{},
		Body:    ChildView2{},
	})
	assert.NoError(t, err)
	assert.Equal(t, template.HTML(`<div>
	<div class="heading">HEADING</div>
	<div class="body">BODY</div>
</div>`), html)

}
