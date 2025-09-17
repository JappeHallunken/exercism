package darts

import "math"

func Score(x, y float64) int {
	d := getDistance(x, y)
	switch {
	case d > 10:
		return 0
	case d > 5:
		return 1
	case d > 1:
		return 5
	case d <= 1:
		return 10
  default:
		return -1
	}
}

func getDistance(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}
