package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	return s.response
}

func (s *SpyStore) Cancle() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	data := "Hello, World!"
	svr := Server(&SpyStore{data, false})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	svr.ServeHTTP(response, request)

	if response.Body.String() != data {
		t.Errorf("got %s, want %s", response.Body.String(), data)
	}

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "Hello, World"
		store := &SpyStore{data, false}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancelFunc := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancelFunc)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if !store.cancelled {
			t.Error("Store was not told to cancel")
		}
	})
}
