package veun_test

import (
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

func (v *personView) Renderable() (Renderable, error) {
	return View{Tpl: personViewTpl, Data: v.Person}, nil
}

func TestRenderPerson(t *testing.T) {
	html, err := Render(PersonView(Person{Name: "Stan"}))
	assert.NoError(t, err)
	assert.Equal(t, html, template.HTML(`<div>Hi, Stan.</div>`))
}
