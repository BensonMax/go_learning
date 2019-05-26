package httptest

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

/*
在标准库中有一个 net/http/httptest 包，它可以让你轻易建立一个 HTTP 模拟服务器（mock HTTP server）。
我们改为使用模拟测试，这样我们就可以控制可靠的服务器来测试了。
*/

func TestHttp(t *testing.T) {

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
	got := slowUrl

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}

	slowServer.Close()
	fastServer.Close()
}
