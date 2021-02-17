package advance_by_offsets

import (
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

func CanReachEnd(maxAdvanceSteps []int) bool {
	maxIdxReachable := 0
	for i := 0; i <= maxIdxReachable && i < len(maxAdvanceSteps); i++ {
		maxIdxReachable = utils.Max(maxIdxReachable, i+maxAdvanceSteps[i])
	}

	return maxIdxReachable >= len(maxAdvanceSteps)-1
}
