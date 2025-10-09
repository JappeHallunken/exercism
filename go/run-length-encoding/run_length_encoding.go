package encode

import (
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(input string) string {
	if input == "" {
		return ""
	}

	var result strings.Builder
	runes := []rune(input)
	count := 1
	writeGroup := func(ch rune, cnt int) {
		if cnt > 1 {
			result.WriteString(strconv.Itoa(cnt))
		}
		result.WriteRune(ch)
	}

	for i := range runes[1:] {
		if runes[i+1] == runes[i] {
			count++
		} else {
			writeGroup(runes[i], count)
			count = 1
		}
	}
	writeGroup(runes[len(runes)-1], count)
	return result.String()
}

func RunLengthDecode(input string) string {
	var result, n strings.Builder

	for _, r := range input {

		if unicode.IsDigit(r) {
			n.WriteRune(r)
		} else {
			if unicode.IsLetter(r) || unicode.IsSpace(r) {
				count, err := strconv.Atoi(n.String())
				if err != nil {
					count = 1
				}
				for i := 0; i < count; i++ {
					result.WriteRune(r)
				}
				n.Reset()
			}
		}
	}
	return result.String()
}
