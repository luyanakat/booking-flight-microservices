package helper

import "time"

func GetIntPointer(value int) *int {
	return &value
}
func GetTimePointer(value time.Time) *time.Time {
	return &value
}
