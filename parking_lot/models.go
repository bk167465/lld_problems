package main

import "time"

type VehicleType int

const (
	motorCycle VehicleType = iota
	car
	truck
)

type ParkingType int

const (
	motorCycleSpot ParkingType = iota
	compactSpot
	largeSpot
)

type Vehicle struct {
	vehicleType   VehicleType
	vehicleNumber string
}

type ParkingLot struct {
	floors      []ParkingFloor
	presenceMap map[string]ParkingTicket
}

type ParkingFloor struct {
	floorNumber int
	spots       map[ParkingType][]ParkingSpot
}

type ParkingSpot struct {
	floorNumber int
	parkingType ParkingType
	spotNumber  int
}

type ParkingTicket struct {
	ticketID    string
	vehicle     Vehicle
	parkedAt    time.Time
	parkingSpot ParkingSpot
}
