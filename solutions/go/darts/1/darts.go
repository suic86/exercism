package darts

import (
	"math"
)

func Score(x, y float64) int {
	r := math.Hypot(x, y)
	circles := [4]float64{10, 5, 1, 0}
	points := [4]int{0, 1, 5, 10}
	for i, c := range circles {
		if r > c {
			return points[i]
		}
	}
	return points[len(points)-1]
}
