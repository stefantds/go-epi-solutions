package graph_clone

func CloneGraph(graph *GraphVertex) *GraphVertex {
	clones := make(map[*GraphVertex]*GraphVertex)
	clones[graph] = &GraphVertex{graph.Label, []*GraphVertex{}}

	queue := make([]*GraphVertex, 0)
	queue = append(queue, graph)

	for len(queue) > 0 {
		var v *GraphVertex
		v, queue = queue[0], queue[1:]
		for _, w := range v.Edges {
			if _, ok := clones[w]; !ok {
				c := &GraphVertex{w.Label, []*GraphVertex{}}
				clones[w] = c
				queue = append(queue, w)
			}
			(*clones[v]).Edges = append((*clones[v]).Edges, clones[w])
		}
	}

	return clones[graph]
}
