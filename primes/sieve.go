package primes

func Sieve(max int) []int {
	marked := make([]bool, max, max)

	for i := 2; i <= sqrt(max); i++ {
		for j := i * i; j < max; j = j + i {
			marked[j] = true
		}
	}

	var primes []int
	for i := 1; i < max; i++ {
		if !marked[i] {
			primes = append(primes, i)
		}
	}
	return primes
}
