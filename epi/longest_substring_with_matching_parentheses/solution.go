package longest_substring_with_matching_parentheses

import "github.com/stefantds/go-epi-judge/test_utils"

func LongestMatchingParenthesesO1Space(s string) int {
	// maxMathingParenDirected iterates through a string (in the direction given by the
	// startPos, endPos and increment parameters) and calculates the longest valid substring
	// ending at each position.
	maxMathingParenDirected := func(s string, startPos, endPos, increment int, openChar rune) int {
		max := 0
		countClosing := 0
		countOpening := 0
		charS := []rune(s)
		for i := startPos; i != endPos; i += increment {
			if charS[i] == openChar {
				countOpening++
			} else {
				countClosing++
			}

			if countClosing > countOpening {
				countClosing = 0
				countOpening = 0
				continue
			}
			if countClosing == countOpening {
				max = test_utils.Max(max, countClosing+countOpening)
			}
		}

		return max
	}

	maxForward := maxMathingParenDirected(s, 0, len(s), 1, '(')
	maxBackward := maxMathingParenDirected(s, len(s)-1, -1, -1, ')')
	return test_utils.Max(maxForward, maxBackward)
}

func LongestMatchingParenthesesUsingStack(s string) int {
	// leftParenIndices will hold the position of the encountered opening parentheses
	leftParenIndices := make([]int, 0, len(s))

	// previousEnd holds the ending position of the previous valid substring
	previousEnd := -1
	max := 0
	for i, c := range s {
		if c == '(' {
			leftParenIndices = append(leftParenIndices, i)
		} else {
			if len(leftParenIndices) == 0 {
				previousEnd = i
			} else {
				// pop the element that was matched
				leftParenIndices = leftParenIndices[:len(leftParenIndices)-1]

				// the start is either the previously unmatched opening parenthesis
				// ot the beginning of the whole valid substring if all the opening
				// prentheses were matched
				var start int
				if len(leftParenIndices) == 0 {
					start = previousEnd
				} else {
					start = leftParenIndices[len(leftParenIndices)-1]
				}
				currentLength := i - start
				max = test_utils.Max(max, currentLength)
			}
		}
	}

	return max
}
