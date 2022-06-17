package _1_bench

import (
	"testing"

	"optimizing-go/primes"
)

func BenchmarkPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primes.Primes(10000)
	}
}
