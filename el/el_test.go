package el_test

import (
	"context"
	"fmt"
	"html/template"
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/stanistan/veun"
	"github.com/stanistan/veun/el"
)

func TestHTMLRender(t *testing.T) {
	t.Parallel()
	for idx, testCase := range []struct {
		in  veun.AsView
		out string
	}{
		{
			el.Div(),
			"<div></div>",
		},
		{
			el.Div().Content(el.Span().InnerText("banana")),
			"<div><span>banana</span></div>",
		},
		{
			el.Div().
				Content(el.Span().InnerText("banana")).
				Class("foo").
				Attr("data-something", "else"),
			`<div class="foo" data-something="else"><span>banana</span></div>`,
		},
		{
			el.Text("<div>"),
			`&lt;div&gt;`,
		},
		{
			el.Br().In(el.Div()),
			`<div><br></div>`,
		},
		{
			el.Img().Attr("src", "/foo.png"),
			`<img src="/foo.png">`,
		},
	} {
		testCase := testCase
		t.Run(fmt.Sprintf("test i-%d", idx), func(t *testing.T) {
			t.Parallel()
			out, err := veun.Render(context.Background(), testCase.in)
			assert.NoError(t, err)
			assert.Equal(t, template.HTML(testCase.out), out)
		})
	}
}
