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
	return RequestHandlerFunc(func(r *http.Request) (AsRenderable, error) {
		v, err := renderable.RequestRenderable(r)
		if err != nil {
			return nil, err
		}

		return html{Body: v}, nil
	})
}

func TestRequestRequestHandler(t *testing.T) {
	var empty = RequestRenderableFunc(func(r *http.Request) (AsRenderable, error) {
		return nil, nil
	})

	mux := http.NewServeMux()

	mux.Handle("/empty", RequestHandlerFunc(empty))
	mux.Handle("/html/empty", HTML(empty))

	mux.Handle("/person", RequestHandlerFunc(func(r *http.Request) (AsRenderable, error) {
		name := r.URL.Query().Get("name")
		if name == "" {
			return nil, fmt.Errorf("missing name")
		}

		return PersonView(Person{Name: name}), nil
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
