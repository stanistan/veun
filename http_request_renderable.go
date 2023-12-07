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

func HTTPHandlerFunc(r RequestRenderableFunc) http.Handler {
	return handler{Renderable: r}
}

func HTTPHandler(r RequestRenderable) http.Handler {
	return handler{Renderable: r}
}

// handler implements http.Handler for a RequestRenderable.
type handler struct {
	Renderable RequestRenderable
}

// ServeHTTP implements http.Handler.
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	renderable, next, err := h.Renderable.RequestRenderable(r)
	if err != nil {
		panic(err)
	}

	html, err := Render(r.Context(), renderable)
	if err != nil {
		panic(err)
	}

	if next != nil {
		next.ServeHTTP(w, r)
	}

	_, err = w.Write([]byte(html))
	if err != nil {
		panic(err)
	}
}
