package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	slowServer := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(20 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
		}))
	defer slowServer.Close()

	fastServer := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastUrl := fastServer.URL

	want := fastUrl
	got := Racer(slowURL, fastUrl)

	if got != want {
		t.Errorf("got %q, want %q", want, got)
	}
}
