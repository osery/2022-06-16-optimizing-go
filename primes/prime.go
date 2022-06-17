package primes

import (
	"math"
)

func Primes(max int) []int {
	var primes []int

OUTER:
	for i := 1; i < max; i++ {
		max := sqrt(i)
		for j := 2; j <= max; j++ {
			if i%j == 0 {
				continue OUTER
			}
		}
		primes = append(primes, i)
	}
	return primes
}

func sqrt(i int) int {
	return int(math.Floor(math.Sqrt(float64(i))))
}
