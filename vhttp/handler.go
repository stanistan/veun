package vhttp

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/stanistan/veun/internal/view"
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
func WithErrorHandler(eh view.ErrorHandler) Option {
	return func(h *handler) {
		h.ErrorHandler = eh
	}
}

// WithErrorHandlerFunc is the same as WithErrorHandler.
func WithErrorHandlerFunc(eh view.ErrorHandlerFunc) Option {
	return WithErrorHandler(eh)
}

// handler implements http.Handler for a request.Handler.
type handler struct {
	RequestHandler request.Handler
	ErrorHandler   view.ErrorHandler
}

// ServeHTTP implements http.Handler.
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v, next, err := h.RequestHandler.ViewForRequest(r)
	if err != nil {
		h.handleError(r.Context(), w, err)

		return
	}

	html, err := view.Render(r.Context(), v)
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
	// N.B. if we successfully executed our error handler
	// and had some actual html output, we write/execute it.
	html, rErr := view.RenderError(ctx, h.ErrorHandler, err)
	if rErr == nil && len(html) > 0 {
		w.WriteHeader(errorCode)
		_, _ = w.Write([]byte(html))

		return
	} else if rErr != nil {
		slog.Error("veun error handler error", "err", rErr)
	}

	// If we can't execute a successful error handler,
	// we do the standard internal service error.
	slog.Error("veun unhandled http error", "err", err)
	http.Error(w, http.StatusText(errorCode), errorCode)
}

const (
	errorCode = http.StatusInternalServerError // default error code
)
