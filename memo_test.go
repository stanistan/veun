package veun_test

import (
	"context"
	"html/template"
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/stanistan/veun"
	"github.com/stanistan/veun/el"
)

func TestMemo(t *testing.T) {
	var (
		in   = el.Div{el.Text("memo")}
		view = veun.MustMemo(in)
	)

	assert.Equal(t, veun.Raw(`<div>memo</div>`), view)

	html, err := veun.Render(context.Background(), view)
	assert.NoError(t, err)
	assert.Equal(t, template.HTML(`<div>memo</div>`), html)

	html, err = veun.Render(context.Background(), in)
	assert.NoError(t, err)
	assert.Equal(t, template.HTML(`<div>memo</div>`), html)
}
