package veun

import (
	"context"
	"log/slog"
	"net/http"
)

func HTTPHandler(r RequestHandler, opts ...HandlerOption) http.Handler {
	return newHandler(r, opts)
}

func HTTPHandlerFunc(r RequestHandlerFunc, opts ...HandlerOption) http.Handler {
	return newHandler(r, opts)
}

func newHandler(r RequestHandler, opts []HandlerOption) handler {
	h := handler{RequestHandler: r}
	for _, option := range opts {
		option(&h)
	}

	if h.ErrorHandler == nil {
		h.ErrorHandler = PassThroughErrorHandler()
	}

	return h
}

// HandlerOption is an option that can be provided to the handler.
type HandlerOption func(h *handler)

// WithErrorHandler creates a HandlerOption that can be provided to HTTPHandler
// or HTTPHandlerFunc.
//
// This can change the default error handling behavior of the handler.
func WithErrorHandler(eh ErrorHandler) HandlerOption {
	return func(h *handler) {
		h.ErrorHandler = eh
	}
}

// WithErrorHandlerFunc is the same as WithErrorHandler.
func WithErrorHandlerFunc(eh ErrorHandlerFunc) HandlerOption {
	return WithErrorHandler(eh)
}

// handler implements http.Handler for a RequestRenderable.
type handler struct {
	RequestHandler RequestHandler
	ErrorHandler   ErrorHandler
}

// ServeHTTP implements http.Handler.
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	renderable, next, err := h.RequestHandler.ViewForRequest(r)
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
	html, rErr := renderError(ctx, h.ErrorHandler, err)
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
