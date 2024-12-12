package handlers

import (
	"net/http"
	"sync/atomic"

	"github.com/polarsignals/pprof-example-app-go/fib"
)

var k atomic.Uint64

func init() {
	k.Store(100_000)
}

func Fib(w http.ResponseWriter, _ *http.Request) {
	n := fib.Fibonacci(uint(k.Add(1)))
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(n.String()))
}
