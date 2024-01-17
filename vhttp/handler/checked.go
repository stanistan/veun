package handler

import "net/http"

// Checked will continue on the handler chain if the
// first handler given ended up executing a 404.
//
// It does this by creating a temporary http.ResponseWriter proxy to
// so it can check if the handler was successful (not 404) or not.
// When it hits a 404, it will execute the next handler, otherwise
// it will replay the response to the real http.ResponseWriter.
func Checked(hs ...http.Handler) http.Handler {
	return &checked{
		handlers: hs,
		status:   http.StatusNotFound,
	}
}

func CheckedFor(status int, hs ...http.Handler) http.Handler {
	return &checked{
		handlers: hs,
		status:   status,
	}
}

type checked struct {
	handlers []http.Handler
	status   int
}

func (c *checked) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var w2 *responseWriter

	for _, next := range c.handlers {
		w2 = newResponseWriter()
		next.ServeHTTP(w2, r)
		if w2.status != c.status {
			break
		}
	}

	w2.WriteTo(w)
}

type responseWriter struct {
	written []byte
	status  int
	h       http.Header
}

func newResponseWriter() *responseWriter {
	return &responseWriter{h: http.Header{}}
}

var _ http.ResponseWriter = &responseWriter{}

func (wr *responseWriter) Write(bs []byte) (int, error) {
	wr.written = append(wr.written, bs...)
	return len(bs), nil
}

func (wr *responseWriter) WriteHeader(statusCode int) {
	wr.status = statusCode
}

func (wr *responseWriter) Header() http.Header {
	return wr.h
}

func (wr *responseWriter) WriteTo(w http.ResponseWriter) {
	if wr.status != 0 {
		w.WriteHeader(wr.status)
	}

	h := w.Header()
	for k, v := range wr.h {
		h[k] = v
	}

	_, _ = w.Write(wr.written)
}
