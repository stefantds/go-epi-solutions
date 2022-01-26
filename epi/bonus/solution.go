package bonus

import "github.com/stefantds/go-epi-judge/test_utils"

func CalculateBonus(productivity []int) int {
	if len(productivity) < 2 {
		return len(productivity)
	}

	bonuses := make([]int, len(productivity))
	for i := 0; i < len(bonuses); i++ {
		bonuses[i] = 1
	}

	for i := 1; i < len(bonuses); i++ {
		if productivity[i] > productivity[i-1] {
			bonuses[i] = bonuses[i-1] + 1
		}
	}
	for i := len(bonuses) - 1; i > 0; i-- {
		if productivity[i] < productivity[i-1] {
			bonuses[i-1] = test_utils.Max(bonuses[i]+1, bonuses[i-1])
		}
	}

	count := 0
	for _, b := range bonuses {
		count += b
	}

	return count
}
