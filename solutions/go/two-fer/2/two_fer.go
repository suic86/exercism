// Package twofer provides a solution for exercism.io go track "Two Fer" exercise.
package twofer

import "fmt"

// ShareWith returns "One for you, and for me." when the name is empty and "One for <name>, and for me." otherwise."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
