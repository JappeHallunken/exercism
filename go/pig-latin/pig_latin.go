package piglatin

import (
	"regexp"
	"strings"
)

type rule struct {
	re       *regexp.Regexp
	useGroup bool
}

var (
	rule1 = regexp.MustCompile(`^(xr|yt|ay|[aeiou])`)
	rule2 = regexp.MustCompile(`^([^aeiou]+)`)
	rule3 = regexp.MustCompile(`^([^aeiou]*qu)`)
	rule4 = regexp.MustCompile(`^([^aeiou]+)y`)
	rules = []rule{ // the sequence of tests matters, first match wins
		{rule1, false},
		{rule4, true},
		{rule3, true},
		{rule2, true},
	}
	suffix string = "ay"
)

func Sentence(sentence string) string {
	var result strings.Builder
	words := strings.Fields(strings.TrimSpace(sentence))

	for i, w := range words {
		if i > 0 {
			result.WriteRune(' ')
		}
		for _, r := range rules {
			if r.re.MatchString(w) {
				if r.useGroup {
					c := r.re.FindStringSubmatch(w)
					l := len([]rune(c[1]))
					runes := []rune(w)
					result.WriteString(string(runes[l:]) + c[1] + suffix)
				} else {
					result.WriteString(w + suffix)
				}
				break
			}
		}
	}
	return result.String()
}
