package main

import "fmt"

// Define the Clock type here.
type Clock struct {
	time int
}

func New(h, m int) Clock {
	total := m + h*60
	total %= 24 * 60
	if total < 0 {
		total += 24 * 60
	}
	return Clock{total}
}

// Add - add minutes to a clock
func (c Clock) Add(minutes int) Clock {
	return New(0, c.time+minutes)
}

// Subtract - subtract minutes from a clock
func (c Clock) Subtract(minutes int) Clock {
	return New(0, c.time-minutes)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.time/60, c.time%60)
}
