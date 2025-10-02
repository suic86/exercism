package spiralmatrix

func SpiralMatrix(size int) [][]int {
	if size == 0 {
		return [][]int{}
	}
	if size == 1 {
		return [][]int{{1}}
	}

	m := make([][]int, size)
	for i := range m {
		m[i] = make([]int, size)
	}

	n, x, y := 1, 0, 0
	dx, dy := 0, 1
	for n <= size*size {
		if 0 <= x && x < size && 0 <= y && y < size {
			m[x][y] = n
		}
		if x+dx < 0 || x+dx == size || y+dy < 0 || y+dy == size || m[x+dx][y+dy] != 0 {
			dx, dy = dy, -dx
		}
		x += dx
		y += dy
		n += 1
	}

	return m
}
