package veun

import (
	"net/http"
)

type RequestHandler interface {
	ViewForRequest(r *http.Request) (AsRenderable, http.Handler, error)
}

type RequestHandlerFunc func(*http.Request) (AsRenderable, http.Handler, error)

func (f RequestHandlerFunc) ViewForRequest(r *http.Request) (AsRenderable, http.Handler, error) {
	return f(r)
}
