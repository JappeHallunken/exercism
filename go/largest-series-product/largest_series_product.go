package lsproduct

import (
	"fmt"
	"unicode"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {

	if span < 0 {
		return -1, fmt.Errorf("span must not be negative")
	}
	if span > len(digits) {
		return -1, fmt.Errorf("span must be smaller than string length")
	}

	for _, r := range digits {
		if !unicode.IsDigit(r) {
			return -1, fmt.Errorf("digits input must only contain digits")
		}
	}

	var dSlice []string

	for i := 0; i < len(digits)-span+1; i++ {
		substr := digits[i : i+span]
		dSlice = append(dSlice, substr)
	}
	return doTheMath(dSlice), nil
}

func doTheMath(dSlice []string) int64 {
	var maximum int64 = 0

	for _, s := range dSlice {
		var prod int64 = 1

		for _, r := range s {
			digit := int64(r - '0')
			prod *= digit
		}

		if prod > maximum {
			maximum = prod
		}
	}
	return maximum
}
