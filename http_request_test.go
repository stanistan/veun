package veun_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert/v2"
	. "github.com/stanistan/veun"
)

var htmlTpl = MustParseTemplate("html", `<html><body>{{ slot "body" }}</body></html>`)

type html struct {
	Body AsRenderable
}

func (v html) Renderable(_ context.Context) (Renderable, error) {
	return View{Tpl: htmlTpl, Slots: Slots{"body": v.Body}}, nil
}

func HTML(renderable RequestRenderable) http.Handler {
	return HTTPHandlerFunc(func(r *http.Request) (AsRenderable, http.Handler, error) {
		v, next, err := renderable.RequestRenderable(r)
		if err != nil {
			return nil, next, err
		}

		return html{Body: v}, next, nil
	})
}

func TestRequestRequestHandler(t *testing.T) {
	var statusCode = func(code int) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(code)
		})
	}

	var empty = RequestRenderableFunc(func(r *http.Request) (AsRenderable, http.Handler, error) {
		switch r.URL.Query().Get("not_found") {
		case "default":
			return nil, http.NotFoundHandler(), nil
		case "nil_404":
			return nil, statusCode(http.StatusNotFound), nil
		default:
			return nil, nil, nil
		}
	})

	mux := http.NewServeMux()

	mux.Handle("/empty", HTTPHandler(empty))
	mux.Handle("/html/empty", HTML(empty))

	mux.Handle("/person", HTTPHandlerFunc(func(r *http.Request) (AsRenderable, http.Handler, error) {
		name := r.URL.Query().Get("name")
		if name == "" {
			return nil, nil, fmt.Errorf("missing name")
		}

		return PersonView(Person{Name: name}), nil, nil
	}))

	server := httptest.NewServer(mux)
	defer server.Close()

	var sendRequest = func(t *testing.T, to string) (string, int, error) {
		t.Helper()

		req, err := http.NewRequestWithContext(context.TODO(), "GET", server.URL+to, nil)
		assert.NoError(t, err)

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return "", 0, err
		}

		defer res.Body.Close()

		data, err := ioutil.ReadAll(res.Body)
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

	t.Run("person renders (name=someone)", func(t *testing.T) {
		body, code, _ := sendRequest(t, "/person?name=someone")
		assert.Equal(t, "<div>Hi, someone.</div>", body)
		assert.Equal(t, 200, code)
	})

	t.Run("/html/empty", func(t *testing.T) {
		body, code, _ := sendRequest(t, "/html/empty")
		assert.Equal(t, "<html><body></body></html>", body)
		assert.Equal(t, 200, code)
	})
}
