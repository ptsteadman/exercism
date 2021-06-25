// Package clock implements a 24 hour wall clock times that handles overflows
package clock

import "fmt"

// Clock has hour and minute int members
type Clock struct {
	hours   int
	minutes int
}

// New takes hours and minutes and ints and returns a normalized wall clock time
func New(hours int, minutes int) Clock {
	return Clock{hours: 0, minutes: 0}.Add(hours*60 + minutes)
}

// String outputs the 24 hour clock time in a format like 00:00
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}

func (c Clock) addHour() Clock {
	c.hours += 1
	if c.hours > 23 {
		c.hours -= 24
	}
	return c
}

// Add adds the given number minutes and returns a new Clock
func (c Clock) Add(minutes int) Clock {
	if minutes < 0 {
		return c.Subtract(-minutes)
	}

	for minutes > 59 {
		minutes -= 60
		c = c.addHour()
	}
	c.minutes += minutes
	if c.minutes > 59 {
		c.minutes -= 60
		c = c.addHour()
	}
	return c
}

func (c Clock) subHour() Clock {
	c.hours -= 1
	if c.hours < 0 {
		c.hours += 24
	}
	return c
}

// Subtract subtracts the given number minutes and returns a new Clock
func (c Clock) Subtract(minutes int) Clock {
	if minutes < 0 {
		return c.Add(-minutes)
	}

	for minutes > 59 {
		minutes -= 60
		c = c.subHour()
	}
	c.minutes -= minutes
	if c.minutes < 0 {
		c.minutes += 60
		c = c.subHour()
	}
	return c
}
