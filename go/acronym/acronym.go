package acronym

import (
	"strings"
)

func Abbreviate(s string) string {
	var out strings.Builder

	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")
	words := strings.Fields(s)

	for _, s := range words {
		runes := []rune(s)
		out.WriteRune(runes[0])
	}
	return strings.ToUpper(out.String())
}
