package matrix_enclosed_regions

import (
	"github.com/stefantds/go-epi-judge/data_structures/queue"
)

type coordinates struct {
	x int
	y int
}

func FillSurroundedRegions(board [][]Color) {
	// mark the boundary regions
	for i := range board {
		markRegion(board, i, 0, White, Neutral)
		markRegion(board, i, len(board[i])-1, White, Neutral)
	}
	for j := range board[0] {
		markRegion(board, 0, j, White, Neutral)
		markRegion(board, len(board)-1, j, White, Neutral)
	}

	for i := range board {
		for j := range board[i] {
			markRegion(board, i, j, White, Black) // flip surrounded white regions to black
		}
	}

	for i := range board {
		for j := range board[i] {
			markRegion(board, i, j, Neutral, White) // flip boundary regions back to white
		}
	}
}

// markRegion marks the white region connected to position [i, j]
// It only replaces the "fromColor" color with the "toColor" one, leaving
// all other colors unchanged
func markRegion(board [][]Color, startX, startY int, fromColor, toColor Color) {
	q := make(queue.Queue, 0)
	q = q.Enqueue(coordinates{x: startX, y: startY})
	for !q.IsEmpty() {
		var value interface{}
		q, value = q.Dequeue()
		c := value.(coordinates)
		if c.x < 0 || c.x >= len(board) || c.y < 0 || c.y >= len(board[0]) {
			continue
		}
		if board[c.x][c.y] == fromColor {
			board[c.x][c.y] = toColor
			q = q.Enqueue(coordinates{x: c.x + 1, y: c.y})
			q = q.Enqueue(coordinates{x: c.x - 1, y: c.y})
			q = q.Enqueue(coordinates{x: c.x, y: c.y + 1})
			q = q.Enqueue(coordinates{x: c.x, y: c.y - 1})
		}
	}
}
