package grains

import (
	"errors"
	"math"
)

const squares = 64

func Square(number int) (uint64, error) {
	switch {
	case number <= 0:
		return 0, errors.New("not a positive number") //makes no sense to allow negative amount of grains
	case number > squares:
		return 0, errors.New("a chessboard has only 64 squares")
	default:
		return uint64(math.Pow(2, float64(number-1))), nil
	}
}

func Total() uint64 {
	// var sum uint64
	// for i := 1; i <= squares; i++ {
	// 	sub, err := Square(i)
	// 	if err != nil {
	// 		return 0
	// 	}
	// 	sum += sub
	// }
	// return sum
	return 1<<64 - 1
}
