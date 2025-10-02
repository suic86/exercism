package clock

import "fmt"

// Clock represents a 00:00-format 24h clock.
type Clock struct {
	Minute int
}

// New creates a new Clock.
func New(hour, minute int) Clock {
	minute = (1440 + (minute+60*hour)%1440) % 1440
	return Clock{Minute: minute}
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.Minute/60, clock.Minute%60)
}

// Add minutes to a clock.
func (clock Clock) Add(minutes int) Clock {
	return New(clock.Minute/60, clock.Minute%60+minutes)
}

// Subtract minutes from a clock.
func (clock Clock) Subtract(minutes int) Clock {
	return New(clock.Minute/60, clock.Minute%60-minutes)
}
