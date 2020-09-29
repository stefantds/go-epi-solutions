package int_as_array_increment

func PlusOne(a []int) []int {
	n := len(a) - 1
	a[n] += 1
	for i := n; i > 0 && a[i] == 10; i-- {
		a[i] %= 10
		a[i-1] += 1
	}
	if a[0] == 10 {
		a = append(a, 0)
		a[0] = 1
	}
	return a
}
