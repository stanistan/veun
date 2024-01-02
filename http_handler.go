package veun

import (
	"log/slog"
	"net/http"
)

// HTTPHandler constructs an [http.HTTPHandler] given a [RequestHandler].
func HTTPHandler(r RequestHandler, opts ...HandlerOption) http.Handler {
	h := handler{Request: r}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

// HTTPHandler constructs an http.HTTPHandler given the [RequestHandlerFunc].
func HTTPHandlerFunc(r RequestHandlerFunc, opts ...HandlerOption) http.Handler {
	h := handler{Request: r}
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
func WithErrorHandler(eh ErrorHandler) HandlerOption {
	return func(h *handler) {
		h.Error = eh
	}
}

// WithErrorHandlerFunc is the same as [WithErrorHandler].
func WithErrorHandlerFunc(eh ErrorHandlerFunc) HandlerOption {
	return WithErrorHandler(eh)
}

// handler implements [http.Handler] for a [RequestHandler].
//
// Use [HTTPHandler] or [HTTPHandlerFunc] to construct.
type handler struct {
	Request RequestHandler
	Error   ErrorHandler
}

// ServeHTTP implements [http.Handler].
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	renderable, next, err := h.Request.ViewForRequest(r)
	if err != nil {
		h.handleError(&Error{Err: err, ctx: r.Context()}, w)
		return
	}

	html, err := Render(r.Context(), renderable)
	if err != nil {
		h.handleError(&Error{Err: err, ctx: r.Context()}, w)
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

func (h handler) handleError(e *Error, w http.ResponseWriter) {
	html, rErr := RenderError(e, h.Error)
	if rErr == nil && len(html) > 0 {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(html))
		return
	}

	// TODO: grab the logger from the context
	slog.Error("handler failed", "err", e)
	code := http.StatusInternalServerError
	http.Error(w, http.StatusText(code), code)
}
