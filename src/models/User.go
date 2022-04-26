package models

import (
	"errors"
	"social-network/src/security"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (user *User) validate(step string) error {
	if step != "update-password" && user.Name == "" {
		return errors.New("name required")
	}
	if step != "update-password" && user.Nick == "" {
		return errors.New("nick required")
	}
	if step != "update-password" && user.Email == "" {
		return errors.New("e-mail required")
	}
	if step != "update-password" {
		if error := checkmail.ValidateFormat(user.Email); error != nil {
			return errors.New("e-mail invalid")
		}
	}
	if (step == "create" || step == "update-password") && user.Password == "" {
		return errors.New("password required")
	}
	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
	if step == "create" || step == "update-password" {
		passwordHash, error := security.Hash(user.Password)
		if error != nil {
			return error
		}
		user.Password = string(passwordHash)
	}
	return nil
}

func (user *User) Prepare(step string) error {
	if error := user.validate(step); error != nil {
		return error
	}

	if error := user.format(step); error != nil {
		return error
	}

	return nil
}
