package prime

func Factors(n int64) []int64 {
	var factors []int64
	var i int64 = 2
	for n > 1 {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
		if i != 2 {
			i += 2
		} else {
			i++
		}
	}
	return factors
}
