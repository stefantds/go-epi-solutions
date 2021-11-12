package collatz_checker

func TestCollatzConjecture(n int) bool {
	// checked contains numbers that were verified already.
	// it only stores odd numbers, as for even numbers
	// the conjecture holds if all the numbers smaller have been
	// checked
	checked := make(map[int64]struct{})

	for i := 3; i <= n; i += 2 {
		i64 := int64(i)
		testI := i64
		seen := make(map[int64]struct{})
		for testI >= i64 {
			if _, ok := seen[testI]; ok {
				// we got into a loop
				return false
			}
			seen[testI] = struct{}{}

			if _, ok := checked[testI]; ok {
				// the value has been previously verified
				break
			}
			checked[testI] = struct{}{}

			if testI%2 == 0 {
				testI /= 2
			} else {
				next := 3*testI + 1
				if next < testI {
					panic("integer overflow")
				}
				testI = next
			}
		}
	}

	return true
}
