package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix struct {
	data     [][]int
	rowcount int
	colcount int
}

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
	return Matrix{m, len(m), len(m[0])}, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	cols := make([][]int, m.colcount)
	for i := 0; i < m.colcount; i++ {
		c := make([]int, m.rowcount)
		for j := 0; j < m.rowcount; j++ {
			c[j] = m.data[j][i]
		}
		cols[i] = c
	}
	return cols
}

func (m Matrix) Rows() [][]int {
	nm := make([][]int, len(m.data))
	for i, v := range m.data {
		r := make([]int, len(v))
		copy(r, m.data[i])
		nm[i] = r
	}
	return nm
}

func (m Matrix) Set(row, col, val int) bool {
	if row > m.rowcount {
		return false
	}
	if col > m.colcount {
		return false
	}

	m.data[row][col] = val
	return true
}
