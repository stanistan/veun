package veun_test

import (
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

func NewExpensiveView(shouldErr bool) *ExpensiveView {
	errCh := make(chan error)
	dataCh := make(chan ExpensiveViewData)

	go func() {
		defer func() {
			close(errCh)
			close(dataCh)
		}()

		// do data fetching and either write to
		// one thing or the other
		time.Sleep(1 * time.Millisecond)
		if shouldErr {
			errCh <- fmt.Errorf("fetch failed")
		} else {
			dataCh <- ExpensiveViewData{Title: "hi"}
		}
	}()

	return &ExpensiveView{Data: dataCh, Err: errCh}
}

func (v *ExpensiveView) Renderable() (Renderable, error) {
	select {
	case err := <-v.Err:
		return nil, err
	case data := <-v.Data:
		return View{Tpl: expensiveViewTpl, Data: data}, nil
	}
}

func TestViewWithChannels(t *testing.T) {
	t.Run("successful", func(t *testing.T) {
		html, err := Render(NewExpensiveView(false))
		assert.NoError(t, err)
		assert.Equal(t, template.HTML(`hi success`), html)
	})

	t.Run("failed", func(t *testing.T) {
		_, err := Render(NewExpensiveView(true))
		assert.Error(t, err)
	})
}
