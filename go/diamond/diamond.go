package diamond

import (
	"fmt"
	"strings"
)

var ph string = " "

func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", fmt.Errorf("input must be a letter between 'A' and 'Z'")
	}

	var result strings.Builder
	limit := int(char - 'A') // for A it's 0, for Z it's 25

	if limit == 0 { // edge case 'A'
		result.WriteByte(char)
		return result.String(), nil
	}
	// upper half + middle line
	for i := 0; i <= limit; i++ {
		result.WriteString(drawLine(char, limit, i))
		result.WriteByte(byte('\n'))
	}
	//lower half
	for i := limit - 1; i >= 0; i-- {
		result.WriteString(drawLine(char, limit, i))
		if i != 0 { // no newLine on last line
			result.WriteByte(byte('\n'))
		}
	}
	return result.String(), nil
}

func drawLine(char byte, limit, i int) string {
	var line strings.Builder
	line.WriteString(strings.Repeat(ph, limit-i))
	line.WriteByte(byte(int(char) - limit + i))

	if i != 0 {
		line.WriteString(strings.Repeat(ph, i*2-1))
		line.WriteByte(byte(int(char) - limit + i))
	}
	line.WriteString(strings.Repeat(ph, limit-i))

	return line.String()
}
