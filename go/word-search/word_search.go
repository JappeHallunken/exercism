package wordsearch

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	result := make(map[string][2][2]int, 0)
	for _, v := range words {
		result[v] = [2][2]int{{-1, -1}, {-1, -1}}
	}

	// create the 2D puzzle matrix
	matrix := make([][]rune, len(puzzle))
	for i, s := range puzzle {
		matrix[i] = []rune(s)
	}

	top, left := 0, 0
	bottom, right := len(puzzle)-1, len(puzzle[0])-1


	return nil, nil
}
