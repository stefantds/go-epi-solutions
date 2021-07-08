package group_equal_entries

func GroupByAge(people []Person) {
	countByAge := make(map[int]int)
	for _, p := range people {
		countByAge[p.Age]++
	}

	offsetByAge := make(map[int]int)
	offset := 0
	for a := range countByAge {
		offsetByAge[a] = offset
		offset += countByAge[a]
	}

	for len(offsetByAge) > 0 {
		for _, o := range offsetByAge {
			elem := people[o]
			correctPos := offsetByAge[elem.Age]
			people[o], people[correctPos] = people[correctPos], people[o]
			countByAge[elem.Age]--
			if countByAge[elem.Age] > 0 {
				offsetByAge[elem.Age] = correctPos + 1
			} else {
				delete(offsetByAge, elem.Age)
			}
		}
	}
}
