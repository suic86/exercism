// Package weather contains weather report related functions.
package weather

// CurrentCondition is the current wheater condition.
var CurrentCondition string

// CurrentLocation is the current location.
var CurrentLocation string

// Forecast return a string displaying the current location an weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
