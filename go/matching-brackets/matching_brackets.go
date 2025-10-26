package brackets

import "strings"

const openBrackets = "([{"
const closeBrackets = ")]}"

func Bracket(input string) bool {
	var stack []rune

	for _, r := range input {
		idx := strings.IndexRune(openBrackets, r)
		if idx >= 0 {
			stack = append(stack, r)
			continue
		}

		idx = strings.IndexRune(closeBrackets, r)
		if idx >= 0 {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if strings.IndexRune(openBrackets, top) != idx {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
