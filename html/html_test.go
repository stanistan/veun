package html_test

import (
	"context"
	"fmt"
	"html/template"
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/stanistan/veun"
	"github.com/stanistan/veun/html"
)

func TestHTMLRender(t *testing.T) {
	for idx, testCase := range []struct {
		in  veun.AsView
		out string
	}{
		{
			html.Div(nil, nil),
			"<div></div>",
		},
		{
			html.Div(html.Attrs{"class": "blue"}, nil),
			`<div class="blue"></div>`,
		},
		{
			html.Text("<div>"),
			`&lt;div&gt;`,
		},
	} {
		testCase := testCase
		t.Run(fmt.Sprintf("test i-%d", idx), func(t *testing.T) {
			out, err := veun.Render(context.Background(), testCase.in)
			assert.NoError(t, err)
			assert.Equal(t, template.HTML(testCase.out), out)
		})
	}
}
