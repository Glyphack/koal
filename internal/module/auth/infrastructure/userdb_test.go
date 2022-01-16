package authinfra_test

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"

	"github.com/glyphack/koal/ent"
	"github.com/glyphack/koal/ent/enttest"
	authuser "github.com/glyphack/koal/internal/module/auth/domain/user"
	authinfra "github.com/glyphack/koal/internal/module/auth/infrastructure"
)

func TestCreateUser(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	userRepo := authinfra.UserRepositoryDB{Client: client.User}
	user := &authuser.User{}
	user.SetEmailAddress("email@test.com")
	user.SetPassword("password")
	err := userRepo.CreateUser(context.Background(), user)
	assert.Nil(t, err)
}

func TestCreateUserDuplicate(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	userRepo := authinfra.UserRepositoryDB{Client: client.User}
	user := &authuser.User{}
	user.SetEmailAddress("email@test.com")
	user.SetPassword("password")
	err := userRepo.CreateUser(context.Background(), user)
	assert.Nil(t, err)

	err = userRepo.CreateUser(context.Background(), user)
	assert.Error(t, ent.ConstraintError{}, err)
}

func TestDeleteUser(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	userRepo := authinfra.UserRepositoryDB{Client: client.User}
	user := &authuser.User{}
	user.SetEmailAddress("email@test.com")
	user.SetPassword("password")
	err := userRepo.CreateUser(context.Background(), user)
	assert.Nil(t, err)

	err = userRepo.DeleteUser(context.Background(), user.Email)
	assert.Nil(t, err)
}

func TestDeleteUserNotExists(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	userRepo := authinfra.UserRepositoryDB{Client: client.User}
	err := userRepo.DeleteUser(context.Background(), "some_email")
	assert.NotNil(t, err)
}

func TestUpdateUser(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	userRepo := authinfra.UserRepositoryDB{Client: client.User}
	user := &authuser.User{}
	user.SetEmailAddress("email@test.com")
	user.SetPassword("password")
	err := userRepo.CreateUser(context.Background(), user)
	assert.Nil(t, err)

	newPassword := "NewPassword"
	updatedUser := &authuser.User{Email: user.Email, Password: newPassword}
	err = userRepo.UpdateUser(context.Background(), updatedUser)
	assert.Nil(t, err)

	user, err = userRepo.GetUser(context.Background(), user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.Password, newPassword)
}

func TestUpdateUserNotExists(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	userRepo := authinfra.UserRepositoryDB{Client: client.User}
	updatedUser := &authuser.User{Email: "email", Password: "NewPassword"}
	err := userRepo.UpdateUser(context.Background(), updatedUser)
	assert.NotNil(t, err)
}
