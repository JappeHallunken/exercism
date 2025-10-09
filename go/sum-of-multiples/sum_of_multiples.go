package summultiples


func SumMultiples(limit int, divisors ...int) int {
	n := len(divisors)
	slices := make([][]int, n)

	for i := range slices {
		div := divisors[i]
		if div == 0 {
			slices[i] = append(slices[i], 0)
			continue
		}
		for j := div; j < limit; j += div {
			slices[i] = append(slices[i], j)
		}
	}

	seen := make(map[int]bool)
	comb := []int{}

	for _, sl := range slices {
		for _, v := range sl {
			if !seen[v] {
				seen[v] = true
				comb = append(comb, v)
			}
		}
	}
	sum := 0

	for _, v := range comb {
		sum += v
	}

	return sum
}
