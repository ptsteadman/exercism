package grains

import (
	"errors"
	"math/big"
)

// Square returns the number of grains on a given chessboard space where the number
// of grains double on each successive place. The first chessboard space is given index 1
func Square(input int) (uint64, error) {
	if input < 1 {
		return 0, errors.New("Input must be greater than 1.")
	}
	if input > 64 {
		return 0, errors.New("Input must be less than 64.")
	}
	var result big.Int
	result.Exp(big.NewInt(2), big.NewInt(int64(input-1)), nil)

	return result.Uint64(), nil
}

// Total returns the sum of all grains on the chessboard (all 64 squares).
// This can be calculated in closed form of 2^64 - 1, thanks to the following relationship:
// total(n) = 2^0 + 2^1 + 2^2 + ... 2^n-1
// 2 * total(n) = 2^1 + 2^2 + ... + 2^n-1 + 2^n
// 2 * total(n) - total(n) = total(n) = 2^n - 2^0 = 2^n - 1
func Total() uint64 {
	var result big.Int
	result.Exp(big.NewInt(2), big.NewInt(64), nil)
	return result.Uint64() - 1
}
