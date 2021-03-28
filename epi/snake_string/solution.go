package snake_string

import "strings"

func SnakeString(s string) string {
	builder := strings.Builder{}
	runeS := []rune(s)
	for i := 1; i < len(runeS); i += 4 {
		builder.WriteRune(runeS[i])
	}
	for i := 0; i < len(runeS); i += 2 {
		builder.WriteRune(runeS[i])
	}
	for i := 3; i < len(runeS); i += 4 {
		builder.WriteRune(runeS[i])
	}

	return builder.String()
}
