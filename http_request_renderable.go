package veun

import (
	"context"
	"log/slog"
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

func HTTPHandlerFunc(r RequestRenderableFunc, opts ...HandlerOption) http.Handler {
	h := handler{Renderable: r}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

func HTTPHandler(r RequestRenderable, opts ...HandlerOption) http.Handler {
	h := handler{Renderable: r}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

type HandlerOption func(h *handler)

func WithErrorHandler(eh ErrorRenderable) HandlerOption {
	return func(h *handler) {
		h.ErrorHandler = eh
	}
}

func WithErrorHandlerFunc(eh ErrorRenderableFunc) HandlerOption {
	return func(h *handler) {
		h.ErrorHandler = eh
	}
}

// handler implements http.Handler for a RequestRenderable.
type handler struct {
	Renderable   RequestRenderable
	ErrorHandler ErrorRenderable
}

// ServeHTTP implements http.Handler.
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	renderable, next, err := h.Renderable.RequestRenderable(r)
	if err != nil {
		h.handleError(r.Context(), w, err)
		return
	}

	html, err := Render(r.Context(), renderable)
	if err != nil {
		h.handleError(r.Context(), w, err)
		return
	}

	if next != nil {
		next.ServeHTTP(w, r)
	}

	_, err = w.Write([]byte(html))
	if err != nil {
		panic(err)
	}
}

func (h handler) handleError(ctx context.Context, w http.ResponseWriter, err error) {
	html, rErr := handleRenderError(ctx, err, h.ErrorHandler)
	if rErr == nil && len(html) > 0 {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(html))
		return
	}

	// TODO: grab the logger from the context
	slog.Error("handler failed", "err", err)
	code := http.StatusInternalServerError
	http.Error(w, http.StatusText(code), code)
}
