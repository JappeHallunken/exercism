package isbn

import "strings"

func IsValidISBN(isbn string) bool {

	isbn = removeHyphens(isbn)

	if len(isbn) != 10 {
		return false
	}

	nums := make([]int, 10)

	if isbn[len(isbn)-1] == 'X' {
		nums[9] = 10
		isbn = isbn[:9]
	}

	for i, v := range isbn {
		nums[i] = int(v - '0')
	}

	return calculateChecksum(nums) == 0
}

func removeHyphens(isbn string) string {
	return strings.ReplaceAll(isbn, "-", "")
}

func calculateChecksum(nums []int) int {
	sum := 0

	for i, v := range nums {
		sum += v * (10 - i)
	}

	return sum % 11
}
