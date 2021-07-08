package string_decompositions_into_dictionary_words

func FindAllSubstrings(s string, words []string) []int {
	wordSize := len(words[0])
	numWords := len(words)
	sChars := []rune(s)
	wordsFreq := make(map[string]int)
	for _, v := range words {
		wordsFreq[v]++
	}

	result := make([]int, 0)

	matchAllWords := func(start int) bool {
		currFreq := make(map[string]int)
		for i := 0; i < numWords; i++ {
			currWord := string(sChars[start+i*wordSize : start+(i+1)*wordSize])
			if wordsFreq[currWord] == 0 {
				return false
			}
			currFreq[currWord]++
			if currFreq[currWord] > wordsFreq[currWord] {
				return false
			}
		}

		return true
	}

	for i := 0; i+wordSize*numWords <= len(sChars); i++ {
		if matchAllWords(i) {
			result = append(result, i)
		}
	}

	return result
}
