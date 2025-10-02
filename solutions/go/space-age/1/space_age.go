package space

const (
	earthYearInSeconds float64 = 31557600.0
)

type Planet string

func Age(seconds float64, planet Planet) float64 {
	var coef float64 = 1
	switch planet {
	case "Earth":
		coef = 1
	case "Mercury":
		coef = 0.2408467
	case "Venus":
		coef = 0.61519726
	case "Mars":
		coef = 1.8808158
	case "Jupiter":
		coef = 11.862615
	case "Saturn":
		coef = 29.447498
	case "Uranus":
		coef = 84.016846
	case "Neptune":
		coef = 164.79132
	}
	return seconds / (coef * earthYearInSeconds)
}
