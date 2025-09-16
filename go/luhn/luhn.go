package luhn

import (
	"strings"
	"unicode"
)

func Valid(id string) bool {

	id = strings.ReplaceAll(id, " ", "")

	if id == "0" {
		return false
	}

	nums := make([]int, 0, len(id))

	for _, r := range id {
		if !unicode.IsDigit(r) {
			return false
		}
		nums = append(nums, int(r-'0'))
	}

	for i := len(nums) - 2; i >= 0; i -= 2 {
		nums[i] = 2 * nums[i]
		if nums[i] > 9 {
			nums[i] -= 9
		}
	}
	return sum(nums)%10 == 0
}

func sum(nums []int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	return sum
}
