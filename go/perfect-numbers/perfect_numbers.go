package perfect

import "math"

// Define the Classification type here.
type Classification int

const (
	ClassificationAbundant Classification = iota
	ClassificationDeficient
	ClassificationPerfect
)

func Classify(n int64) (Classification, error) {
	limit := math.Sqrt(n)
}
