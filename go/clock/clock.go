// Package clock implements a 24 hour wall clock times that handles overflows
package clock

import "fmt"

// Clock has hour and minute int members
type Clock struct {
	hours   int
	minutes int
}

// mod is a modulus function where the result has the same sign as the divisor
func mod(a, b int) int {
	return (a%b + b) % b
}

// roundToNegativeInfinityDivision implements integer division where results are rounded towards negative
// infinity in the case of negative operands
func roundToNegativeInfinityDivision(a, b int) int {
	if a < 0 {
		return (a - (b - 1)) / b
	}
	return a / b
}

// New takes hours and minutes and ints and returns a normalized wall clock time
func New(hours int, minutes int) Clock {
	c := Clock{}
	hoursFromMinuteOverflow := roundToNegativeInfinityDivision(minutes, 60)
	c.hours = hours + hoursFromMinuteOverflow
	c.hours = mod(c.hours, 24)
	c.minutes = mod(minutes, 60)
	return c
}

// String outputs the 24 hour clock time in a format like 00:00
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}

// Add adds the given number of minutes and returns a new Clock
func (c Clock) Add(minutes int) Clock {
	return New(c.hours, c.minutes+minutes)
}

// Subtract subtracts the given number of minutes and returns a new Clock
func (c Clock) Subtract(minutes int) Clock {
	return New(c.hours, c.minutes-minutes)
}
