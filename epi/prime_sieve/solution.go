package prime_sieve

func GeneratePrimes(n int) []int {
	candidates := make(map[int]bool, n+1)
	for i := 0; i <= n; i++ {
		candidates[i] = true
	}
	candidates[0] = false
	candidates[1] = false

	primes := make([]int, 0, n)

	for p := 2; p <= n; p++ {
		if candidates[p] {
			// we know that p is a prime because it was not sieved (i.e. is divisible) by any previous value
			primes = append(primes, p)

			for i := p * 2; i <= n; i = i + p {
				candidates[i] = false
			}
		}
	}

	return primes
}
