package meteorology

import (
	"fmt"
	"log"
)

type TemperatureUnit int

func (tu TemperatureUnit) String() string {
	s := ""
	switch tu {
	case Celsius:
		s = "°C"
	case Fahrenheit:
		s = "°F"
	default:
		log.Fatal("invalid temperature unit.")
	}
	return s
}

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (t Temperature) String() string {
	return fmt.Sprintf("%d %s", t.degree, t.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

func (su SpeedUnit) String() string {
	s := ""
	switch su {
	case 0:
		s = "km/h"
	case 1:
		s = "mph"
	default:
		log.Fatal("invalid speed unit")
	}
	return s
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (s Speed) String() string {
	return fmt.Sprintf("%d %s", s.magnitude, s.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (md MeteorologyData) String() string {
	return fmt.Sprintf("%s: %d %s, Wind %s at %d %s, %d%% Humidity", md.location, md.temperature.degree, md.temperature.unit, md.windDirection, md.windSpeed.magnitude, md.windSpeed.unit, md.humidity)
}
