package wordy

import (
	"strconv"
	"strings"
)

var ops = map[string]func(int, int) (int, bool){
	"plus":          func(a, b int) (int, bool) { return a + b, true },
	"minus":         func(a, b int) (int, bool) { return a - b, true },
	"multiplied_by": func(a, b int) (int, bool) { return a * b, true },
	"divided_by": func(a, b int) (int, bool) {
		if b == 0 {
			return 0, false
		}
		return a / b, true
	},
}

func Answer(question string) (int, bool) {
	
	parts := strings.Fields(cleanString(question))

	// we should have a slice ["what", "is", <number>] possibly followed by <operation>, <number> ...
	if len(parts) < 3 || parts[0] != "what" || parts[1] != "is" || len(parts)%2 == 0 {
		return 0, false
	}

	for i, s := range parts[2:] {
		// even position: number, uneven position: operation
		if (i%2 == 0 && !isNumber(s)) || (i%2 == 1 && ops[s] == nil) {
			return 0, false
		}
	}

	result, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, false
	}

	for i := 3; i < len(parts); i += 2 {
		op := ops[parts[i]]
		num, err := strconv.Atoi(parts[i+1])
		if err != nil {
			return 0, false
		}

		r, ok := op(result, num)
		if !ok {
			return 0, false
		}
		result = r
	}
	return result, true
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func cleanString(s string) string {
	s = strings.ToLower(s)
	s = strings.TrimSuffix(s, "?")
	s = strings.ReplaceAll(s, "multiplied by", "multiplied_by")
	s = strings.ReplaceAll(s, "divided by", "divided_by")
	return s
}
