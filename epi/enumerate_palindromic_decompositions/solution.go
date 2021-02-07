package enumerate_palindromic_decompositions

func PalindromeDecompositions(text string) [][]string {
	result := make([][]string, 0)
	doPalindromicDecomposition(text, 0, []string{}, &result)
	return result
}

func doPalindromicDecomposition(text string, offset int, partialResult []string, result *[][]string) {
	if offset == len(text) {
		newPartialRes := make([]string, len(partialResult))
		copy(newPartialRes, partialResult)
		*result = append(*result, newPartialRes)
		return
	}

	for i := offset + 1; i <= len(text); i++ {
		prefix := text[offset:i]
		if isPalindrome(prefix) {
			partialResult = append(partialResult, prefix)
			doPalindromicDecomposition(text, i, partialResult, result)
			partialResult = partialResult[:len(partialResult)-1]
		}
	}
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
