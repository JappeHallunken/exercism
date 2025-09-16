package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) == len(b) {
		diff := 0

		for i := range a {
			if a[i] != b[i] {
				diff++
			}
		}
		return diff, nil
	}
	return -1, errors.New("sequences differ in length")
}
