package main

import (
	"io"
	"net/http"
	"time"

	"github.com/ajwlforever/goutils/network"
	"github.com/ajwlforever/goutils/ratelimit"
)

func NetWork1() {
	go network.StartServer()
	network.StartClient()

	time.Sleep(10000)

}

var limiter *ratelimit.FixedWindowLimiter
var l ratelimit.Limiter

func handleHello(w http.ResponseWriter, r *http.Request) {
	if !limiter.TryAcquire().Ok {
		w.WriteHeader(http.StatusTooManyRequests)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "<h1>hello, world</h1>")
}

func FixedWindowLimiterTest() {
	limiter = ratelimit.NewFixedWindowLimiter(time.Second, 1)
	http.HandleFunc("/h", handleHello)
	http.ListenAndServe("0.0.0.0:8080", nil)

}

func handleHello1(w http.ResponseWriter, r *http.Request) {
	if !l.TryAcquire().Ok {
		w.WriteHeader(http.StatusTooManyRequests)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "<h1>Ooption mode</h1>")
}
func FixedWindowLimiteroptionTest() {
	l = ratelimit.NewLimiter(ratelimit.WithFixedWindowLimiter(time.Second, 1))

	http.HandleFunc("/h", handleHello1)
	http.ListenAndServe("0.0.0.0:8080", nil)

}

func main() {
	// network.PollURL()
	//network.FetchURL("https://www.jb51.net/softs/708254.html")
	//
	// network.StartWeb()
	FixedWindowLimiteroptionTest()
	for {
		time.Sleep(time.Second)
	}
}
