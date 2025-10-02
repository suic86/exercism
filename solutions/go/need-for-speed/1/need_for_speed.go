package speed

type Car struct {
	battery int
	batteryDrain int
    speed int
	distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	c := Car{}
	c.battery = 100
	c.batteryDrain = batteryDrain
	c.speed = speed
	c.distance = 0
	return c
}

type Track struct {
    distance int
}


// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{distance}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if (car.battery - car.batteryDrain < 0) {
		return car
	}
	car.battery -= car.batteryDrain
	car.distance += car.speed
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	return car.speed * (car.battery / car.batteryDrain) >= track.distance
}
