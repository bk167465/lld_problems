package main

import (
	"fmt"
	"time"
)

func main() {
	floor1 := ParkingFloor{
		floorNumber: 1,
		spots:       map[ParkingType][]ParkingSpot{},
	}

	motorcycle := []ParkingSpot{
		{floorNumber: 1, parkingType: motorCycleSpot, spotNumber: 1},
		{floorNumber: 1, parkingType: motorCycleSpot, spotNumber: 2},
	}

	compact := []ParkingSpot{
		{floorNumber: 1, parkingType: compactSpot, spotNumber: 1},
		{floorNumber: 1, parkingType: compactSpot, spotNumber: 2},
	}

	large := []ParkingSpot{
		{floorNumber: 1, parkingType: largeSpot, spotNumber: 1},
	}

	floor1.spots[motorCycleSpot] = motorcycle
	floor1.spots[compactSpot] = compact
	floor1.spots[largeSpot] = large

	parkingLot := ParkingLot{
		floors:      []ParkingFloor{floor1},
		presenceMap: make(map[string]ParkingTicket),
	}

	car1 := Vehicle{
		vehicleType:   car,
		vehicleNumber: "DL01AA1111",
	}

	ticket, err := parkingLot.ParkVehicle(car1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Vehicle Parked")
	fmt.Printf("Ticket: %+v\n", *ticket)

	// Simulate 2.5 hours of parking
	t := parkingLot.presenceMap[car1.vehicleNumber]
	t.parkedAt = time.Now().Add(-2*time.Hour - 30*time.Minute)
	parkingLot.presenceMap[car1.vehicleNumber] = t

	amount, err := parkingLot.ExitVehicle(car1.vehicleNumber)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Parking Fee = ₹%d\n", amount)
}
