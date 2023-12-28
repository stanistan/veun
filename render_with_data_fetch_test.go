package veun_test

import (
	"context"
	"fmt"
	"html/template"
	"testing"
	"time"

	"github.com/alecthomas/assert/v2"

	. "github.com/stanistan/veun"
)

type ExpensiveViewData struct {
	Title string `json:"title"`
}

var expensiveViewTpl = MustParseTemplate("expensiveView", `{{ .Title }} success`)

type ExpensiveView struct {
	Data chan ExpensiveViewData
	Err  chan error
}

func NewExpensiveView(shouldErr bool, sleepFor time.Duration) *ExpensiveView {
	errCh := make(chan error)
	dataCh := make(chan ExpensiveViewData)

	go func() {
		defer func() {
			close(errCh)
			close(dataCh)
		}()

		// do data fetching and either write to
		// one thing or the other
		time.Sleep(sleepFor)
		if shouldErr {
			errCh <- fmt.Errorf("fetch failed")
		} else {
			dataCh <- ExpensiveViewData{Title: "hi"}
		}
	}()

	return &ExpensiveView{Data: dataCh, Err: errCh}
}

func (v *ExpensiveView) Renderable(ctx context.Context) (*Renderable, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case err := <-v.Err:
		return nil, err
	case data := <-v.Data:
		return R(View{Tpl: expensiveViewTpl, Data: data}), nil
	}
}

type ViewWithTimeout struct {
	Delegate AsR
	Timeout  time.Duration
}

func (v ViewWithTimeout) Renderable(ctx context.Context) (*Renderable, error) {
	ctx, cancel := context.WithTimeout(ctx, v.Timeout)
	defer cancel()

	return v.Delegate.Renderable(ctx)
}

func TestViewWithChannels(t *testing.T) {
	t.Run("successful", func(t *testing.T) {
		html, err := Render(context.Background(), NewExpensiveView(false, 1*time.Millisecond))
		assert.NoError(t, err)
		assert.Equal(t, template.HTML(`hi success`), html)
	})

	t.Run("failed", func(t *testing.T) {
		_, err := Render(context.Background(), NewExpensiveView(true, 1*time.Millisecond))
		assert.Error(t, err)
	})

	t.Run("context timed out", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
		defer cancel()

		_, err := Render(ctx, NewExpensiveView(false, 2*time.Millisecond))
		assert.Error(t, err)
	})

	t.Run("context timeout not reached", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
		defer cancel()

		_, err := Render(ctx, NewExpensiveView(false, 2*time.Millisecond))
		assert.NoError(t, err)
	})

	t.Run("with timeout and fallible", func(t *testing.T) {
		html, err := Render(context.Background(), FallibleView{
			Child: ViewWithTimeout{
				Delegate: NewExpensiveView(false, 10*time.Millisecond),
				Timeout:  2 * time.Millisecond,
			},
			CapturesErr: context.DeadlineExceeded,
		})
		assert.NoError(t, err)
		assert.Equal(t, template.HTML(`HEADING`), html)
	})
}
