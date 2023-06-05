package internal

import (
	"mock-project/grpc/booking-grpc/ent/booking"
	"strings"
)

var bookingStatus = map[string]booking.Status{
	"scheduled": booking.StatusScheduled,
	"completed": booking.StatusCompleted,
	"cancel":    booking.StatusCancel,
}

func ParseString(str string) (booking.Status, bool) {
	c, ok := bookingStatus[strings.ToLower(str)]
	return c, ok
}
