package models

import (
	"errors"
	"social-network/src/security"
)

type Password struct {
	ActualPassword string `json:"actual_password"`
	NewPassword    string `json:"new_password"`
}

func (password *Password) validate() error {
	if password.ActualPassword == "" {
		return errors.New("actual password required")
	}
	if password.NewPassword == "" {
		return errors.New("new password required")
	}
	return nil
}

func (password *Password) format() error {
	newPasswordHash, error := security.Hash(password.NewPassword)
	if error != nil {
		return error
	}
	password.NewPassword = string(newPasswordHash)

	return nil
}

func (password *Password) Prepare() error {
	if error := password.validate(); error != nil {
		return error
	}

	if error := password.format(); error != nil {
		return error
	}

	return nil
}
