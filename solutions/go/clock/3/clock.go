package clock

import "fmt"

// Clock represents a 00:00-format 24h clock.
type Clock struct {
	Hour   int
	Minute int
}

// New creates a new Clock.
func New(hour, minute int) Clock {
	minute = (1440 + (minute+60*hour)%1440) % 1440
	return Clock{Hour: (minute / 60) % 24, Minute: minute % 60}
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.Hour, clock.Minute)
}

// Add minutes to a clock.
func (clock Clock) Add(minutes int) Clock {
	return New(clock.Hour, clock.Minute+minutes)
}

// Subtract minutes from a clock.
func (clock Clock) Subtract(minutes int) Clock {
	return clock.Add(-minutes)
}
