package twelve

import (
	"fmt"
	"strings"
)

var words = map[int][]string{
	1:  {"first", "a Partridge"},
	2:  {"second", "two Turtle Doves, and"},
	3:  {"third", "three French Hens,"},
	4:  {"fourth", "four Calling Birds,"},
	5:  {"fifth", "five Gold Rings,"},
	6:  {"sixth", "six Geese-a-Laying,"},
	7:  {"seventh", "seven Swans-a-Swimming,"},
	8:  {"eighth", "eight Maids-a-Milking,"},
	9:  {"ninth", "nine Ladies Dancing,"},
	10: {"tenth", "ten Lords-a-Leaping,"},
	11: {"eleventh", "eleven Pipers Piping,"},
	12: {"twelfth", "twelve Drummers Drumming,"},
}

func Verse(i int) string {
	if i < 0 || i > 12 {
		return ""
	}
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", words[i][0]))

	for j := i; j > 0; j-- {
		sb.WriteString(words[j][1] + " ")
	}
	sb.WriteString("in a Pear Tree.")

	return sb.String()
}

func Song() string {
	var sb strings.Builder

	for i := 1; i <= len(words); i++ {
		sb.WriteString(Verse(i))
		if i < len(words) {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}
