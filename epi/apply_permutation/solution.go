package apply_permutation

func ApplyPermutation(perm []int, a []int) {
	for idx := range perm {
		for perm[idx] != idx {
			newIdx := perm[idx]

			// move element at index idx to the right place (swap positions)
			a[idx], a[newIdx] = a[newIdx], a[idx]

			// perm make perm also reflect the swap
			perm[idx], perm[newIdx] = perm[newIdx], perm[idx]
		}
	}
}
