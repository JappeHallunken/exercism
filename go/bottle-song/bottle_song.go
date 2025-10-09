package bottlesong

import (
	"fmt"
	"strings"
)

func Recite(startBottles, takeDown int) []string {

	invalid := startBottles < 0 || takeDown < 0 || takeDown > 10 || startBottles > 10 || takeDown > startBottles
	if invalid {
		return nil
	}

	words := []string{"No", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}
	result := make([]string, 0, takeDown*5)
	line3 := "And if one green bottle should accidentally fall,"

	b := startBottles
	for i := 0; i < takeDown; i++ {
		if i > 0 {
			result = append(result, "")
		}

		bottleWord := "bottles"
		if b == 1 {
			bottleWord = "bottle"
		}
		nextBottleWord := "bottles"
		if b-1 == 1 {
			nextBottleWord = "bottle"
		}

		line1 := fmt.Sprintf("%s green %s hanging on the wall,", words[b], bottleWord)
		line4 := fmt.Sprintf("There'll be %s green %s hanging on the wall.", strings.ToLower(words[b-1]), nextBottleWord)

		result = append(result, line1, line1, line3, line4)
		b--
	}

	return result
}
