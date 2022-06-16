package main

import (
	"optimizing-go/primes"
)

func main() {
	// TODO: Profile me.
	for i := 0; i < 1000; i++ {
		_ = primes.Primes(10000)
	}
}
