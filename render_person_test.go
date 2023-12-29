package veun_test

import (
	"context"
	"html/template"
	"testing"

	"github.com/alecthomas/assert/v2"
	. "github.com/stanistan/veun"
)

type Person struct {
	Name string
}

type personView struct {
	Person Person
}

func PersonView(person Person) *personView {
	return &personView{Person: person}
}

var personViewTpl = template.Must(
	template.New("PersonView").Parse(`<div>Hi, {{ .Name }}.</div>`),
)

func (v *personView) Renderable(_ context.Context) (*View, error) {
	return R(Template{Tpl: personViewTpl, Data: v.Person}), nil
}

func TestRenderPerson(t *testing.T) {
	html, err := Render(context.Background(), PersonView(Person{Name: "Stan"}))
	assert.NoError(t, err)
	assert.Equal(t, html, template.HTML(`<div>Hi, Stan.</div>`))
}

func BenchmarkRender(b *testing.B) {
	ctx := context.Background()
	person := PersonView(Person{Name: "Stan"})
	for i := 0; i < b.N; i++ {
		_, _ = Render(ctx, person)
	}
}
