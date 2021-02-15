package is_circuit_wirable

type Color int

const (
	none  Color = 0
	white Color = 1
	black Color = 2
)

type GraphVertex struct {
	Edges       []*GraphVertex
	VertexColor Color
}
