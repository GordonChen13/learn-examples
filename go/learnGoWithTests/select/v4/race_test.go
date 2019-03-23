package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T)  {
	t.Run("return fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)
		slowURL := slowServer.URL
		fastURL := fastServer.URL

		defer slowServer.Close()
		defer fastServer.Close()

		got := fastURL
		want ,err := Racer(slowURL, fastURL)
		if err != nil{
			t.Error("don't expect error here")
		}

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("timeout", func(t *testing.T) {
		slowServer := makeDelayedServer(25 * time.Millisecond)
		fastServer := makeDelayedServer(30 * time.Millisecond)
		slowURL := slowServer.URL
		fastURL := fastServer.URL

		defer slowServer.Close()
		defer fastServer.Close()

		_,err := ConfigurableRacer(slowURL, fastURL, 20*time.Millisecond)
		if err == nil{
			t.Error("expect timeout error here")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
