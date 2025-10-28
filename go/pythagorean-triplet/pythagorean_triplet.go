package pythagorean

import "math"

type Triplet [3]int

func Range(min, max int) []Triplet {
	var triplets []Triplet

	for a := min; a <= max-2; a++ {
		for b := a + 1; b <= max-1; b++ {
			c2 := a*a + b*b
			c := int(math.Sqrt(float64(c2)))

			if c*c == c2 && c <= max && c > b {
				triplets = append(triplets, Triplet{a, b, c})
			}
		}
	}
	return triplets
}

func Sum(p int) []Triplet {
	var r []Triplet
	tri := Range(1, p/2)
	for _, t := range tri {
		if t[0]+t[1]+t[2] == p {
			r = append(r, t)
		}
	}
	return r
}
