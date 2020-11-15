package spreadsheet_encoding

func SsDecodeColID(col string) int {
	const base = 26 // nb of characters from 'A' to 'Z'

	result := 0
	for _, char := range col {
		result = result*base + int(char-'A') + 1
	}

	return result
}
