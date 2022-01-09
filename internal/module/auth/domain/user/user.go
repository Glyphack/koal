package authuser

import (
	"errors"

	"github.com/glyphack/koal/pkg/email"
)

type User struct {
	Email    string
	Password string
}

func (u *User) SetEmailAddress(emailAddress string) error {
	if !email.IsEmailValid(emailAddress) {
		return errors.New("Email is not valid")
	}
	u.Email = emailAddress
	return nil
}

func (u *User) SetPassword(password string) error {
	if len(password) < 7 {
		return errors.New("Password must at least be 7 characters long")
	}

	u.Password = password
	return nil
}
