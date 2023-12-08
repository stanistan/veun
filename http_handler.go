package veun

import (
	"context"
	"log/slog"
	"net/http"
)

// HTTPHandler constructs an http.HTTPHandler given the RequestRenderable.
func HTTPHandler(r RequestRenderable, opts ...HandlerOption) http.Handler {
	h := handler{Renderable: r}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

// HTTPHandler constructs an http.HTTPHandler given the RequestRenderableFunc.
func HTTPHandlerFunc(r RequestRenderableFunc, opts ...HandlerOption) http.Handler {
	h := handler{Renderable: r}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

// HandlerOption is an option that can be provided to the handler.
type HandlerOption func(h *handler)

// WithErrorHandler creates a HandlerOption that can be provided to HTTPHandler
// or HTTPHandlerFunc.
//
// This can change the default error handling behavior of the handler.
func WithErrorHandler(eh ErrorRenderable) HandlerOption {
	return func(h *handler) {
		h.ErrorHandler = eh
	}
}

// WithErrorHandlerFunc is the same as WithErrorHandler.
func WithErrorHandlerFunc(eh ErrorRenderableFunc) HandlerOption {
	return WithErrorHandler(eh)
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
