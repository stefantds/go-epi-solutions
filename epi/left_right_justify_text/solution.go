package left_right_justify_text

import (
	"strings"

	"github.com/stefantds/go-epi-judge/test_utils"
)

func JustifyText(words []string, l int) []string {
	result := make([]string, 0)
	currentLineLen := 0
	currentLine := make([]string, 0)
	for _, w := range words {
		if currentLineLen+len(w)+len(currentLine) > l {
			// the current word doesn't fit on the current line
			remainingSpaces := l - currentLineLen
			wordsWithSpaces := test_utils.Max(len(currentLine)-1, 1)

			// distribute the remaining spaces
			for i := 0; i < remainingSpaces; i++ {
				currentLine[i%wordsWithSpaces] += " "
			}

			result = append(result, strings.Join(currentLine, ""))
			currentLine = make([]string, 0)
			currentLineLen = 0
		}

		// there is enough space for the current word on the current line
		currentLine = append(currentLine, w)
		currentLineLen += len(w)
	}

	// handle the last line
	lastLineBuilder := strings.Builder{}
	lastLineBuilder.WriteString(strings.Join(currentLine, " "))
	writeSpaces(&lastLineBuilder, l-currentLineLen-len(currentLine)+1)
	result = append(result, lastLineBuilder.String())

	return result
}

func writeSpaces(b *strings.Builder, count int) {
	for i := 0; i < count; i++ {
		b.WriteRune(' ')
	}
}
