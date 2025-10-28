package minesweeper

var directions = [8][2]int{
	{0, 1}, {1, 1}, {1, 0}, {1, -1}, // → ↘ ↓ ↙
	{0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, // ← ↖ ↑ ↗
}

// Annotate returns an annotated board
func Annotate(board []string) []string {
	m := make([][]rune, len(board))

	for i, row := range board {
		m[i] = []rune(row)
	}

	for row, runes := range m {
		for col, field := range runes {
			if field == '*' {
				continue
			}
			counter := 0
			for _, dir := range directions {
				r := row + dir[0]
				c := col + dir[1]
				if r >= 0 && r < len(m) && c >= 0 && c < len(m[0]) && m[r][c] == '*' {
					counter++
				}
			}
			if counter > 0 {
				m[row][col] = rune('0' + counter)
			}
		}
	}

	result := make([]string, len(m))
	for i, runes := range m {
		result[i] = string(runes)
	}
	return result
}
