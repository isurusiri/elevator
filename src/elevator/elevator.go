package elevator

import (
	"math/rand"
)

var totalFloors = 21
var totalPassengerCount = 20
var noOfPassengers = 0
var currentFloor = 1

// DoWork starts the elevator
func DoWork() {
	currentFloor = getNextFloor()

}

func getNextFloor() int {
	return rand.Intn(totalFloors-1) + 1
}

func passengersIn() int {
	return rand.Intn(totalPassengerCount-1) + 1
}

func updateCurrentPassengerCount(passengersIn int) int {
	if noOfPassengers+passengersIn > totalPassengerCount {
		// display an erro since total passengers cannot exceed
		// maximum passengers
		return totalPassengerCount
	}
	return noOfPassengers + passengersIn
}
