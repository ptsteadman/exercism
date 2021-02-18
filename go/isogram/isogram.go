// Contains function for determining is a string is an isogram
package isogram

import (
	"regexp"
	"strings"
)

// An isogram is a string with no repeated alpha characters
func IsIsogram(input string) bool {
	r := regexp.MustCompile(`[^a-z]`)
	input = r.ReplaceAllString(strings.ToLower(input), "")
	letterSet := map[byte]bool{}
	for i := 0; i < len(input); i++ {
		_, exists := letterSet[input[i]]
		if exists {
			return false
		}
		letterSet[input[i]] = true
	}
	return true
}
