package look_and_say

import (
	"strconv"
	"strings"
)

func LookAndSay(n int) string {
	current := "1"
	for i := 1; i < n; i++ {
		current = lookAndSayPrevious(current)
	}

	return current
}

func lookAndSayPrevious(input string) string {
	var resultBuilder strings.Builder
	inputChars := []rune(input)
	if len(inputChars) == 0 {
		return ""
	}

	prevChar := inputChars[0]
	count := 1

	for i := 1; i < len(inputChars); i++ {
		if inputChars[i] == prevChar {
			count++
		} else {
			resultBuilder.WriteString(strconv.Itoa(count)) // write the count
			resultBuilder.WriteRune(prevChar)              // write the character counted
			count = 1
			prevChar = inputChars[i]
		}
	}

	resultBuilder.WriteString(strconv.Itoa(count)) // write the count
	resultBuilder.WriteRune(prevChar)              // write the character counted

	return resultBuilder.String()
}
