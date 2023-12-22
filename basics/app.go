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

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 {
		return option1 + " is clearly the better choice."
	}

	return option2 + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
// si orginalPrice -> 100%
//
//	x            -> 80%
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		cost := (originalPrice * 80) / 100
		return cost
	}
	if age > 3 && age < 10 {
		cost := (originalPrice * 70) / 100
		return cost
	}
	cost := (originalPrice * 50) / 100
	return cost

}

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	var helper string
	if table < 10 {
		helper = "00"
	} else if table < 100 {
		helper = "0"
	}
	return fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table %s%d. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.", name, helper, table, direction, distance, neighbor)
}

func main() {
	fmt.Println(AssignTable("Christiane", 2, "Frank", "on the left", 23.7834298))
}
