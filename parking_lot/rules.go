package main

var parkingRules = map[VehicleType][]ParkingType{
	motorCycle: {motorCycleSpot, compactSpot, largeSpot},
	car:        {compactSpot, largeSpot},
	truck:      {largeSpot},
}

var parkingPricePerHour = map[VehicleType]int{
	motorCycle: 10,
	car:        20,
	truck:      50,
}
