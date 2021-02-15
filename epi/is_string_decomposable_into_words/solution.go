package is_string_decomposable_into_words

func DecomposeIntoDictionaryWords(domain string, dictionary map[string]struct{}) []string {
	// lastLength[i] will contain a value != -1 if domain[0:i+1] can be decomposed into valid values.
	// the value of lastLength[i] is the length of the word that ends at i.
	lastLength := make([]int, len(domain))
	for i := 0; i < len(lastLength); i++ {
		lastLength[i] = -1
	}

	for i := 0; i < len(domain); i++ {
		prefix := domain[0 : i+1]
		if _, ok := dictionary[prefix]; ok { // prefix found in dictionary
			lastLength[i] = i + 1
			continue
		}

		for j := 0; j < i; j++ {
			if lastLength[j] != -1 {
				substr := domain[j+1 : i+1]
				if _, ok := dictionary[substr]; ok {
					lastLength[i] = i - j
					break
				}
			}
		}
	}

	result := make([]string, 0)
	if lastLength[len(domain)-1] != -1 {
		idx := len(domain) - 1
		for idx >= 0 {
			result = append(result, domain[idx+1-lastLength[idx]:idx+1])
			idx -= lastLength[idx]
		}
	}

	reverse(result)
	return result
}

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
