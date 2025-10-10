package perfect

import (
	"errors"
)

type Classification int

const (
	ClassificationAbundant Classification = iota
	ClassificationDeficient
	ClassificationPerfect
)

var ErrOnlyPositive = errors.New("only positive numbers allowed")

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return 0, ErrOnlyPositive
	}
	if n == 1 {
		return ClassificationDeficient, nil
	}
	var sum int64 = 1 // 1 is always part of the sum

	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			sum += i // add number
			if i != n/i {
				sum += n / i // add partner, if not a quadratic number
			}
		}
	}
	switch {
	case sum == n:
		return ClassificationPerfect, nil
	case sum > n:
		return ClassificationAbundant, nil
	default:
		return ClassificationDeficient, nil
	}
}
