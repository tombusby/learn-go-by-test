package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	slowServer := makeDelayedServer(20)
	fastServer := makeDelayedServer(0)

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastUrl := fastServer.URL

	want := fastUrl
	got := Racer(slowURL, fastUrl)

	if got != want {
		t.Errorf("got %q, want %q", want, got)
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(delay * time.Millisecond)
			w.WriteHeader(http.StatusOK)
		}))
}
