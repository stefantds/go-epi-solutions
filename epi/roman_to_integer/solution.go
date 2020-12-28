package roman_to_integer

import "fmt"

type romanDigit rune

const (
	I romanDigit = 'I'
	V romanDigit = 'V'
	X romanDigit = 'X'
	L romanDigit = 'L'
	C romanDigit = 'C'
	D romanDigit = 'D'
	M romanDigit = 'M'
)

var values = map[romanDigit]int{
	I: 1,
	V: 5,
	X: 10,
	L: 50,
	C: 100,
	D: 500,
	M: 1000,
}

func RomanToInteger(s string) int {
	result := 0
	previousValue := 0
	chars := []rune(s)
	for i := len(chars) - 1; i >= 0; i-- {
		digit := chars[i]
		value, ok := values[romanDigit(digit)]
		if !ok {
			panic(fmt.Errorf("expected a valid value, got %s", s))
		}

		if value < previousValue {
			result = result - value
		} else {
			result = result + value
		}
		previousValue = value
	}

	return result
}
