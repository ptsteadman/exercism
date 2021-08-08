package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix struct {
	values [][]int
}

func New(s string) (*Matrix, error) {
	m := Matrix{}
	rows := strings.Split(s, "\n")
	m.values = make([][]int, len(rows))
	var numCols int
	for i, row := range rows {
		stringRowValues := strings.Split(strings.TrimSpace(row), " ")
		if i == 0 {
			numCols = len(stringRowValues)
		}
		if len(stringRowValues) != numCols {
			return nil, errors.New("number of columns must be the same in each row")
		}
		rowValues := make([]int, numCols)
		for c, v := range stringRowValues {
			intVal, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			rowValues[c] = intVal
		}
		m.values[i] = rowValues
	}
	return &m, nil
}

func (m *Matrix) Rows() [][]int {
	rowsCopy := make([][]int, len(m.values))
	for i := range m.values {
		rowsCopy[i] = make([]int, len(m.values[i]))
		copy(rowsCopy[i], m.values[i])
	}
	return rowsCopy
}

func (m *Matrix) Cols() [][]int {
	if len(m.values) == 0 {
		return make([][]int, 0)
	}
	numCols := len(m.values[0])
	colsCopy := make([][]int, numCols)
	for i := 0; i < numCols; i++ {
		colsCopy[i] = make([]int, len(m.values))
		for j, row := range m.values {
			colsCopy[i][j] = row[i]
		}
	}
	return colsCopy
}

func (m *Matrix) Set(r int, c int, v int) bool {
	if r < 0 || r >= len(m.values) {
		return false
	}
	if c < 0 || c >= len(m.values[0]) {
		return false
	}
	m.values[r][c] = v
	return true
}
