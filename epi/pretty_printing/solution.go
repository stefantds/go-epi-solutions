package pretty_printing

import (
	"math"

	"github.com/stefantds/go-epi-judge/test_utils"
)

func MinimumMessiness(words []string, lineLength int) int {
	if len(words) == 0 {
		return lineLength * lineLength
	}
	// cache[i] stores the minimum messiness for words[0:i+1]
	cache := make([]int, len(words))
	for i := 0; i < len(cache); i++ {
		cache[i] = math.MaxInt32
	}

	const spaceLen = 1

	remainingBlanks := lineLength - len(words[0])
	cache[0] = remainingBlanks * remainingBlanks

	for i := 1; i < len(words); i++ {
		remainingBlanks := lineLength - len(words[i])
		cache[i] = cache[i-1] + remainingBlanks*remainingBlanks

		for j := i - 1; j >= 0; j-- {
			remainingBlanks = remainingBlanks - len(words[j]) - spaceLen
			if remainingBlanks < 0 {
				break
			}
			var prevLinesMessiness int
			if j == 0 {
				prevLinesMessiness = 0
			} else {
				prevLinesMessiness = cache[j-1]
			}
			currentLineMessiness := remainingBlanks * remainingBlanks
			cache[i] = test_utils.Min(cache[i], currentLineMessiness+prevLinesMessiness)
		}
	}

	return cache[len(words)-1]
}
