package smallest_subarray_covering_set

func FindSmallestSubarrayCoveringSet(paragraph []string, keywords map[string]struct{}) (start, end int) {
	start, end = 0, len(paragraph)-1
	keywordInRange := make(map[string]int)

	isKeyword := func(k string) bool {
		_, ok := keywords[k]
		return ok
	}

	addRangeKeyword := func(k string) {
		if _, ok := keywordInRange[k]; !ok {
			keywordInRange[k] = 0
		}
		keywordInRange[k]++
	}

	removeRangeKeyword := func(k string) {
		if _, ok := keywordInRange[k]; ok {
			if keywordInRange[k] == 1 {
				delete(keywordInRange, k)
			} else {
				keywordInRange[k]--
			}
		}
	}

	isAllKeywordsInRange := func() bool {
		return len(keywordInRange) == len(keywords)
	}

	for currentEnd, currentStart := 0, 0; currentEnd < len(paragraph); currentEnd++ {
		currentWord := paragraph[currentEnd]
		if isKeyword(currentWord) {
			addRangeKeyword(currentWord)
		}

		for isAllKeywordsInRange() {
			if end-start > currentEnd-currentStart {
				start, end = currentStart, currentEnd
			}
			currentWord := paragraph[currentStart]
			if isKeyword(currentWord) {
				removeRangeKeyword(currentWord)
			}
			currentStart++
		}
	}

	return start, end
}
