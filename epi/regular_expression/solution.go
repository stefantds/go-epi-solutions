package regular_expression

func IsMatch(regex string, s string) bool {
	regexChar, sChar := []rune(regex), []rune(s)
	if len(regex) > 0 && regex[0] == '^' {
		return isMatchAtPos(regexChar[1:], sChar)
	}
	for pos := 0; pos <= len(sChar); pos++ {
		if isMatchAtPos(regexChar, sChar[pos:]) {
			return true
		}
	}
	return false
}

func isMatchAtPos(regex []rune, s []rune) bool {
	if len(regex) == 0 {
		return true // the regex was completely matched
	}
	if regex[0] == '$' {
		return len(s) == 0 // s needs to be at the end of the string to match
	}
	if len(regex) >= 2 && regex[1] == '*' {
		if len(s) > 0 && isMatch(regex[0], s[0]) {
			return isMatchAtPos(regex, s[1:]) || // regex consumes the first character but can consume others
				isMatchAtPos(regex[2:], s[1:]) || // regex consumes just the first character
				isMatchAtPos(regex[2:], s) // regex consumes no character
		} else {
			return isMatchAtPos(regex[2:], s) // regex consumes no character but it is still a match (* can match no character)
		}
	}

	// this is a simple one character match
	if len(s) > 0 && isMatch(regex[0], s[0]) {
		return isMatchAtPos(regex[1:], s[1:])
	}

	return false
}

func isMatch(regexChar, sChar rune) bool {
	return regexChar == '.' || regexChar == sChar
}
