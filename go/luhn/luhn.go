package luhn

import (
	"strconv"
	"strings"
)

// Valid returns true if the input is valid per the Luhn formula.
func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")
	// Must at least two digits in length
	if len(input) < 2 {
		return false
	}

	luhnSum := 0
	// Double every second digit, starting from the right
	checksumDigit := len(input)%2 == 0
	for _, r := range input {
		toAdd, err := strconv.Atoi(string(r))
		// If any of the input characters are not digits, not a valid luhn code
		if err != nil {
			return false
		}
		if checksumDigit {
			toAdd *= 2
			if toAdd > 9 {
				toAdd -= 9
			}
		}
		luhnSum += toAdd
		checksumDigit = !checksumDigit
	}

	return luhnSum%10 == 0
}
