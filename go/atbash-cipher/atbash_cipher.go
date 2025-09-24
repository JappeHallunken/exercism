package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	var sb strings.Builder
	s = strings.ToLower(s)
	counter := 0

	for _, r := range s {
		encoded := ' '
		if unicode.IsLetter(r)  {
			encoded = flip(r)
			counter++
		}else if unicode.IsDigit(r) {
			encoded = r
			counter++
		}else {
			continue
		}
		sb.WriteRune(encoded)

		if counter%5 == 0 {
			sb.WriteRune(' ')
		}
	}
	return strings.TrimSpace(sb.String())
}

// flips the letter. a -> z | b -> y | e -> v
func flip(r rune) rune {
	return rune('a' + ('z' - r))
}
