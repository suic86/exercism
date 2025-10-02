package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Robot represents a robot with random name in [A-Z]{2}[0-9]{3} format.
type Robot struct {
	name string
}

// Reset sets the robot's name to a new random value.
func (r *Robot) Reset() {
	name, err := randomName()
	if err != nil {
		panic(err)
	}
	r.name = name
}

// Name return the robot's name.
func (r *Robot) Name() (string, error) {
	if len(r.name) == 0 {
		r.Reset()
	}
	return r.name, nil
}

func randomNumber() int {
	return rand.Intn(1000)
}

func randomLetter() int {
	return 65 + rand.Intn(26)
}

var names = make(map[string]bool)
var availableNames = 26 * 26 * 1000

func randomName() (string, error) {
	if availableNames == 0 {
		return "", fmt.Errorf("no more available names.")
	}

	var name string
	for {
		name = fmt.Sprintf("%c%c%03d", randomLetter(), randomLetter(), randomNumber())
		if _, exists := names[name]; !exists {
			names[name] = true
			availableNames--
			break
		}
	}

	return name, nil
}
