package request

import (
	"net/http"

	"github.com/stanistan/veun/internal/view"
)

type Handler interface {
	ViewForRequest(r *http.Request) (view.AsView, http.Handler, error)
}

type HandlerFunc func(*http.Request) (view.AsView, http.Handler, error)

func (f HandlerFunc) ViewForRequest(r *http.Request) (view.AsView, http.Handler, error) {
	return f(r)
}

func Always(v view.AsView) Handler {
	return HandlerFunc(func(_ *http.Request) (view.AsView, http.Handler, error) {
		return v, nil, nil
	})
}
