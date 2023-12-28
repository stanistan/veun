package veun_test

import (
	"context"
	"html/template"
	"testing"

	"github.com/alecthomas/assert/v2"

	. "github.com/stanistan/veun"
)

func Div(contents AsR) AsR {
	return Views{
		Raw("<div>"), contents, Raw("</div>"),
	}
}

func TestCommonViews(t *testing.T) {
	html, err := Render(context.Background(), Div(ChildView1{}))
	assert.NoError(t, err)
	assert.Equal(t, template.HTML("<div>HEADING</div>"), html)
}

func BenchmarkCommonViews(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Render(context.Background(), Div(ChildView1{}))
	}
}
