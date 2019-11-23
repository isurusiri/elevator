package elevator

import (
	"fmt"
	"queue"
)

// Elevator represents the data structure of the elevator
type Elevator struct {
	elevatorID         int
	currentFloorNumber int          // start floor is 0
	direction          int          // -1 == down, 1 == up
	goalFloorNumber    map[int]bool // stores the floors where the elevetor needs to stop
}

// PickupReq represents the data structure that uses
// when requesting an elevator to a floor
type PickupReq struct {
	pickupFloor int
	direction   int // -1 == down, 1 == up
}

// ControlSystem represents all elevators and pickup requests
type ControlSystem struct {
	elevator       []*Elevator
	pickupRequests *queue.Queue
}

// NewElevator create a pointer to a new elevator
func NewElevator(ID int) *Elevator {
	e := Elevator{elevatorID: ID, direction: 1}
	e.goalFloorNumber = make(map[int]bool)
	return &e
}

// GetElevatorID reads the elevator id
func (e *Elevator) GetElevatorID() int {
	return e.elevatorID
}

// GetCurrentFloorNumber reads the current floor number of the elevator
func (e *Elevator) GetCurrentFloorNumber() int {
	return e.currentFloorNumber
}

// GetDirection reads the direction the elevator is headed
func (e *Elevator) GetDirection() int {
	return e.direction
}

// GetNumGoalFloors reads the number of floors that elevator has to stop
func (e *Elevator) GetNumGoalFloors() int {
	n := len(e.goalFloorNumber)
	return n
}

// GetGoalFloorNumbers reads all floors that elevator has to stop
func (e *Elevator) GetGoalFloorNumbers() []int {
	goalFloors := make([]int, 0)

	for k := range e.goalFloorNumber {
		goalFloors = append(goalFloors, k)
	}

	return goalFloors
}

func (e *Elevator) addGoalFloor(floorNumber int) {
	e.goalFloorNumber[floorNumber] = true
}

func (e *Elevator) removeGloalFloor(floorNumber int) {
	delete(e.goalFloorNumber, floorNumber)
}

func (e *Elevator) canMove() bool {
	if e.GetNumGoalFloors() > 0 {
		return true
	}

	return false
}

func (e *Elevator) canAddGoalFloor(goalFloorNumber int, direction int) bool {
	// if there are no goalFloors
	if e.GetNumGoalFloors() == 0 {
		e.direction = direction
		return true
		// if the move direction of the elevator is the same was requested
	} else if e.direction == direction {
		// if move up
		if direction > 0 && e.currentFloorNumber <= goalFloorNumber {
			return true
			// if move down
		} else if direction < 0 && e.currentFloorNumber >= goalFloorNumber {
			return true
		}
	}

	return false
}

// Update moves the elevator
func (e *Elevator) Update(currentFloorNum int, goalFloorNum int, direction int) bool {
	// update the goalFloorNumber map
	if e.canMove() {
		e.currentFloorNumber = currentFloorNum
		e.removeGloalFloor(e.currentFloorNumber)
	}

	// if elevator moves in the same direction or the elevator is empty
	if e.canAddGoalFloor(goalFloorNum, direction) {
		e.addGoalFloor(goalFloorNum)
		return true
	}

	return false
}

// GetNextFloor returns the next floor that elevator should go
func (e *Elevator) GetNextFloor() int {
	// move down
	if e.direction == -1 && e.currentFloorNumber > 0 {
		return e.currentFloorNumber - 1
	}
	// move up
	return e.currentFloorNumber + 1
}

// NewElevatorControlSystem creates new control system
func NewElevatorControlSystem(NumberOfElevators int) *ControlSystem {
	cs := ControlSystem{}

	for i := 0; i < NumberOfElevators; i++ {
		cs.elevator = append(cs.elevator, NewElevator(i))
	}
	cs.pickupRequests = queue.NewQueue()

	return &cs
}

// Status prints where all elevators are located at and their work load
func (cs *ControlSystem) Status() []string {
	seq := make([]string, 0)

	for _, elev := range cs.elevator {
		seq = append(seq, fmt.Sprintf("elevatorID: %v, currentFloor: %v, direction: %v, goalFloors: %v", elev.GetElevatorID(), elev.GetCurrentFloorNumber(), elev.GetDirection(), elev.GetGoalFloorNumbers()))
	}

	return seq
}

// Pickup new pickup request
func (cs *ControlSystem) Pickup(pickupFloorNumber int, direction int) {
	cs.pickupRequests.Push(PickupReq{pickupFloorNumber, direction})
}

func (cs *ControlSystem) update(elev *Elevator, currentFloor int, goalFloor int, direction int) bool {
	return elev.Update(currentFloor, goalFloor, direction)
}

// Step evaluate the next step of the elevator
func (cs *ControlSystem) Step() {
	for _, elev := range cs.elevator {
		// check if there is a pickup request
		if cs.pickupRequests.Len() > 0 {
			req := cs.pickupRequests.Peek()

			if e, ok := req.(PickupReq); ok {
				// if the elevator moves in the same direction as the new pickup request
				success := cs.update(elev, elev.GetNextFloor(), e.pickupFloor, e.direction)
				if success {
					_ = cs.pickupRequests.Pop()
				}
			}
			// check if the goalFloorNumber map is not empty
			// if the goalFloorNumber map is not empty, the elevator needs to move to the next goal floor
		} else if elev.GetNumGoalFloors() > 0 {
			_ = cs.update(elev, elev.GetNextFloor(), -1, elev.GetDirection())
		}
	}
}
