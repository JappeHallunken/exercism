package prime

import (
	"fmt"
	"math"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {

	if n < 1 {
		return -1, fmt.Errorf("there is no zeroth prime")
	}
	count := 0
	num := 1
	for {
		num++
		if isPrime(num) {
			count++
			if count == n {
				return num, nil
			}
		}
	}
}

func isPrime(n int) bool {
	switch {
	case n < 2:
		return false
	case n == 2:
		return true
	case n%2 == 0:
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrtN; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
