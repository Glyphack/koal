package authapi

import (
	"context"

	authv1 "github.com/glyphack/koal/gen/proto/go/auth/v1"
	authuser "github.com/glyphack/koal/internal/module/auth/domain/user"
	authinfra "github.com/glyphack/koal/internal/module/auth/infrastructure"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	userRepository authinfra.UserRepository
}

func NewServer() *server {
	return &server{}
}

func RegisterServer(s *grpc.Server) {
	authServer := NewServer()
	authv1.RegisterAuthServiceServer(s, authServer)
}

func (s *server) Register(ctx context.Context, in *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	newUser := authuser.User{}
	if err := newUser.SetEmailAddress(in.GetEmail()); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := newUser.SetPassword(in.GetPassword()); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	s.userRepository.CreateUser(ctx, newUser)

	return &authv1.RegisterResponse{Token: "dummy"}, nil
}
