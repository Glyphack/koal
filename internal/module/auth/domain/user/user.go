package authuser

import (
	"errors"
	"time"

	"github.com/glyphack/koal/pkg/email"
	"github.com/glyphack/koal/pkg/passwordutils"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

var PasswordIsNotValidError = errors.New("Password must at least be 7 characters long")
var EmailIsNotValidError = errors.New("Email is not valid")

type User struct {
	Email    string
	Password string // Password is always a hash of real password
}

func (u *User) SetEmailAddress(emailAddress string) error {
	if !email.IsEmailValid(emailAddress) {
		return EmailIsNotValidError
	}
	u.Email = emailAddress
	return nil
}

func (u *User) SetPassword(password string) error {
	if len(password) < 7 {
		return PasswordIsNotValidError
	}
	var err error
	u.Password, err = passwordutils.HashPassword(password)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) CheckPassword(inputPassword string) bool {
	return passwordutils.CheckPasswordHash(inputPassword, u.Password)
}

func (u *User) GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": u.Email,
		"nbf":     time.Now().UTC().Unix(),
	})

	tokenString, err := token.SignedString([]byte(viper.GetString("jwt_secret")))
	return tokenString, err
}
