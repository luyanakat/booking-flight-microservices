package helper

import "time"

func ConvertTime(input string) (time.Time, error) {
	date, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}
