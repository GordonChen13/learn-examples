package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100*time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestHandler(t *testing.T) {
	data := "hello, world"
	t.Run("returns data from store", func(t *testing.T) {
		store := &SpyStore{response:data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}

		if store.cancelled {
			t.Errorf("it should not have cancelled the store")
		}
	})
	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		store := &SpyStore{response:data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5 * time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if !store.cancelled {
			t.Errorf("store was not told to cancel")
		}
	})
}

