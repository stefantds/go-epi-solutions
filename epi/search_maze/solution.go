package search_maze

import "reflect"

func SearchMaze(maze [][]Color, s Coordinate, e Coordinate) []Coordinate {
	path := make([]Coordinate, 0)
	searchMazeHelper(maze, s, e, &path)
	return path
}

func searchMazeHelper(maze [][]Color, cur, e Coordinate, path *[]Coordinate) bool {
	if cur.X < 0 || cur.X >= len(maze) || cur.Y < 0 || cur.Y >= len(maze[0]) || maze[cur.X][cur.Y] != White {
		return false
	}

	*path = append(*path, cur)
	maze[cur.X][cur.Y] = Black

	if reflect.DeepEqual(cur, e) {
		return true
	}

	for _, nextMove := range []Coordinate{
		{cur.X, cur.Y + 1},
		{cur.X, cur.Y - 1},
		{cur.X + 1, cur.Y},
		{cur.X - 1, cur.Y},
	} {
		if searchMazeHelper(maze, nextMove, e, path) {
			return true
		}
	}

	// remove current element, since it didn't lead to the end
	*path = (*path)[0 : len(*path)-1]

	return false
}
