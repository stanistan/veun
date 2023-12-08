package veun

import (
	"net/http"
)

// RequestRenderable represents a method that
// can create a view out of an http.Request.
type RequestRenderable interface {
	RequestRenderable(r *http.Request) (AsRenderable, http.Handler, error)
}

// RequestRenderableFunc is the function representation of a
// RequestRenderable.
type RequestRenderableFunc func(*http.Request) (AsRenderable, http.Handler, error)

// RequestRenderable conforms RequestRenderableFunc to
// RequestRenderable interface.
func (f RequestRenderableFunc) RequestRenderable(r *http.Request) (AsRenderable, http.Handler, error) {
	return f(r)
}
