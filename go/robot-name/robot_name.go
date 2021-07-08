package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

// UppercaseChars are all the possible characters in the first part of a robot name
const UppercaseChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// NumNames is total number of possible robot names
const NumNames = 26 * 26 * 10 * 10 * 10

// GenerateNames generates all possible robot names
func GenerateNames() []string {
	names := make([]string, NumNames)

	i := 0
	for _, c1 := range UppercaseChars {
		for _, c2 := range UppercaseChars {
			for d1 := 0; d1 < 10; d1++ {
				for d2 := 0; d2 < 10; d2++ {
					for d3 := 0; d3 < 10; d3++ {
						names[i] = fmt.Sprintf("%s%s%v%v%v", string(c1), string(c2), d1, d2, d3)
						i++
					}
				}
			}
		}
	}
	return names
}

// UnusedNames is a list of robot names that have never been assigned to a Robot
var UnusedNames = GenerateNames()

// Robot has a name in the format of two uppercase letters followed by three digits
type Robot struct {
	name string
}

// Name returns a Robot's name, setting it if is not yet set
// If all Robot names are exhausted, an error is thrown
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	if len(UnusedNames) == 0 {
		return "", errors.New("all valid robot names are taken")
	}
	nameIndex := rand.Intn(len(UnusedNames))
	r.name = UnusedNames[nameIndex]
	// remove the name from the set of possible names
	UnusedNames[nameIndex] = UnusedNames[len(UnusedNames)-1]
	UnusedNames = UnusedNames[:len(UnusedNames)-1]

	return r.name, nil
}

// Reset unsets the Robot's name. It is not returned to the pool of UnusedNames
func (r *Robot) Reset() {
	r.name = ""
}
