package line_through_most_points

import (
	"github.com/stefantds/go-epi-judge/test_utils"
)

type slope struct {
	xDiff, yDiff int
}

func FindLineWithMostPoints(points []Point) int {
	result := 0
	for i, pi := range points {
		slopeTable := make(map[slope]int)
		overlapPoints := 1
		for j := i + 1; j < len(points); j++ {
			pj := points[j]

			// pj is the same point as pi, so no slope can be calculated.
			// we keep track of these points separately
			if pi.X == pj.X && pi.Y == pj.Y {
				overlapPoints++
				continue
			}
			s := slope{
				xDiff: pi.X - pj.X,
				yDiff: pi.Y - pj.Y,
			}
			if s.xDiff == 0 {
				s.yDiff = 1
			} else {
				gcd := gcd(s.xDiff, s.yDiff)
				s.xDiff /= gcd
				s.yDiff /= gcd
				if s.xDiff < 0 {
					s.xDiff = -s.xDiff
					s.yDiff = -s.yDiff
				}
			}
			slopeTable[s] += 1
		}

		// find the line with most points passing through pi
		var max = 0
		for _, v := range slopeTable {
			max = test_utils.Max(max, v)
		}

		// pi (And the overlapping points) are on all lines passing through pi
		maxPoints := overlapPoints + max
		result = test_utils.Max(result, maxPoints)
	}

	return result
}

// greatest common divisor (GCD) via Euclidean algorithm
func gcd(x, y int) int {
	for y != 0 {
		temp := y
		y = x % y
		x = temp
	}
	return x
}
