package sieve

import "sort"

func Sieve(limit int) []int {
	isPrime := makeMap(limit)
  var result []int
	for i, prime := range isPrime {
    if prime {
      result = append(result, i)
		}
	}
	sort.Ints(result)

	return result
}

// creates a map with keys from 2 up to the given limit
func makeMap(l int) map[int]bool {
	nums := make(map[int]bool, l-1)
	for i := 2; i <= l; i++ {
		if _, ok := nums[i]; !ok {
			nums[i] = true
			for j := i + i; j <= l; j += i {
				nums[j] = false
			}
		}
	}
	return nums
}
