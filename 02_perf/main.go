package main

import (
	"net/http"
	_ "net/http/pprof"

	"optimizing-go/primes"
)

func main() {
	go http.ListenAndServe("localhost:6060", nil)

	for i := 0; i < 100000; i++ {
		_ = primes.Primes(1000000)
	}
}
