package clock

import "fmt"

// Clock represents a 00:00-format 24h clock.
type Clock struct {
	Minute int
}

// New creates a new Clock.
func New(hour, minute int) Clock {
	minute = hour*60 + minute
	minute %= 24 * 60
	if minute < 0 {
		minute += 24 * 60
	}
	return Clock{Minute: minute}
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.Minute/60, clock.Minute%60)
}

// Add minutes to a clock.
func (clock Clock) Add(minutes int) Clock {
	return New(0, clock.Minute+minutes)
}

// Subtract minutes from a clock.
func (clock Clock) Subtract(minutes int) Clock {
	return New(0, clock.Minute-minutes)
}
