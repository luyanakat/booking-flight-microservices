package helper

import "time"

func CheckCancelBooking(departure time.Time) bool {
	currentTime := time.Now()

	bookingHour := departure.Add(-time.Hour * 48)

	return currentTime.Before(bookingHour)
}

func CheckCreateBooking(departure time.Time) bool {
	currentTime := time.Now()

	bookingHour := departure.Add(-time.Hour * 4)

	return currentTime.Before(bookingHour)
}
