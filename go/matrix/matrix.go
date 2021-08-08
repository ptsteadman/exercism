package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix is a m*n two dimensional slice of ints with Rows, Cols and Set methods
type Matrix [][]int

// New returns a new Matrix
func New(s string) (Matrix, error) {
	m := Matrix{}
	rows := strings.Split(s, "\n")
	m = make([][]int, len(rows))
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
		m[i] = rowValues
	}
	return m, nil
}

// Rows returns a two dimensional row-wise copy of Matrix values
func (m Matrix) Rows() [][]int {
	rowsCopy := make([][]int, len(m))
	for i := range m {
		rowsCopy[i] = make([]int, len(m[i]))
		copy(rowsCopy[i], m[i])
	}
	return rowsCopy
}

// Cols returns a two dimensional column-wise copy of Matrix values
func (m Matrix) Cols() [][]int {
	if len(m) == 0 {
		return make([][]int, 0)
	}
	numCols := len(m[0])
	colsCopy := make([][]int, numCols)
	for i := 0; i < numCols; i++ {
		colsCopy[i] = make([]int, len(m))
		for j, row := range m {
			colsCopy[i][j] = row[i]
		}
	}
	return colsCopy
}

// Set modifies a Matrix value at the specified row r, columni c, to a new value v
// Returns false if rows and columns were out of range, otherwise returns true
func (m Matrix) Set(r int, c int, v int) bool {
	if r < 0 || r >= len(m) {
		return false
	}
	if c < 0 || c >= len(m[0]) {
		return false
	}
	m[r][c] = v
	return true
}
