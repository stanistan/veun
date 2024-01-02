package veun

import (
	"net/http"
)

// RequestHandler represents a method that can create a view out of an [http.Request].
type RequestHandler interface {
	ViewForRequest(r *http.Request) (AsRenderable, http.Handler, error)
}

// RequestHandlerFunc is the function representation of a [RequestHandler].
type RequestHandlerFunc func(*http.Request) (AsRenderable, http.Handler, error)

// ViewForRequest conforms [RequestRenderableFunc] to [RequestHandler].
func (f RequestHandlerFunc) ViewForRequest(r *http.Request) (AsRenderable, http.Handler, error) {
	return f(r)
}
