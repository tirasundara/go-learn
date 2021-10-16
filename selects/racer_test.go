package selects

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowURL := makeDelayedServer(20 * time.Millisecond)
	fastURL := makeDelayedServer(0 * time.Millisecond)
	defer slowURL.Close()
	defer fastURL.Close()

	got := Racer(slowURL.URL, fastURL.URL)
	want := fastURL.URL

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func makeDelayedServer(d time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(d)
		w.WriteHeader(http.StatusOK)
	}))
}
