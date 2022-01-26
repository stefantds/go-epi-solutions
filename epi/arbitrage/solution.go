package arbitrage

import "math"

func IsArbitrageExist(graph [][]float64) bool {
	for i := range graph {
		for j := range graph[i] {
			graph[i][j] = -math.Log10(graph[i][j])
		}
	}

	return hasNegativeWeightCycle(graph, 0)
}

// hasNegativeWeightCycle implements the Bellman-Ford algorithm to
// check if the graph contains a negative weight cycle
func hasNegativeWeightCycle(graph [][]float64, source int) bool {
	distToSource := make([]float64, len(graph))
	for i := range distToSource {
		distToSource[i] = math.MaxFloat64
	}
	distToSource[source] = 0.0

	for times := 1; times < len(graph); times++ {
		haveUpdate := false
		for i := 0; i < len(graph); i++ {
			for j := 0; j < len(graph[i]); j++ {
				if distToSource[j] > distToSource[i]+graph[i][j] {
					haveUpdate = true
					distToSource[j] = distToSource[i] + graph[i][j]
				}
			}
		}

		// no updates in a interation means no negative cycle
		if !haveUpdate {
			return false
		}
	}

	// check if there is any further updates: if yes, there is a cycle
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			if distToSource[j] > distToSource[i]+graph[i][j] {
				return true
			}
		}
	}

	return false
}
