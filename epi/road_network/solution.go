package road_network

import (
	"math"

	"github.com/stefantds/go-epi-judge/test_utils"
)

func FindBestProposals(h []HighwaySection, p []HighwaySection, n int) HighwaySection {
	// build a graph with the given distances. The distance between nodes that are
	// not directly connected is set to MaxInt64. Distance from a node to itself is 0.
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				graph[i][j] = 0
			} else {
				graph[i][j] = math.MaxInt64
			}
		}
	}

	for _, v := range h {
		graph[v.X][v.Y] = v.Distance
		graph[v.Y][v.X] = v.Distance
	}

	// compute the shortest path between all the nodes
	floydWarshall(graph)

	bestImprovement := 0
	bestSection := HighwaySection{}

	// evaluate each proposal
	for _, v := range p {
		improvement := 0
		for a := 0; a < n; a++ {
			for b := 0; b < n; b++ {
				// check if the distance from a to b is improved by proposal v
				newABDistance := test_utils.Min(
					graph[a][v.X]+v.Distance+graph[v.Y][b],
					graph[a][v.Y]+v.Distance+graph[v.X][b],
				)
				if newABDistance < graph[a][b] {
					improvement += graph[a][b] - newABDistance
				}
			}
		}
		if improvement > bestImprovement {
			bestImprovement = improvement
			bestSection = v
		}
	}

	return bestSection
}

func floydWarshall(graph [][]int) {
	for k := 0; k < len(graph); k++ {
		for i := 0; i < len(graph); i++ {
			for j := 0; j < len(graph); j++ {
				if graph[i][k] != math.MaxInt64 && graph[k][j] != math.MaxInt64 {
					if graph[i][k]+graph[k][j] < graph[i][j] {
						graph[i][j] = graph[i][k] + graph[k][j]
					}
				}
			}
		}
	}
}
