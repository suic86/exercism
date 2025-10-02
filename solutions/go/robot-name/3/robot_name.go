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
func (r *Robot) Reset() error {
	name, err := randomName()
	r.name = name
	return err
}

// Name returns the robot's name.
func (r *Robot) Name() (string, error) {
	if len(r.name) != 0 {
		return r.name, nil
	}
	err := r.Reset()
	return r.name, err
}

var used = map[string]bool{}
var availableNames = 26 * 26 * 1000

func randomNumber() int {
	return rand.Intn(1000)
}

func randomLetter() int {
	return 65 + rand.Intn(26)
}

func newName() string {
	return fmt.Sprintf("%c%c%03d", randomLetter(), randomLetter(), randomNumber())
}

func randomName() (string, error) {
	if availableNames == 0 {
		return "", fmt.Errorf("no more available names")
	}

	var name string
	for {
		name = newName()
		if !used[name] {
			used[name] = true
			availableNames--
			break
		}
	}

	return name, nil
}
