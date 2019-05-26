package httptest

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHttpV2(t *testing.T) {

	slowServer := MakeDelayServer(20 * time.Millisecond)
	fastServer := MakeDelayServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	want := fastUrl
	got := slowUrl

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

// 通过httptest 库生成一个指定延时的httpserver
func MakeDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
