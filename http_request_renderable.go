package veun

import (
	"net/http"
)

// RequestRenderable represents a method that
// can create a view out of an http.Request.
type RequestRenderable interface {
	RequestRenderable(r *http.Request) (AsRenderable, error)
}

// RequestRenderableFunc is the function representation of a
// RequestRenderable.
type RequestRenderableFunc func(*http.Request) (AsRenderable, error)

// RequestRenderable conforms RequestRenderableFunc to
// RequestRenderable interface.
func (f RequestRenderableFunc) RequestRenderable(r *http.Request) (AsRenderable, error) {
	return f(r)
}

func RequestHandlerFunc(r RequestRenderableFunc) http.Handler {
	return HTTPHandler{Renderable: r}
}

func RequestHandler(r RequestRenderable) http.Handler {
	return HTTPHandler{Renderable: r}
}

// HTTPHandler implements http.Handler for a RequestRenderable.
type HTTPHandler struct {
	Renderable RequestRenderable
}

// ServeHTTP implements http.Handler.
func (h HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	renderable, err := h.Renderable.RequestRenderable(r)
	if err != nil {
		panic(err)
	}

	html, err := Render(r.Context(), renderable)
	if err != nil {
		panic(err)
	}

	_, err = w.Write([]byte(html))
	if err != nil {
		panic(err)
	}
}
