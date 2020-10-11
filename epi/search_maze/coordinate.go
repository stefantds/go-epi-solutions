package search_maze

type Coordinate struct {
	X int
	Y int
}

func (c Coordinate) Equal(t Coordinate) bool {
	return c.X == t.X && c.Y == t.Y
}
