package deadlock_detection

type VertexState int

type GraphVertex struct {
	Edges   []*GraphVertex
	Visited bool
}
