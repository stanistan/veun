package veun_test

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"testing"

	"github.com/alecthomas/assert/v2"

	. "github.com/stanistan/veun"
)

type FailingView struct {
	Err error
}

func (v FailingView) Renderable(_ context.Context) (*Renderable, error) {
	return nil, fmt.Errorf("FailingView.Renderable(): %w", v.Err)
}

type FallibleView struct {
	CapturesErr error
	Child       AsR
}

func (v FallibleView) Renderable(_ context.Context) (*Renderable, error) {
	return R(v.Child).WithErrorHandler(v), nil
}

func (v FallibleView) ViewForError(ctx context.Context, err error) (AsR, error) {
	if v.CapturesErr == nil {
		return nil, err
	}

	if errors.Is(err, v.CapturesErr) {
		return ChildView1{}, nil
	}

	return nil, nil
}

func TestRenderContainerWithFailingView(t *testing.T) {
	_, err := Render(context.Background(), ContainerView2{
		Heading: ChildView1{},
		Body: FailingView{
			Err: fmt.Errorf("construction: %w", errSomethingFailed),
		},
	})
	assert.IsError(t, err, errSomethingFailed)
}

func TestRenderContainerWithCapturedError(t *testing.T) {
	t.Run("errors_bubble_out", func(t *testing.T) {
		_, err := Render(context.Background(), ContainerView2{
			Heading: ChildView1{},
			Body: FallibleView{
				Child: FailingView{Err: errSomethingFailed},
			},
		})
		assert.IsError(t, err, errSomethingFailed)
	})

	t.Run("errors_can_push_replacement_views", func(t *testing.T) {
		html, err := Render(context.Background(), ContainerView2{
			Heading: ChildView1{},
			Body: FallibleView{
				Child:       FailingView{Err: errSomethingFailed},
				CapturesErr: errSomethingFailed,
			},
		})
		assert.NoError(t, err)
		assert.Equal(t, template.HTML(`<div>
	<div class="heading">HEADING</div>
	<div class="body">HEADING</div>
</div>`), html)
	})

	t.Run("errors_can_return_nil_views", func(t *testing.T) {
		html, err := Render(context.Background(), ContainerView2{
			Heading: ChildView1{},
			Body: FallibleView{
				Child:       FailingView{Err: errors.New("hi")},
				CapturesErr: errSomethingFailed,
			},
		})
		assert.NoError(t, err)
		assert.Equal(t, template.HTML(`<div>
	<div class="heading">HEADING</div>
	<div class="body"></div>
</div>`), html)
	})

}

var errSomethingFailed = errors.New("an error")
