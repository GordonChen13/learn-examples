package service

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTestHandler(t *testing.T) {
	server := NewServer()

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	res := httptest.NewRecorder()

	server.ServeHTTP(res, req)

	got := res.Body.String()
	want := "This is a test1"

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
