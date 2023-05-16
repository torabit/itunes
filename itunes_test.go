package itunes

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
)

func testClient(code int, body io.Reader, validators ...func(*http.Request)) (*Client, *httptest.Server) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, v := range validators {
			v(r)
		}
		w.WriteHeader(code)
		_, _ = io.Copy(w, body)
		r.Body.Close()
		if closer, ok := body.(io.Closer); ok {
			closer.Close()
		}
	}))
	client := &Client{
		http:    http.DefaultClient,
		baseURL: server.URL + "/",
	}
	return client, server
}

func testClientFile(code int, filename string, validators ...func(*http.Request)) (*Client, *httptest.Server) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return testClient(code, f, validators...)
}
