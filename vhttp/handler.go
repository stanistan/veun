package vhttp

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/stanistan/veun"
	"github.com/stanistan/veun/vhttp/request"
)

// Handler wraps request.Handler to conform it to http.Handler.
func Handler(r request.Handler, opts ...Option) http.Handler {
	return newHandler(r, opts)
}

// HandlerFunc wraps request.HandlerFunc to conform it to http.Handler.
func HandlerFunc(r request.HandlerFunc, opts ...Option) http.Handler {
	return newHandler(r, opts)
}

func newHandler(r request.Handler, opts []Option) handler {
	h := handler{RequestHandler: r, ErrorHandler: nil}
	for _, option := range opts {
		option(&h)
	}

	return h
}

// Option is an option that can be provided to vhttp handler.
type Option func(h *handler)

// WithErrorHandler is an Option that can be provided to Handler
// and HandlerFunc.
//
// This can change the default error handling behavior of the handler.
func WithErrorHandler(eh veun.ErrorHandler) Option {
	return func(h *handler) {
		h.ErrorHandler = eh
	}
}

// WithErrorHandlerFunc is the same as WithErrorHandler.
func WithErrorHandlerFunc(eh veun.ErrorHandlerFunc) Option {
	return WithErrorHandler(eh)
}

// handler implements http.Handler for a request.Handler.
type handler struct {
	RequestHandler request.Handler
	ErrorHandler   veun.ErrorHandler
}

// ServeHTTP implements http.Handler.
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v, next, err := h.RequestHandler.ViewForRequest(r)
	if err != nil {
		h.handleError(r.Context(), w, err)

		return
	}

	html, err := veun.Render(r.Context(), v)
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
	html, rErr := veun.RenderError(ctx, h.ErrorHandler, err)
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
