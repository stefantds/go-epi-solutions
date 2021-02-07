package combinations

func Combinations(n int, k int) [][]int {
	return combinationsHelper(1, n+1, k)
}

// combinationsHelper returns all the sets of size `setSize` with values
// between low (incl) and high (excl)
func combinationsHelper(low, high int, setSize int) [][]int {
	if setSize == 0 {
		// the empty set is the only set of size 0
		return [][]int{{}}
	}
	if high-low < setSize {
		// can't get a set of setSize with values between low and high
		return [][]int{}
	}

	withoutLow := combinationsHelper(low+1, high, setSize)

	withLow := make([][]int, 0)
	for _, c := range combinationsHelper(low+1, high, setSize-1) {
		withLow = append(withLow, append(c, low))
	}

	return append(withoutLow, withLow...)
}
