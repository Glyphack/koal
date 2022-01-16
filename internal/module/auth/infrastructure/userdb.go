package authinfra

import (
	"context"
	"errors"

	"github.com/glyphack/koal/ent"
	"github.com/glyphack/koal/ent/user"
	authuser "github.com/glyphack/koal/internal/module/auth/domain/user"
	log "github.com/sirupsen/logrus"
)

type UserRepositoryDB struct {
	Client *ent.UserClient
}

func (u *UserRepositoryDB) CreateUser(ctx context.Context, newUser *authuser.User) error {
	err := u.Client.Create().SetEmail(newUser.Email).SetPassword(newUser.Password).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepositoryDB) GetUser(ctx context.Context, id string) (error, *authuser.User) {
	dbUser, err := u.Client.Query().Where(user.Email(id)).First(ctx)
	if err != nil {
		return err, nil
	}
	return nil, &authuser.User{
		Email:    dbUser.Email,
		Password: dbUser.Password,
	}
}
func (u *UserRepositoryDB) DeleteUser(ctx context.Context, id string) error {
	count, err := u.Client.Delete().Where(user.Email(id)).Exec(ctx)
	if err != nil {
		return err
	}
	if count < 1 {
		return errors.New("No user deleted")
	}
	if count > 1 {
		log.WithFields(log.Fields{"count": count, "id": id}).Warn("Multiple users deleted with one id")
	}
	return nil
}

func (u *UserRepositoryDB) UpdateUser(ctx context.Context, newUser *authuser.User) error {
	updated, err := u.Client.Update().Where(user.Email(newUser.Email)).SetPassword(newUser.Password).Save(ctx)

	if err != nil {
		return err
	}

	if updated != 1 {
		return errors.New("No user updated")
	}

	return nil
}
