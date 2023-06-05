package helper

import (
	"errors"
	"mock-project/graphql/graph/model"
	"net/mail"
)

func CheckBookingInput(input *model.CreateBookingInput) error {
	if input.Email == nil || input.IdentifyNumber == nil || input.Address == nil || input.DateOfBirth == nil || input.PhoneNumber == nil {
		return errors.New("not enough required field")
	}
	if CheckEmailValid(*input.Email) == false {
		return errors.New("email not valid")
	}
	return nil
}

func CheckCustomerInput(input *model.UpdateCustomerInput) error {
	if *input.Email == "" || *input.IdentifyNumber == "" || *input.Address == "" || *input.DateOfBirth == "" || *input.PhoneNumber == "" || *input.Name == "" {
		return errors.New("not enough required field")
	}
	if CheckEmailValid(*input.Email) == false {
		return errors.New("email not valid")
	}
	return nil
}

func CheckLoginInput(email, password string) error {
	if !CheckEmailValid(email) {
		return errors.New("email not valid")
	}
	if !CheckPasswordLength(password) {
		return errors.New("password length must be >= 8")
	}
	return nil
}

func CheckEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func CheckPasswordLength(password string) bool {
	if len(password) >= 8 {
		return true
	}
	return false
}
