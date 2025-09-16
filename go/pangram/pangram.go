package pangram

import "strings"

func IsPangram(input string) bool {

	input = strings.ToLower(input)

	for chr := 'a'; chr <= 'z'; chr++ {
		if !strings.ContainsRune(input, chr) {
			return false
		}
	}
	return true
}
