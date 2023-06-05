package internal

import (
	"mock-project/grpc/flight-grpc/ent"
)

func CheckFlightEmptyInput(flightInput, flight *ent.Flight) *ent.Flight {
	if flightInput.Name == "" {
		flightInput.Name = flight.Name
	}
	if flightInput.From == "" {
		flightInput.From = flight.From
	}
	if flightInput.To == "" {
		flightInput.To = flight.To
	}
	if flightInput.AvailableEconomySlot == 0 {
		flightInput.AvailableEconomySlot = flight.AvailableEconomySlot
	}
	if flightInput.AvailableFirstSlot == 0 {
		flightInput.AvailableFirstSlot = flight.AvailableFirstSlot
	}
	if flightInput.Status == "" {
		flightInput.Status = flight.Status
	}
	return flightInput
}
