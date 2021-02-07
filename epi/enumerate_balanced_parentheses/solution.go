package enumerate_balanced_parentheses

func GenerateBalancedParentheses(numPairs int) []string {
	return generateParentheses("", numPairs, numPairs)
}

func generateParentheses(prefix string, leftParenNeeded int, rightParenNeeded int) []string {
	if rightParenNeeded == 0 {
		return []string{prefix}
	}

	result := make([]string, 0)
	if leftParenNeeded > 0 {
		result = append(result,
			generateParentheses(prefix+"(", leftParenNeeded-1, rightParenNeeded)...)
	}
	if leftParenNeeded < rightParenNeeded {
		result = append(result,
			generateParentheses(prefix+")", leftParenNeeded, rightParenNeeded-1)...)
	}

	return result
}
