package request

import (
	"net/http"

	"github.com/stanistan/veun"
)

type Handler interface {
	ViewForRequest(r *http.Request) (veun.AsView, http.Handler, error)
}

type HandlerFunc func(*http.Request) (veun.AsView, http.Handler, error)

func (f HandlerFunc) ViewForRequest(r *http.Request) (veun.AsView, http.Handler, error) {
	return f(r)
}

func Always(v veun.AsView) Handler {
	return HandlerFunc(func(_ *http.Request) (veun.AsView, http.Handler, error) {
		return v, nil, nil
	})
}
