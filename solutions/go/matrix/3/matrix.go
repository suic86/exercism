package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(s string) (Matrix, error) {
	m := [][]int{}
	rowlen := -1
	for _, row := range strings.Split(s, "\n") {
		fs := strings.Split(strings.Trim(row, " "), " ")
		if rowlen != -1 {
			if len(fs) != rowlen {
				return Matrix{}, errors.New("invalid matrix: differnt row lengths")
			}
		}
		rowlen = len(fs)

		r := []int{}
		for _, col := range fs {
			v, err := strconv.Atoi(col)
			if err != nil {
				return Matrix{}, err
			}
			r = append(r, v)
		}
		m = append(m, r)
	}
	return m, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	colcount := len(m[0])
	rowcount := len(m)
	cols := make([][]int, colcount)
	for i := 0; i < colcount; i++ {
		c := make([]int, rowcount)
		for j := 0; j < rowcount; j++ {
			c[j] = m[j][i]
		}
		cols[i] = c
	}
	return cols
}

func (m Matrix) Rows() [][]int {
	nm := make([][]int, len(m))
	for i, v := range m {
		r := make([]int, len(v))
		copy(r, m[i])
		nm[i] = r
	}
	return nm
}

func (m Matrix) Set(row, col, val int) bool {
	if !(0 <= row && row < len(m)) {
		return false
	}
	if !(0 <= col && col < len(m[0])) {
		return false
	}

	m[row][col] = val
	return true
}
