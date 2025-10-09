package allyourbase

import "fmt"

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {

	if inputBase < 2 {
		return nil, fmt.Errorf("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, fmt.Errorf("output base must be >= 2")
	}

	inputTrimmed := trimLeadingZeros(inputDigits)

	if len(inputTrimmed) == 0 || len(inputTrimmed) == 1 && inputTrimmed[0] == 0 {
		return []int{0}, nil
	}

	n := 0
	for _, d := range inputTrimmed {
		if d < 0 || d >= inputBase {
			return nil, fmt.Errorf("all digits must satisfy 0 <= d < input base")
		}
		n = n*inputBase + d
	}

	var output []int
	for n > 0 {
		output = append(output, n%outputBase)
		n /= outputBase
	}

	//invert slice
	for i, j := 0, len(output)-1; i < j; i, j = i+1, j-1 {
		output[i], output[j] = output[j], output[i]
	}
	return output, nil
}

func trimLeadingZeros(digits []int) []int {
	i := 0
	for i < len(digits) && digits[i] == 0 {
		i++
	}

	if i == len(digits) {
		return []int{0}
	}
	return digits[i:]
}
