package veun

import (
	"net/http"
)

type RequestHandler interface {
	ViewForRequest(r *http.Request) (AsView, http.Handler, error)
}

type RequestHandlerFunc func(*http.Request) (AsView, http.Handler, error)

func (f RequestHandlerFunc) ViewForRequest(r *http.Request) (AsView, http.Handler, error) {
	return f(r)
}
