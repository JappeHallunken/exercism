package yacht

import "bytes"

func Score(dice []int, category string) int {
	count := make([]byte, 6)
	for _, d := range dice {
		count[d-1]++
	}
	switch category {
	case "ones":
		return int(count[0])
	case "twos":
		return int(2 * count[1])
	case "threes":
		return int(3 * count[2])
	case "fours":
		return int(4 * count[3])
	case "fives":
		return int(5 * count[4])
	case "sixes":
		return int(6 * count[5])
	case "full house":
		if twos, threes := bytes.IndexByte(count, 2), bytes.IndexByte(count, 3); twos != -1 && threes != -1 {
			return (twos+1)*2 + (threes+1)*3
		}
	case "four of a kind":
		for i, d := range count {
			if d >= 4 {
				return (i + 1) * 4
			}
		}
	case "little straight":
		if bytes.Equal(count, []byte{1, 1, 1, 1, 1, 0}) {
			return 30
		}
	case "big straight":
		if bytes.Equal(count, []byte{0, 1, 1, 1, 1, 1}) {
			return 30
		}
	case "choice":
		sum := 0
		for _, d := range dice {
			sum += d
		}
		return sum
	case "yacht":
		if bytes.Contains(count, []byte{5}) {
			return 50
		}
	}
	return 0
}
