package veun_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/stanistan/veun"
	"github.com/stanistan/veun/el"
	"github.com/stanistan/veun/vhttp"
	"github.com/stanistan/veun/vhttp/request"
)

func HTML(rh request.Handler) http.Handler {
	return vhttp.HandlerFunc(func(r *http.Request) (veun.AsView, http.Handler, error) {
		v, next, err := rh.ViewForRequest(r)
		if err != nil {
			return nil, nil, err
		} else if v == nil {
			return nil, next, nil
		}

		return el.HTML().Content(
			el.Body().Content(v),
		), next, nil
	})
}

func newErrorView(_ context.Context, err error) (veun.AsView, error) {
	return el.Text("Error: " + err.Error()), nil
}

func TestHTTPHandler(t *testing.T) {

	var statusCode = func(code int) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(code)
		})
	}

	var empty = request.HandlerFunc(func(r *http.Request) (veun.AsView, http.Handler, error) {
		switch r.URL.Query().Get("not_found") {
		case "default":
			return nil, http.NotFoundHandler(), nil
		case "nil_404":
			return nil, statusCode(http.StatusNotFound), nil
		default:
			return nil, nil, nil
		}
	})

	var person = request.HandlerFunc(func(r *http.Request) (veun.AsView, http.Handler, error) {
		name := r.URL.Query().Get("name")
		if name == "" {
			return nil, nil, fmt.Errorf("missing name")
		}

		return PersonView(Person{Name: name}), nil, nil
	})

	mux := http.NewServeMux()

	mux.Handle("/empty", vhttp.Handler(empty))
	mux.Handle("/html/empty", HTML(empty))

	mux.Handle("/person", vhttp.Handler(person, vhttp.WithErrorHandlerFunc(newErrorView)))
	mux.Handle("/html/person", HTML(person))

	server := httptest.NewServer(mux)
	defer server.Close()

	var sendRequest = func(t *testing.T, to string) (string, int, error) {
		t.Helper()

		req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, server.URL+to, nil)
		assert.NoError(t, err)

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return "", 0, err
		}

		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		assert.NoError(t, err)

		return string(data), res.StatusCode, nil
	}

	t.Run("the root path is a real server that 404s", func(t *testing.T) {
		_, code, _ := sendRequest(t, "/")
		assert.Equal(t, 404, code)
	})

	t.Run("empty handler is indeed empty", func(t *testing.T) {
		body, code, _ := sendRequest(t, "/empty")
		assert.Equal(t, "", body)
		assert.Equal(t, 200, code)
	})

	t.Run("empty handler can 404", func(t *testing.T) {
		body, code, _ := sendRequest(t, "/empty?not_found=default")
		assert.Equal(t, "404 page not found\n", body)
		assert.Equal(t, 404, code)
	})

	t.Run("empty handler can 404 and nil", func(t *testing.T) {
		body, code, _ := sendRequest(t, "/empty?not_found=nil_404")
		assert.Equal(t, "", body)
		assert.Equal(t, 404, code)
	})

	t.Run("person renders (name=Stan)", func(t *testing.T) {
		body, code, _ := sendRequest(t, "/person?name=Stan")
		assert.Equal(t, "<div>Hi, Stan.</div>", body)
		assert.Equal(t, 200, code)
	})

	t.Run("person (name=)", func(t *testing.T) {
		body, code, _ := sendRequest(t, "/person?name=")
		assert.Equal(t, 500, code)
		assert.Equal(t, "Error: missing name", body)
	})

	t.Run("person renders (name=someone)", func(t *testing.T) {
		body, code, _ := sendRequest(t, "/person?name=someone")
		assert.Equal(t, "<div>Hi, someone.</div>", body)
		assert.Equal(t, 200, code)
	})

	t.Run("/html/empty", func(t *testing.T) {
		body, code, _ := sendRequest(t, "/html/empty")
		assert.Equal(t, "", body)
		assert.Equal(t, 200, code)
	})

	t.Run("/html/person (name=Stan)", func(t *testing.T) {
		body, code, _ := sendRequest(t, "/html/person?name=Stan")
		assert.Equal(t, "<html><body><div>Hi, Stan.</div></body></html>", body)
		assert.Equal(t, 200, code)
	})

	t.Run("/html/person (name=)", func(t *testing.T) {
		body, code, _ := sendRequest(t, "/html/person?name=")
		assert.Equal(t, "Internal Server Error\n", body)
		assert.Equal(t, 500, code)
	})
}
