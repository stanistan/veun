package veun_test

import (
	"context"
	"html/template"
	"testing"

	"github.com/alecthomas/assert/v2"

	. "github.com/stanistan/veun"
)

type ContainerView2 struct {
	Heading, Body AsView
}

func (v ContainerView2) View(ctx context.Context) (*View, error) {
	return V(Template{
		Tpl:   containerViewTpl,
		Slots: Slots{"heading": v.Heading, "body": v.Body},
	}), nil
}

func TestRenderContainerAsView(t *testing.T) {
	html, err := Render(context.Background(), ContainerView2{
		Heading: ChildView1{},
		Body:    ChildView2{},
	})
	assert.NoError(t, err)
	assert.Equal(t, template.HTML(`<div>
	<div class="heading">HEADING</div>
	<div class="body">BODY</div>
</div>`), html)
}

func BenchmarkRenderContainer(b *testing.B) {
	ctx := context.Background()

	b.Run("empty_container", func(b *testing.B) {
		container := ContainerView2{}
		for i := 0; i < b.N; i++ {
			_, _ = Render(ctx, container)
		}
	})

	b.Run("container_one_child", func(b *testing.B) {
		container := ContainerView2{Heading: ChildView1{}}
		for i := 0; i < b.N; i++ {
			_, _ = Render(ctx, container)
		}
	})

	b.Run("container_two_children", func(b *testing.B) {
		container := ContainerView2{Heading: ChildView1{}, Body: ChildView2{}}
		for i := 0; i < b.N; i++ {
			_, _ = Render(ctx, container)
		}
	})

	b.Run("container_nested", func(b *testing.B) {
		container := ContainerView2{
			Heading: ContainerView2{
				Heading: ChildView1{},
				Body:    ChildView2{},
			},
			Body: ChildView2{},
		}
		for i := 0; i < b.N; i++ {
			_, _ = Render(ctx, container)
		}
	})

	b.Run("container_nested_double", func(b *testing.B) {
		container := ContainerView2{
			Heading: ContainerView2{
				Heading: ChildView1{},
				Body:    ChildView2{},
			},
			Body: ContainerView2{
				Heading: ChildView1{},
				Body:    PersonView(Person{Name: "hi"}),
			},
		}
		for i := 0; i < b.N; i++ {
			_, _ = Render(ctx, container)
		}
	})

	b.Run("container_nested_more", func(b *testing.B) {
		container := ContainerView2{
			Heading: ContainerView2{
				Heading: ContainerView2{
					Heading: ContainerView2{
						Heading: ContainerView2{
							Heading: ChildView1{},
							Body:    PersonView(Person{Name: "hi"}),
						},
					},
					Body: ChildView2{},
				},
				Body: ChildView2{},
			},
			Body: ContainerView2{
				Heading: ChildView1{},
				Body:    PersonView(Person{Name: "hi"}),
			},
		}
		for i := 0; i < b.N; i++ {
			_, _ = Render(ctx, container)
		}
	})
}
