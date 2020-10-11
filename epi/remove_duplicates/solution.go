package remove_duplicates

import (
	"sort"
)

func EliminateDuplicate(names *[]Name) {
	if len(*names) < 2 {
		return
	}

	sort.Slice(*names, func(i, j int) bool {
		return (*names)[i].FirstName <= (*names)[j].FirstName
	})

	write := 0
	for read := 1; read < len(*names); read++ {
		if (*names)[read].FirstName != (*names)[write].FirstName {
			(*names)[write+1] = (*names)[read]
			write++
		}
	}

	*names = (*names)[:write+1]
}
