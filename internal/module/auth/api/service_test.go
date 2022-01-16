package authapi_test

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/glyphack/koal/ent/enttest"
	authv1 "github.com/glyphack/koal/gen/proto/go/auth/v1"
	authapi "github.com/glyphack/koal/internal/module/auth/api"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestRegisterValidInput(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	server := authapi.NewServer(client)
	response, err := server.Register(context.Background(), &authv1.RegisterRequest{Email: "mail@test.com", Password: "password"})
	assert.Nil(t, err)
	token, err := jwt.Parse(response.GetToken(), func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("jwt_secret")), nil
	})
	assert.True(t, token.Valid)
}

func TestRegisterInvalidEmail(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	server := authapi.NewServer(client)
	_, err := server.Register(context.Background(), &authv1.RegisterRequest{Email: "invalidmail", Password: "password"})
	responseStatus, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, responseStatus.Code(), codes.InvalidArgument)
	assert.Contains(t, responseStatus.Err().Error(), "Email")
}

func TestRegisterInvalidPassword(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	server := authapi.NewServer(client)
	_, err := server.Register(context.Background(), &authv1.RegisterRequest{Email: "email@test.com", Password: "weak"})
	responseStatus, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, responseStatus.Code(), codes.InvalidArgument)
	assert.Contains(t, responseStatus.Err().Error(), "Password")
}
