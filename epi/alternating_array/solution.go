package alternating_array

import "sort"

func Rearrange(a []int) {
	sort.Ints(a)

	for i := 1; i < len(a)-1; i += 2 {
		a[i], a[i+1] = a[i+1], a[i]
	}
}
