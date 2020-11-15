package reverse_words

func ReverseWords(input []rune) {
	n := len(input)

	// reverse the whole string
	reverse(input, 0, n-1)

	// reverse each word again
	start, finish := 0, 0
	for start < n {
		for start < n && input[start] == ' ' {
			start++
		}
		finish = start
		for finish < n && input[finish] != ' ' {
			finish++
		}
		reverse(input, start, finish-1)
		start = finish
	}
}

// reverse reverses the values of slice s between the start position
// and the end position (inclusive)
func reverse(s []rune, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
