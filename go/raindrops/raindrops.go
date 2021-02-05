// Contains convert function for number input to raindrop string
package raindrops

import "fmt"

// Returns raindrop string based on factors of number input
func Convert(number int) string {
	var output string
	if number%3 == 0 {
		output += "Pling"
	}
	if number%5 == 0 {
		output += "Plang"
	}
	if number%7 == 0 {
		output += "Plong"
	}

	if output == "" {
		return fmt.Sprint(number)
	}
	return output
}
