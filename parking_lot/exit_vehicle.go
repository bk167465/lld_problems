package main

import (
	"errors"
	"math"
	"time"
)

func (p *ParkingLot) ExitVehicle(vehicleNumber string) (int, error) {
	ticket, ok := p.presenceMap[vehicleNumber]
	if !ok {
		return 0, errors.New("Vehicle is not parked with us")
	}

	for i := range p.floors {
		if p.floors[i].floorNumber == ticket.parkingSpot.floorNumber {
			p.floors[i].freeSpot(ticket.parkingSpot)
			break
		}
	}

	amount := calculateFee(ticket)

	delete(p.presenceMap, vehicleNumber)

	return amount, nil
}

func calculateFee(ticket ParkingTicket) int {
	duration := time.Since(ticket.parkedAt)
	hours := duration.Hours()

	ceilHours := int(math.Ceil(hours))

	return ceilHours * parkingPricePerHour[ticket.vehicle.vehicleType]
}

func (pf *ParkingFloor) freeSpot(spot ParkingSpot) {
	pf.spots[spot.parkingType] = append(pf.spots[spot.parkingType], spot)
}
