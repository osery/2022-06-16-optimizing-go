package _1_bench

import (
	"testing"

	"optimizing-go/primes"
)

func BenchmarkPrimes(b *testing.B) {
	// TODO: Benchmark the following call.
	primes.Primes(100)
}
