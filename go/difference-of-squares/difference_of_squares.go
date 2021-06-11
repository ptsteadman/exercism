package diffsquares

// Returns the sum of the first n positive integers.
// Computes in O(1) time thanks to
// telescoping the binomial expansion of the series.
func SumOfIntegerSeries(n int) int {
	return n * (n + 1) / 2
}

func SquareOfSum(n int) int {
	sumOfOneToN := SumOfIntegerSeries(n)
	return sumOfOneToN * sumOfOneToN
}

// Returns the sum of the squares of the first n positive integers.
// Computes in O(1) time thanks to
// telescoping the binomial expansion of the series
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// Return the difference betwen the square of the sum of the first n integers
// and the sum of the squares of the first n integers
// Computes in O(1) time.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
