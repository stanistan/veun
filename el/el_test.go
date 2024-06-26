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

//nolint:funlen
func TestElementExpRender(t *testing.T) {
	t.Parallel()

	for idx, testCase := range []struct {
		in  veun.AsView
		out string
	}{
		{
			el.Div{},
			"<div></div>",
		},
		{
			el.Div{el.Span{el.Text("banana")}},
			"<div><span>banana</span></div>",
		},
		{
			el.Div{
				el.Attrs{
					"class":          "foo",
					"data-something": "else",
				},
				el.Span{
					el.Text("banana"),
				},
			},
			`<div class="foo" data-something="else"><span>banana</span></div>`,
		},
		{
			el.Text("<div>"),
			`&lt;div&gt;`,
		},
		{
			el.Div{el.Br{}},
			`<div><br></div>`,
		},
		{
			el.Img{el.Attr{"src", "/foo.png"}},
			`<img src="/foo.png">`,
		},
		{
			el.Script{
				el.Attr{"src", "some-file.js"},
			},
			`<script src="some-file.js"></script>`,
		},
		{
			el.Hr{},
			"<hr>",
		},
		{
			el.Div{
				el.Class("foo", "bar", "baz"),
			},
			`<div class="foo bar baz"></div>`,
		},
		{
			el.Div{
				el.Attr{"class", "banana"},
				el.Class("foo", "bar", "baz"),
			},
			`<div class="foo bar baz banana"></div>`,
		},
		{
			el.Div{
				el.Class("foo", "bar", "baz"),
				el.Attr{"class", "banana"},
			},
			`<div class="banana"></div>`,
		},
	} {
		testCase := testCase

		t.Run(fmt.Sprintf("test i-%d", idx), func(t *testing.T) {
			t.Parallel()

			out, err := veun.Render(context.Background(), testCase.in)

			assert.NoError(t, err)
			assert.Equal(t, template.HTML(testCase.out), out) //nolint:gosec
		})
	}
}
