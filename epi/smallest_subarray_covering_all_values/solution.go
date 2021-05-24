package smallest_subarray_covering_all_values

import "math"

func FindSmallestSequentiallyCoveringSubset(paragraph, keywords []string) (start, end int) {
	minLength := math.MaxInt64

	// keywordIdx maps a keyword to its position (index) in the keywords array
	keywordIdx := make(map[string]int)

	// keywordLastOccurrence keeps for a keyword k the last occurrence (index)
	// in the paragraph, with the additional constraint that all previous keywords
	// appear in the paragraph
	keywordLastOccurrence := make(map[string]int)

	// subarrayLength keeps for a keyword k the subarray length
	// for a subarray that contains all leywords until k
	subarrayLength := make(map[string]int)

	for i, k := range keywords {
		keywordIdx[k] = i
		keywordLastOccurrence[k] = -1
		subarrayLength[k] = -1
	}

	for i, w := range paragraph {
		idx, ok := keywordIdx[w]
		if !ok {
			// not a keyword, just ignore
			continue
		}

		keywordLastOccurrence[w] = i

		if idx == 0 { // the first keyword
			subarrayLength[w] = 1
		} else {
			prevKeyword := keywords[idx-1]
			if subarrayLength[prevKeyword] != -1 {
				subarrayLength[w] = subarrayLength[prevKeyword] + i - keywordLastOccurrence[prevKeyword]

				if idx == len(keywords)-1 { // last keyword
					if subarrayLength[w] < minLength {
						// we've found a shorter subarray covering all values
						end = i
						start = i - subarrayLength[w] + 1
						minLength = subarrayLength[w]
					}
				}
			}
		}
	}

	return start, end
}
