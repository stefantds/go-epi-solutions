package is_valid_parenthesization

func IsWellFormed(s string) bool {
	matching := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	unmatched := make([]rune, 0)
	var lastUnmatched rune

	for _, r := range s {
		switch r {
		case '(', '[', '{':
			unmatched = append(unmatched, r)
		case ')', ']', '}':
			if len(unmatched) == 0 {
				// too many close parentheses
				return false
			}
			l := len(unmatched)
			unmatched, lastUnmatched = unmatched[:l-1], unmatched[l-1]
			if r != matching[lastUnmatched] {
				// not matching the last open
				return false
			}
		default:
			// other characters present
			return false
		}
	}

	// returns false if too many open parentheses
	return len(unmatched) == 0
}
