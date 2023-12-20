package main

import (
	"fmt"
	"strings"
)

// CaluclteWorkingCarsPerHour take two args and return a float
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate)) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	return (uint(carsCount)/10)*95000 + (uint(carsCount)%10)*10000
}

// CanFastAttack can be executed only when the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
	if knightIsAwake {
		return false
	}
	return true
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	if knightIsAwake || archerIsAwake || prisonerIsAwake {
		return true
	}
	return false
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	if !archerIsAwake && prisonerIsAwake {
		return true
	}
	return false
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	if prisonerIsAwake && !archerIsAwake && !knightIsAwake || petDogIsPresent && !archerIsAwake {
		return true
	}

	return false
}

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	customerToUpper := strings.ToUpper(customer)
	return "Welcome to the Tech Palace, " + customerToUpper
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	repeatStars := strings.Repeat("*", numStarsPerLine)
	return repeatStars + "\n" + welcomeMsg + "\n" + repeatStars

}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	removeAsterix := strings.Replace(oldMsg, "*", " ", 100)
	return strings.TrimSpace(removeAsterix)
}

func main() {

	message := `
**************************
*    BUY NOW, SAVE 10%   *
**************************
`

	fmt.Println(CleanupMessage(message))
}
