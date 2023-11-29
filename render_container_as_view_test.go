package veun_test

import (
	"html/template"
	"testing"

	"github.com/alecthomas/assert/v2"

	. "github.com/stanistan/veun"
)

func TestRenderContainerAsView(t *testing.T) {
	html, err := Render(View{
		Tpl: containerViewTpl,
		Slots: map[string]Renderable{
			"heading": ChildView1{},
			"body":    ChildView2{},
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, template.HTML(`<div>
	<div class="heading">HEADING</div>
	<div class="body">BODY</div>
</div>`), html)

}
