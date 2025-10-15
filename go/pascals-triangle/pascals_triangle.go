package pascal

func Triangle(n int) [][]int {
	tri := make([][]int, n)
	for line := range n {
		tri[line] = make([]int, line+1)
		tri[line][0], tri[line][line] = 1, 1
		for idx := 1; idx < line; idx++ {
			tri[line][idx] = tri[line-1][idx-1] + tri[line-1][idx]
		}
	}
	return tri
}
