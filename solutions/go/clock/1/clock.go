package clock

import "fmt"

// Clock represents a 00:00-format 24h clock.
type Clock struct {
	Hour   int
	Minute int
}

// New creates a new Clock.
func New(hour, minute int) Clock {
	dh := minute / 60

	if minute < 0 && minute%60 != 0 {
		dh--
	}

	hour += dh
	hour %= 24

	if hour < 0 {
		hour += 24
	}

	if hour == 24 {
		hour = 0
	}

	minute %= 60
	if minute < 0 {
		minute += 60
	}

	return Clock{Hour: hour, Minute: minute}
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
