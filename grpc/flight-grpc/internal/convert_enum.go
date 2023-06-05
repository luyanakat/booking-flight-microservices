package internal

import (
	"mock-project/grpc/flight-grpc/ent/flight"
	"strings"
)

var flightStatus = map[string]flight.Status{
	"available": flight.StatusAvailable,
	"arrived":   flight.StatusArrived,
	"cancel":    flight.StatusCancel,
}

func ParseString(str string) (flight.Status, bool) {
	c, ok := flightStatus[strings.ToLower(str)]
	return c, ok
}
