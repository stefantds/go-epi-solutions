package minimum_distance_3_sorted_arrays

type ArrayData struct {
	Val int
	Idx int
}

func less(a, b ArrayData) bool {
	if a.Val < b.Val {
		return true
	}
	if a.Val == b.Val {
		return a.Idx < b.Idx
	}

	return false
}

func more(a, b ArrayData) bool {
	return (a.Val > b.Val) ||
		(a.Val == b.Val && a.Idx > b.Idx)
}
