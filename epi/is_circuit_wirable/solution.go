package is_circuit_wirable

import "github.com/stefantds/go-epi-judge/utils"

func IsAnyPlacementFeasible(graph []GraphVertex) bool {
	if len(graph) == 0 {
		return true
	}

	for _, node := range graph {
		if node.VertexColor != none {
			continue // the node is part of a tree that was already traversed
		}

		if !isPlacementFeasible(&node) {
			return false
		}
	}

	return true
}

func isPlacementFeasible(startNode *GraphVertex) bool {
	bfsQueue := make(utils.Queue, 0)

	// assign once color to the starting vertex
	startNode.VertexColor = white
	bfsQueue = bfsQueue.Enqueue(startNode)

	for !bfsQueue.IsEmpty() {
		var value interface{}
		bfsQueue, value = bfsQueue.Dequeue()
		current := value.(*GraphVertex)

		for _, neighbor := range current.Edges {
			if neighbor.VertexColor == current.VertexColor {
				return false
			}

			// only add vertices without a color to the queue; the other ones are already visited
			if neighbor.VertexColor == none {
				neighbor.VertexColor = inverseColor(current.VertexColor)
				bfsQueue = bfsQueue.Enqueue(neighbor)
			}
		}
	}

	return true
}

func inverseColor(c Color) Color {
	switch c {
	case white:
		return black
	case black:
		return white
	default:
		panic("doesn't have a known inverse")
	}
}
