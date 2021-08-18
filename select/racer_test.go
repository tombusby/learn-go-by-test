package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("compares the speed of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, err := Racer(slowURL, fastUrl)

		if err != nil {
			t.Fatalf("Got an error when one wasn't expected %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", want, got)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10 seconds", func(t *testing.T) {
		serverA := makeDelayedServer(100 * time.Millisecond)
		serverB := makeDelayedServer(110 * time.Millisecond)

		defer serverA.Close()
		defer serverB.Close()

		_, err := ConfigurableRacer(serverA.URL, serverB.URL, 50*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(delay)
			w.WriteHeader(http.StatusOK)
		}))
}
