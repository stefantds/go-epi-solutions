package string_integer_interconversion

import (
	"strings"
)

type StringIntegerConverter struct{}

func (c StringIntegerConverter) IntToString(x int) string {
	builder := strings.Builder{}
	reverseOrder := make([]rune, 0)
	var isNegative bool
	if x == 0 {
		return "0"
	}

	if x < 0 {
		isNegative = true
		x = -x
	}

	for x > 0 {
		reverseOrder = append(reverseOrder, '0'+rune(x%10))
		x = x / 10
	}

	if isNegative {
		reverseOrder = append(reverseOrder, '-')
	}

	for i := len(reverseOrder) - 1; i >= 0; i-- {
		builder.WriteRune(reverseOrder[i])
	}

	return builder.String()
}

func (c StringIntegerConverter) StringToInt(s string) int {
	var isNegative bool
	switch {
	case s[0] == '-':
		isNegative = true
		s = s[1:]
	case s[0] == '+':
		s = s[1:]
	}

	result := 0
	for _, c := range s {
		result = result*10 + int(c-'0')
	}

	if isNegative {
		return -result
	}

	return result
}
