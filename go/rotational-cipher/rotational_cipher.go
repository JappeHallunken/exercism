package rotationalcipher

import "strings"

const min, max = 0, 26

func RotationalCipher(plain string, shiftKey int) string {
	if shiftKey < min || shiftKey > max {
		return plain
	}

	var sb strings.Builder

	for _, r := range plain {
		switch {
		case r >= 'A' && r <= 'Z':
			r2 := r + rune(shiftKey)
			if r2 > 'Z' {
				r2 -= max
			}
			sb.WriteRune(r2)

		case r >= 'a' && r <= 'z':
			r2 := r + rune(shiftKey)
			if r2 > 'z' {
				r2 -= max
			}
			sb.WriteRune(r2)
		default:
			sb.WriteRune(r)

		}
	}
	return sb.String()
}
