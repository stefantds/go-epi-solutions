package search_maze

func SearchMaze(maze [][]Color, s Coordinate, e Coordinate) []Coordinate {
	path := make([]Coordinate, 0)
	searchPath(maze, s, e, &path)
	return path
}

func searchPath(maze [][]Color, s Coordinate, e Coordinate, pathSoFar *[]Coordinate) bool {
	if s.X < 0 || s.Y < 0 || s.X >= len(maze) || s.Y >= len(maze[s.X]) || maze[s.X][s.Y] == Black {
		return false
	}

	maze[s.X][s.Y] = Black
	*pathSoFar = append(*pathSoFar, s)

	if s.Equal(e) {
		return true
	}

	for _, c := range []Coordinate{
		{X: s.X - 1, Y: s.Y},
		{X: s.X + 1, Y: s.Y},
		{X: s.X, Y: s.Y - 1},
		{X: s.X, Y: s.Y + 1},
	} {
		if searchPath(maze, c, e, pathSoFar) {
			return true
		}
	}

	*pathSoFar = (*pathSoFar)[:len(*pathSoFar)-1]
	return false
}
