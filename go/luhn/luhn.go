package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

// Returns true if the input is valid per the Luhn formula.
func Valid(input string) bool {
	inputSpacesStripped := strings.ReplaceAll(input, " ", "")
	// Must be a string of all digits, at least two in length
	r, _ := regexp.Compile("^\\d{2,}$")
	if ok := r.MatchString(inputSpacesStripped); ok == false {
		return false
	}

	luhnSum := 0
	lastIndexOfInput := len(inputSpacesStripped) - 1
	for i := lastIndexOfInput; i >= 0; i-- {
		toAdd, _ := strconv.Atoi(string(inputSpacesStripped[i]))
		if (lastIndexOfInput-i)%2 != 0 {
			// double every second digit, starting from the right
			toAdd = toAdd * 2
			if toAdd > 9 {
				toAdd = toAdd - 9
			}
		}
		luhnSum += toAdd
	}

	return luhnSum%10 == 0
}
