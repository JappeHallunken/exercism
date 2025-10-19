package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix struct {
	data Data
}

type Data [][]int

type Pair struct {
	row int // east - west
	col int // north - south
}

func (m *Matrix) rowMaxes() []int {
	maxes := make([]int, len(m.data))
	for i, row := range m.data {
		max := row[0]
		for _, v := range row[1:] {
			if v > max {
				max = v
			}
		}
		maxes[i] = max
	}
	return maxes
}
func (m *Matrix) colMins() []int {
	if len(m.data) == 0 {
		return nil
	}
	cols := len(m.data[0])
	mins := make([]int, cols)
	for j := range cols {
		min := m.data[0][j]
		for i := 1; i < len(m.data); i++ {
			if m.data[i][j] < min {
				min = m.data[i][j]
			}
		}
		mins[j] = min
	}
	return mins
}

func New(s string) (*Matrix, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return new(Matrix), nil
	}
	lines := strings.Split(s, "\n")

	data := make(Data, len(lines))

	for i, line := range lines {
		fields := strings.Fields(line)

		row := make([]int, len(fields))
		for j, f := range fields {
			val, err := strconv.Atoi(f)
			if err != nil {
				return nil, fmt.Errorf("invalid number %q, %w", f, err)
			}
			row[j] = val
		}
		data[i] = row
	}
	return &Matrix{data: data}, nil
}

func (m *Matrix) Saddle() []Pair {
	if len(m.data) == 0 {
		return nil
	}
	rowMax := m.rowMaxes()
	colMin := m.colMins()

	var result []Pair
	for i, row := range m.data {
		for j, val := range row {
			if val == rowMax[i] && val == colMin[j] {
				result = append(result, Pair{row: i + 1, col: j + 1})
			}
		}
	}
	return result
}
