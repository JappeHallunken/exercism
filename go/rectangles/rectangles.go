package rectangles

func Count(diagram []string) int {
	count := 0
	rows := len(diagram)

	for r1, row := range diagram {
		var plusCols []int
		for c, ch := range row {
			if ch == '+' {
				plusCols = append(plusCols, c)
			}
		}

		for i := 0; i < len(plusCols); i++ {
			for j := i + 1; j < len(plusCols); j++ {
				c1, c2 := plusCols[i], plusCols[j]

				if !isHorizontalEdge(diagram, r1, c1, c2) {
					continue
				}

				for r2 := r1 + 1; r2 < rows; r2++ {
					if diagram[r2][c1] == '+' && diagram[r2][c2] == '+' &&
						isVerticalEdge(diagram, c1, r1, r2) &&
						isVerticalEdge(diagram, c2, r1, r2) &&
						isHorizontalEdge(diagram, r2, c1, c2) {
						count++
					}
				}
			}
		}
	}
	return count
}

func isHorizontalEdge(grid []string, r, c1, c2 int) bool {
	for c := c1 + 1; c < c2; c++ {
		ch := grid[r][c]
		if ch != '-' && ch != '+' {
			return false
		}
	}
	return true
}

func isVerticalEdge(grid []string, c, r1, r2 int) bool {
	for r := r1 + 1; r < r2; r++ {
		ch := grid[r][c]
		if ch != '|' && ch != '+' {
			return false
		}
	}
	return true
}
