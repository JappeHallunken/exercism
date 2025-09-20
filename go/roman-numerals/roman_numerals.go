// Package romannumerals converts an integer to a roman numeral string
package romannumerals

import (
	"errors"
	"strconv"
)

var rn = []string{"I", "V", "X", "L", "C", "D", "M"}

func ToRomanNumeral(input int) (string, error) {
	for input < 1 || input > 3999 {
		return "", errors.New("not possible as roman number")
	}

	var result string
	inputStr := strconv.Itoa(input)
	length := len(inputStr)

	for i := 0; i < length; i++ {
		digit := int(inputStr[i] - '0')
		offset := (length - 1 - i) * 2 // 0 for ones, 2 for tens, 4 for hundreds, 6 for thousands
		result += check(digit, offset)
	}
	return result, nil
}


// "translates" arabic number to roman with a given number and and an offset to get the right latin letters
func check(num, offset int) string {
	var result string
	switch {
	case num < 4: // we dan only repeat one letter 3 times
		for i := 0; i < num; i++ { // so for a 2 the loop runs twice and adds "I" two times
			result += rn[offset]
		}
	case num == 4 && offset != 6: // if the offset is 6, we are in the thousands and there is no 4000 in roman numerals
		return rn[offset] + rn[offset+1]
	case num >= 5 && num < 9 && offset != 6: // same as in the first case, we can only repeat 3 times, like the number 8 is VIII
		result = rn[offset+1] // first letter is V, because the number is between 5 and 8
		for i := 0; i < num-5; i++ { 
			result += rn[offset]
		}
	case num == 9 && offset != 6:
		result = rn[offset] + rn[offset+2]
	}
	return result
}
