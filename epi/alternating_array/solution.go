package alternating_array

import "sort"

func Rearrange(a []int) {
	sort.Ints(a)

	for i := 1; i < len(a)-1; i += 2 {
		a[i], a[i+1] = a[i+1], a[i]
	}
}

func RearrangeLinearTime(a []int) {
	for i := 1; i < len(a); i++ {
		if i%2 == 0 && a[i-1] < a[i] ||
			i%2 == 1 && a[i-1] > a[i] {
			a[i], a[i-1] = a[i-1], a[i]
		}
	}
}
