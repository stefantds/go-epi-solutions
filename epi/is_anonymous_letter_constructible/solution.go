package is_anonymous_letter_constructible

func IsLetterConstructibleFromMagazine(letterText string, magazineText string) bool {
	letterCount := make(map[rune]int)
	for _, c := range magazineText {
		if _, ok := letterCount[c]; !ok {
			letterCount[c] = 1
		} else {
			letterCount[c] = letterCount[c] + 1
		}
	}

	for _, c := range letterText {
		if count, ok := letterCount[c]; !ok || count == 0 {
			return false
		}
		letterCount[c] = letterCount[c] - 1
	}

	return true
}
