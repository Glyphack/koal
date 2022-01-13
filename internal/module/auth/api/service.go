package authapi

import (
	"context"

	authv1pb "github.com/glyphack/koal/gen/proto/go/auth/v1"
	authuser "github.com/glyphack/koal/internal/module/auth/domain/user"
	authinfra "github.com/glyphack/koal/internal/module/auth/infrastructure"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	userRepository authinfra.UserRepository
}

func NewServer() *server {
	return &server{}
}

func (s *server) Register(ctx context.Context, in *authv1pb.RegisterRequest) (*authv1pb.RegisterResponse, error) {
	newUser := authuser.User{}
	if err := newUser.SetEmailAddress(in.GetEmail()); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := newUser.SetPassword(in.GetPassword()); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// s.userRepository.CreateUser(ctx, newUser)

	return &authv1pb.RegisterResponse{Token: "dummy"}, nil
}
