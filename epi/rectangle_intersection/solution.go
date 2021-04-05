package rectangle_intersection

import utils "github.com/stefantds/go-epi-judge/test_utils"

func IntersectRectangle(r1 Rect, r2 Rect) Rect {
	if !isIntersect(r1, r2) {
		return Rect{X: 0, Y: 0, Height: -1, Width: -1}
	}
	return Rect{
		X:      utils.Max(r1.X, r2.X),
		Y:      utils.Max(r1.Y, r2.Y),
		Width:  utils.Min(r1.X+r1.Width, r2.X+r2.Width) - utils.Max(r1.X, r2.X),
		Height: utils.Min(r1.Y+r1.Height, r2.Y+r2.Height) - utils.Max(r1.Y, r2.Y),
	}
}

func isIntersect(r1, r2 Rect) bool {
	switch {
	case r1.Y+r1.Height < r2.Y:
		return false
	case r2.Y+r2.Height < r1.Y:
		return false
	case r1.X+r1.Width < r2.X:
		return false
	case r2.X+r2.Width < r1.X:
		return false
	default:
		return true
	}
}
