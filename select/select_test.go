package selected

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacker(t *testing.T) {

	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	want := fastUrl
	got := Racer(fastUrl, slowUrl)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	slowServer.Close()
	fastServer.Close()
}
