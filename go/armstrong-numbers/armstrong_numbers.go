package armstrong

import (
	"math"
)

func IsNumber(n int) bool {
	var length int
	switch n {
	case 0:
		length = 1
	default:
		length = int(math.Log10(float64(n))) + 1

	}
	slice := make([]int, length)

	num := n
	for i := length - 1; i >= 0; i-- {
		slice[i] = num % 10
		num /= 10
	}

	var sum int
	for _, v := range slice {
		sum += int(math.Pow(float64(v), float64(length)))
	}
	return sum == n

}
