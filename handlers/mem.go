package handlers

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func Mem(w http.ResponseWriter, _ *http.Request) {
	allocBlocks()
	w.WriteHeader(200)
}

func allocBlocks() {
	buf := []byte{}
	mb := 1024 * 1024

	n := rand.Intn(10)
	for i := 0; i < n; i++ {
		buf = append(buf, make([]byte, mb)...)
		log.Println("total allocated memory", len(buf))
		time.Sleep(10 * time.Millisecond)
	}
}
