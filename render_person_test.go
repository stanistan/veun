package veun_test

import (
	"html/template"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/stanistan/veun"
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

var _ veun.Renderable = &personView{}

var personViewTpl = template.Must(
	template.New("PersonView").Parse(`<div>Hi, {{ .Name }}.</div>`),
)

func (v *personView) Template() (*template.Template, error) {
	return personViewTpl, nil
}

func (v *personView) TemplateData() (any, error) {
	return v.Person, nil
}

func (v *personView) Renderable() (veun.Renderable, error) {
	return v, nil
}

func TestRenderPerson(t *testing.T) {
	html, err := veun.Render(PersonView(Person{Name: "Stan"}))
	assert.NoError(t, err)
	assert.Equal(t, html, template.HTML(`<div>Hi, Stan.</div>`))
}
