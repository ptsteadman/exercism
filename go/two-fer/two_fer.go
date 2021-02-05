// Defines ShareWith function
package twofer

import "fmt"

// Returns string statement of who is getting one
func ShareWith(name string) string {
	youOrName := "you"
	if name != "" {
		youOrName = name
	}
	return fmt.Sprintf("One for %s, one for me.", youOrName)
}
