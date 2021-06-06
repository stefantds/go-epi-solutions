package substring_match

const (
	base = 26
)

func RabinKarp(t string, s string) int {
	sChars := []rune(s)
	tChars := []rune(t)

	searchedLength := len(sChars)

	if len(tChars) < len(sChars) {
		return -1
	}

	hashS := 0
	hashT := 0
	power := 1

	for i := 0; i < len(sChars); i++ {
		if i > 0 {
			power = power * base
		}
		hashS = hashS*base + charIdx(sChars[i])
		hashT = hashT*base + charIdx(tChars[i])
	}

	for i := len(sChars); i < len(tChars); i++ {
		if hashS == hashT && string(tChars[i-searchedLength:i]) == string(sChars) {
			return i - searchedLength
		}
		hashT -= charIdx(tChars[i-searchedLength]) * power
		hashT = hashT*base + charIdx(tChars[i])
	}

	if hashS == hashT && string(tChars[len(tChars)-searchedLength:]) == string(sChars) {
		return len(tChars) - searchedLength
	}

	return -1
}

func charIdx(r rune) int {
	return int(r - 'A')
}
