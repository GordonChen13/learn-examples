package http

import (
	"github.com/magiconair/properties/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestTestHandler(t *testing.T) {
	res := serveRequest(http.MethodGet, "/test", nil)

	got := res.Body.String()
	want := "This is a test"

	assert.Equal(t, res.Code, 200)
	assert.Equal(t, got, want)
}

func TestCreateMatch(t *testing.T) {
	name := "gordon"
	data := url.Values{}
	data.Set("name", name)

	res := serveRequest(http.MethodPost, "/match", strings.NewReader(data.Encode()))
	got := res.Body.String()
	want := name

	assert.Equal(t, res.Code, 201)
	assert.Equal(t, got, want)

}

func serveRequest(method, url string, body io.Reader) *httptest.ResponseRecorder {
	server := NewServer()

	req, _ := http.NewRequest(method, url, body)
	res := httptest.NewRecorder()

	switch method {
	case http.MethodPost:
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	server.ServeHTTP(res, req)
	return  res
}
