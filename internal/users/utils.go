package users

import (
	"errors"
)

func BodyCreateValidation(body CreateRequest) (int, error) {
	// Validations
	if body.FirstName == "" {
		return 400, errors.New("first_name is required")
	}
	if body.LastName == "" {
		return 400, errors.New("last_name is required")
	}
	if body.Email == "" {
		return 400, errors.New("email is required")
	}
	if body.Phone == "" {
		return 400, errors.New("email is required")
	}

	return 201, nil
}
