package next_permutation

func NextPermutation(perm []int) []int {
	if len(perm) <= 1 {
		return []int{} // no next permutation possible
	}

	// look for the position k in perm such that perm[k] < perm[k+1]
	pos := len(perm) - 2
	for pos >= 0 && perm[pos] >= perm[pos+1] {
		pos--
	}

	if pos == -1 {
		// this is the last permutation
		return []int{}
	}

	// look for the smallest value greater that perm[pos] in perm[pos+1:]
	// we use the fact that we know from the previous loop that perm[pos+1:] is sorted
	// in descending order.
	for i := len(perm) - 1; i > pos; i-- {
		if perm[i] > perm[pos] {
			// swap inversion point and value at i
			perm[i], perm[pos] = perm[pos], perm[i]
			break
		}
	}

	reverse(perm[pos+1:])

	return perm
}

func reverse(s []int) {
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}
}
