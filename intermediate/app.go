package main

import "fmt"

// Intermediate go slicing ...

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	cards := []int{2, 6, 9}
	return cards
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index > len(slice) || index < 0 {
		return -1
	}
	return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index > len(slice) || index < 0 {
		return append(slice, value)
	}

	slice[index] = value
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	return append(values, slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	var newSlice []int

	if len(slice) < index || index < 0 {
		return slice
	}

	value := slice[index]
	for i := range slice {
		if slice[i] != value {
			newSlice = append(newSlice, slice[i])
		}
	}

	return newSlice
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	card1Value := ParseCard(card1)
	card2Value := ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)

	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"

	case card1Value+card2Value == 21:
		if dealerCard != "ace" && dealerCard != "figure" && dealerCard != "ten" && dealerCard != "queen" {
			return "W"
		}
		return "S"

	case card1Value+card2Value > 16 && card1Value+card2Value < 21:
		return "S"

	case card1Value+card2Value > 11 && card1Value+card2Value < 17:
		if dealerCardValue >= 7 {
			return "H"
		}
		return "S"

	case card1Value+card2Value <= 11:
		return "H"

	default:
		return ""
	}
}

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
	}
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	var battery int
	if car.batteryDrain < 10 {
		battery = car.battery - car.batteryDrain
	} else {
		battery = (car.battery * 10 / 100) - car.batteryDrain
	}
	distance := car.speed
	return Car{
		battery:      battery,
		distance:     distance,
		batteryDrain: car.batteryDrain,
		speed:        car.speed,
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	panic("Please implement the CanFinish function")
}

func main() {
	speed := 5
	batteryDrain := 2
	car := NewCar(speed, batteryDrain)
	car = Drive(car)
	fmt.Println(car)

}
