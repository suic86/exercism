package resistorcolor

var colors = []string{
	"black",
	"brown",
	"red",
	"orange",
	"yellow",
	"green",
	"blue",
	"violet",
	"grey",
	"white",
}

var colorMap = make(map[string]int)

func init() {
	for i, c := range colors {
		colorMap[c] = i
	}
}

// Colors returns the list of all colors.
func Colors() []string {
	return colors
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	return colorMap[color]
}
