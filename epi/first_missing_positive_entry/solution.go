package first_missing_positive_entry

func FindFirstMissingPositive(a []int) int {
	for i := 0; i < len(a); i++ {
		for a[i] > 0 && a[i] <= len(a) && a[a[i]-1] != a[i] {
			a[a[i]-1], a[i] = a[i], a[a[i]-1]
		}
	}

	for i, v := range a {
		if v != i+1 {
			return i + 1
		}
	}

	return len(a) + 1
}
