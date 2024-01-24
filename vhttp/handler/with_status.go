package handler

import "net/http"

// WithStatus returns an [http.Handler] that _only_ writes the
// status code and does not set anything on the body.
func WithStatus(status int) http.Handler {
	return statusHandler{status: status}
}

type statusHandler struct {
	status int
}

func (h statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(h.status)
}
