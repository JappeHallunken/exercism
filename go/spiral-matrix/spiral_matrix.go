package spiralmatrix

type matrix [][]int

func SpiralMatrix(size int) [][]int {
	row, col := 0, 0
	dr, dc := 0, 1

	m := make(matrix, size)
	for i := range m {
		m[i] = make([]int, size)
	}

	for i := 1; i <= size*size; i++ {
		m[row][col] = i
		nextRow, nextCol := row+dr, col+dc
		if nextRow < 0 || nextRow >= size || nextCol < 0 || nextCol >= size || m[nextRow][nextCol] != 0 {
			dr, dc = dc, -dr // turn 90 degrees
		}
		row, col = row+dr, col+dc
	}
	return m
}
