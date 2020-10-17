package smallest_nonconstructible_value

import (
	"sort"
)

func SmallestNonconstructibleValue(a []int) int {
	sort.Ints(a)
	smallestNC := 1

	for _, v := range a {
		if v <= smallestNC {
			// all values smaller than smallestNC can be created, so by adding v all values until
			// smallestNC+v can be constructed as well.
			smallestNC += v
		}
	}

	return smallestNC
}
