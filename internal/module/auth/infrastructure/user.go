package authinfra

import (
	"context"

	authuser "github.com/glyphack/koal/internal/module/auth/domain/user"
)

type UserRepository interface {
	CreateUser(ctx context.Context, newUser *authuser.User) error
	GetUser(ctx context.Context, id string) (error, *authuser.User)
	DeleteUser(ctx context.Context, id string) error
	UpdateUser(ctx context.Context, newUser *authuser.User) (error)
}
