package authuser_test

import (
	"testing"

	authuser "github.com/glyphack/koal/internal/module/auth/domain/user"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestSetValidEmail(t *testing.T) {
	user := authuser.User{Email: "testEmail@test.com"}
	newEmail := "newEmail@test.com"
	user.SetEmailAddress(newEmail)
	assert.Equal(t, user.Email, newEmail)
}

func TestSetInvalidEmail(t *testing.T) {
	user := authuser.User{Email: "testEmail@test.com"}
	newInvalidEmail := "pass"
	err := user.SetEmailAddress(newInvalidEmail)
	assert.Equal(t, err, authuser.EmailIsNotValidError)
}

func TestSetPasswordValidPassword(t *testing.T) {
	user := authuser.User{Email: "testEmail@test.com", Password: "OldPassword"}
	newPassword := "strong_password"
	_ = user.SetPassword(newPassword)
	assert.Equal(t, user.Password, newPassword)
}

func TestSetPasswordInvalidPassword(t *testing.T) {
	user := authuser.User{Email: "testEmail@test.com", Password: "OldPassword"}
	newPassword := "pass"
	err := user.SetPassword(newPassword)
	assert.Equal(t, err, authuser.PasswordIsNotValidError)
}

func TestGenerateToken(t *testing.T) {
	viper.GetViper().Set("jwt_secret", "test")
	user := authuser.User{Email: "testEmail@test.com", Password: "OldPassword"}
	_, err := user.GenerateToken()
	t.Log(viper.GetString("jwt_secret"))
	assert.Nil(t, err)
}
