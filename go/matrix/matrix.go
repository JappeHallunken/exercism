package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix [][]int

func New(s string) (Matrix, error) {
	rowsStr := strings.Split(s, "\n")
	if len(rowsStr) == 0 {
		return nil, fmt.Errorf("empty input")
	}

	var m Matrix
	var expectedCols int

	for i, row := range rowsStr {
		row = strings.TrimSpace(row)
		if row == "" {
			return nil, fmt.Errorf("row %d is empty", i+1)
		}
		fields := strings.Fields(row)

		if i == 0 {
			expectedCols = len(fields)
		} else if len(fields) != expectedCols {
			return nil, fmt.Errorf("row %d has different length", i+1)
		}

		var nums []int
		for j, f := range fields {
			n, err := strconv.Atoi(f)
			if err != nil {
				return nil, fmt.Errorf("row %d col %d not a number", i+1, j+1)
			}
			nums = append(nums, n)
		}
		m = append(m, nums)
	}

	return m, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	if len(m) == 0 {
		return nil
	}
	cols := make([][]int, len(m[0]))
	for j := range m[0] {
		cols[j] = make([]int, len(m))
		for i := range m {
			cols[j][i] = m[i][j]
		}
	}
	return cols
}

func (m Matrix) Rows() [][]int {
	rows := make([][]int, len(m))
	for i := range m {
		rows[i] = append([]int(nil), m[i]...)
	}
	return rows
}

func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(*m) || col < 0 || col >= len((*m)[0]) {
		return false
	}
	(*m)[row][col] = val
	return true
}
