package railfence

import (
	"strings"
)

func Encode(message string, rails int) string {
	if len(message) < 2 || rails <= 1 {
		return message
	}

	var enc strings.Builder
	runes := []rune(message)

	if rails == 2 {
		for i := 0; i < len(runes); i += 2 {
			enc.WriteRune(runes[i])
		}
		for i := 1; i < len(runes); i += 2 {
			enc.WriteRune(runes[i])
		}
		return enc.String()
	}

	if rails > 2 {
		for i := 0; i < len(runes); i += 2 * (rails - 1) {
			enc.WriteRune(runes[i])
		}

		for row := 1; row < rails-1; row++ {
			i := row
			toggle := true
			step1 := 2 * (rails - row - 1)
			step2 := 2 * row
			for i < len(message) {
				enc.WriteRune(runes[i])
				if toggle {
					i += step1
				} else {
					i += step2
				}
				toggle = !toggle
			}
		}

		for i := rails - 1; i < len(runes); i += 2 * (rails - 1) {
			enc.WriteRune(runes[i])
		}
	}
	return enc.String()
}

func Decode(message string, rails int) string {
	if rails <= 1 || len(message) < 2 {
		return message
	}

	runes := []rune(message)
	var dec strings.Builder

	if rails == 2 {
		l := len(runes)
		ul := (l + 1) / 2

		upper := runes[:ul]
		lower := runes[ul:]

		for i := range ul {
			dec.WriteRune(upper[i])
			if i < len(lower) {
				dec.WriteRune(lower[i])
			}
		}
	}

	if rails > 2 {
		lengths := make([]int, rails)
		row, dir := 0, 1

		for range len(runes) {
			lengths[row]++
			row += dir
			if row == 0 || row == rails-1 {
				dir *= -1
			}
		}
		rows := make([][]rune, rails)
		pos := 0

		for r := range rails {
			rows[r] = runes[pos : pos+lengths[r]]
			pos += lengths[r]
		}

		row, dir = 0, 1

		for range runes {
			dec.WriteRune(rows[row][0])
			rows[row] = rows[row][1:]

			row += dir
			if row == 0 || row == rails-1 {
				dir *= -1
			}
		}
	}

	return dec.String()
}
