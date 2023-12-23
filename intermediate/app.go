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

func main() {
	card := GetItem([]int{1, 2, 4, 1}, 40)
	fmt.Println(card)
}
