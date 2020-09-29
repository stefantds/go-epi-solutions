package convert_base

import (
	"strings"
	"unicode"
)

func ConvertBase(numAsString string, b1 int, b2 int) string {
	isNegative := numAsString[0] == '-'
	var numAsInt int = 0

	if isNegative {
		numAsString = numAsString[1:]
	}

	numAsRunes := []rune(numAsString)

	for i := 0; i < len(numAsRunes); i++ {
		numAsInt *= b1
		currentDigit := numAsRunes[i]
		if unicode.IsDigit(currentDigit) {
			numAsInt += int(currentDigit - '0')
		} else {
			numAsInt += int(currentDigit-'A') + 10
		}
	}

	if numAsInt == 0 {
		return "0"
	}

	var resultBuilder strings.Builder
	if isNegative {
		resultBuilder.WriteRune('-')
	}

	constructFromBase(numAsInt, b2, &resultBuilder)

	return resultBuilder.String()
}

func constructFromBase(numAsInt int, base int, sBuilder *strings.Builder) {
	if numAsInt == 0 {
		return
	}

	constructFromBase(numAsInt/base, base, sBuilder)
	var currentRune rune

	if numAsInt%base >= 10 {
		currentRune = rune('A' + numAsInt%base - 10)
	} else {
		currentRune = rune('0' + numAsInt%base)
	}
	sBuilder.WriteRune(currentRune)
}
