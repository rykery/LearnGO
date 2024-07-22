package thefarm

import (
	"errors"
	"fmt"
)

// DivideFood determines the amount of food per cow
func DivideFood(calculator FodderCalculator, numCows int) (float64, error) {
	fodderAmount, err := calculator.FodderAmount(numCows)
	if err != nil {
		return 0.0, err
	}
	fattenFactor, err := calculator.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return fodderAmount * fattenFactor / float64(numCows), nil
}

// ValidateInputAndDivideFood ensures valid input and then divides the food among the cows
func ValidateInputAndDivideFood(calculator FodderCalculator, numCows int) (float64, error) {
	if numCows > 0 {
		return DivideFood(calculator, numCows)
	}
	return 0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
	numCows int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.numCows, e.message)
}

// ValidateNumberOfCows ensures the number of cows is valid
func ValidateNumberOfCows(numCows int) error {
	if numCows < -0 {
		return &InvalidCowsError{
			numCows: numCows,
			message: "there are no negative cows",
		}
	}
	if numCows == 0 {
		return &InvalidCowsError{
			numCows: numCows,
			message: "no cows don't need food",
		}
	}
	return nil
}
