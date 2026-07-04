package main

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

func (p *ParkingLot) ParkVehicle(vehicle Vehicle) (*ParkingTicket, error) {
	if _, exists := p.presenceMap[vehicle.vehicleNumber]; exists {
		return nil, errors.New("Sorry, vehicle already parked")
	}

	allowedParkingTypes := parkingRules[vehicle.vehicleType]

	for _, floor := range p.floors {
		spot, ok := floor.getAvailableSpot(allowedParkingTypes)
		if !ok {
			continue
		}
		ticket := p.createTicket(vehicle, spot)
		p.presenceMap[vehicle.vehicleNumber] = *ticket
		return ticket, nil
	}

	return nil, errors.New("Sorry, no slot avialble")
}

func (pf *ParkingFloor) getAvailableSpot(allowedTypes []ParkingType) (ParkingSpot, bool) {
	for _, v := range allowedTypes {
		if len(pf.spots[v]) > 0 {
			spot := pf.spots[v][0]
			pf.spots[v] = pf.spots[v][1:]
			return spot, true
		}
	}
	return ParkingSpot{}, false
}

func (p *ParkingLot) createTicket(vehicle Vehicle, spot ParkingSpot) *ParkingTicket {
	return &ParkingTicket{
		ticketID:    uuid.NewString(),
		vehicle:     vehicle,
		parkedAt:    time.Now(),
		parkingSpot: spot,
	}
}
