package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

func Encode(pt string) string {
	normalized := strings.ReplaceAll(pt, " ", "")
	normalized = strings.ToLower(normalized)

	if len(normalized) == 0 {
		return ""
	}

	re := regexp.MustCompile(`[^a-z0-9]+`)
	normalized = re.ReplaceAllString(normalized, "")

	pt_r := []rune(normalized)
	length := len(normalized)

	c := int(math.Ceil(math.Sqrt(float64(length))))
	r := int(math.Ceil(float64(length) / float64(c)))

	for len(pt_r) < r*c {
		pt_r = append(pt_r, ' ')
	}

	var out []rune
	for col := 0; col < c; col++ {
		for row := 0; row < r; row++ {
			idx := row*c + col
			out = append(out, pt_r[idx])
		}
		if col < c-1 {
			out = append(out, ' ')
		}
	}

	return string(out)
}
