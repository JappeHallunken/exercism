package transpose

func Transpose(input []string) []string {
	maxLen := 0

	for _, s := range input {
		if len(s) > maxLen {
			maxLen = len(s)
		}
	}

	output := make([]string, maxLen)
	for i, str := range input {
		for j, r := range str {
			for len(output[j]) < i {
				output[j] += " "
			}
			output[j] += string(r)
		}
	}
	return output
}
