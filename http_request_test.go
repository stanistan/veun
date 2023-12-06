package veun_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/stanistan/veun"
)

func TestRequestBasicHandler(t *testing.T) {
	var handler = veun.RequestHandlerFunc(func(r *http.Request) (veun.AsRenderable, error) {
		return nil, nil
	})

	mux := http.NewServeMux()

	mux.Handle("/empty", handler)

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
}
