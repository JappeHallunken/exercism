package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {

	freq := make(Frequency)

	if phrase == "" {
		return freq
	}

	parts := split(phrase)

	for _, part := range parts {
		if part != "" {
			freq[part]++
		}
	}
	return freq
}

func split(phrase string) []string {
	tphrase := strings.ToLower(phrase)
	re := regexp.MustCompile(`[a-zA-Z]+(?:'[a-zA-Z]+)?|\d+`)
	return re.FindAllString(tphrase, -1)
}
