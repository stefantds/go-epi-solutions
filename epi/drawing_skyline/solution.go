package drawing_skyline

import (
	"math"

	"github.com/stefantds/go-epi-judge/test_utils"
)

func DrawingSkylines(buildings []Rect) []Rect {
	min := math.MaxInt64
	max := math.MinInt64

	for _, b := range buildings {
		min = test_utils.Min(min, b.Left)
		max = test_utils.Max(max, b.Right)
	}

	sk := make([]int, max-min+1)
	for _, b := range buildings {
		for j := b.Left - min; j <= b.Right-min; j++ {
			sk[j] = test_utils.Max(sk[j], b.Height)
		}
	}

	if len(sk) == 0 {
		return []Rect{}
	}

	result := make([]Rect, 0)
	currentRect := Rect{
		Left:   0 + min,
		Height: sk[0],
	}

	for i, h := range sk {
		if h != currentRect.Height {
			currentRect.Right = min + i - 1
			result = append(result, currentRect)
			currentRect = Rect{
				Left:   i + min,
				Height: h,
			}
		}
	}

	// add the rectangle that ends at sk[-1]
	currentRect.Right = min + len(sk) - 1
	result = append(result, currentRect)

	return result
}
