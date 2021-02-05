package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("Number must be greater than 0.")
	}
	var steps int
	for n != 1 {
		steps++
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
	}
	return steps, nil
}
