package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowError struct {
	cows    int
	message string
}

func (ice *InvalidCowError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", ice.cows, ice.message)
}

func DivideFood(fc FodderCalculator, cows int) (float64, error) {
	fa, err := fc.FodderAmount(cows)
	if err != nil {
		return 0.0, err
	}
	ff, err := fc.FatteningFactor()
	if err != nil {
		return 0.0, err
	}
	return (fa * ff) / float64(cows), nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0.0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, cows)
}

func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowError{cows, "there are no negative cows"}
	}
	if cows == 0 {
		return &InvalidCowError{cows, "no cows don't need food"}
	}
	return nil
}
