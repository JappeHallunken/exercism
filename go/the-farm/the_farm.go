package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fodderCalculator FodderCalculator, cows int) (float64, error) {
	fodderAmount, err := fodderCalculator.FodderAmount(cows)
	if err != nil {
		return 0, err
	}
	fatteningFactor, err := fodderCalculator.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return fodderAmount * fatteningFactor / float64(cows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, cows int) (float64, error) {
  if cows > 0 {
		return DivideFood(fodderCalculator, cows)
	}
	return 0, errors.New("invalid number of cows")
}


type InvalidCowsError struct{
	cows int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cows int) error {
	switch {
	case cows < 0:
		return &InvalidCowsError{cows, "there are no negative cows"}
	case cows == 0:
		return &InvalidCowsError{cows, "no cows don't need food"}
	default:
		return nil
	}

}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
