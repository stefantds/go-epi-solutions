package is_string_palindromic_punctuation

import "unicode"

func IsPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		switch {
		case !unicode.IsLetter(rune(s[i])) && !unicode.IsDigit(rune(s[i])):
			i++
		case !unicode.IsLetter(rune(s[j])) && !unicode.IsDigit(rune(s[j])):
			j--
		case unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])):
			return false
		default:
			i, j = i+1, j-1
		}
	}

	return true
}
