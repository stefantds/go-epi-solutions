package run_length_compression

import (
	"strconv"
	"strings"
	"unicode"
)

type RLCompression struct{}

func (r *RLCompression) Decoding(s string) string {
	var result strings.Builder
	var currentCount int
	for _, c := range s {
		if unicode.IsDigit(c) {
			v, _ := strconv.Atoi(string(c))
			currentCount = currentCount*10 + v
		} else {
			for i := 0; i < currentCount; i++ {
				result.WriteRune(c)
			}
			currentCount = 0
		}
	}

	return result.String()
}

func (r *RLCompression) Encoding(s string) string {
	var result strings.Builder

	count := 1
	sChars := []rune(s)
	for i := 1; i < len(sChars); i++ {
		if sChars[i] == sChars[i-1] {
			count++
		} else {
			result.WriteString(strconv.Itoa(count))
			result.WriteRune(sChars[i-1])
			count = 1
		}
	}

	result.WriteString(strconv.Itoa(count))
	result.WriteRune(sChars[len(sChars)-1])

	return result.String()
}
