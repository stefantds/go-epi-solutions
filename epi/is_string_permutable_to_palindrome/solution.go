package is_string_permutable_to_palindrome

func CanFormPalindrome(s string) bool {
	charsFound := make(map[rune]bool)
	for _, c := range s {
		if ok, _ := charsFound[c]; ok {
			delete(charsFound, c)
		} else {
			charsFound[c] = true
		}
	}

	return len(charsFound) == 0 || len(charsFound) == 1
}
