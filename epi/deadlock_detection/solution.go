package deadlock_detection

func IsDeadlocked(graph []GraphVertex) bool {
	for _, g := range graph {
		if !g.Visited {
			if hasCycle(&g) {
				return true
			}
		}
	}

	return false
}

func hasCycle(v *GraphVertex) bool {
	if v.Visited {
		// vertex has been seen before -> cycle
		return true
	}
	v.Visited = true
	for _, w := range v.Edges {
		if hasCycle(w) {
			return true
		}
	}
	return false
}
