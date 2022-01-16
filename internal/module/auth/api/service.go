package authapi

import (
	"context"

	"github.com/glyphack/koal/ent"
	authv1pb "github.com/glyphack/koal/gen/proto/go/auth/v1"
	authuser "github.com/glyphack/koal/internal/module/auth/domain/user"
	authinfra "github.com/glyphack/koal/internal/module/auth/infrastructure"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	userRepository authinfra.UserRepository
}

func NewServer(dbConnection *ent.Client) *server {
	return &server{userRepository: &authinfra.UserRepositoryDB{Client: dbConnection.User}}
}

func (s *server) Register(ctx context.Context, in *authv1pb.RegisterRequest) (*authv1pb.RegisterResponse, error) {
	newUser := authuser.User{}
	if err := newUser.SetEmailAddress(in.GetEmail()); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := newUser.SetPassword(in.GetPassword()); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	s.userRepository.CreateUser(ctx, &newUser)

	token, err := newUser.GenerateToken()

	if err != nil {
		log.WithError(err).Error("Failed generating token")
		return nil, status.Error(codes.Internal, "Cannot generate JWT token")
	}

	return &authv1pb.RegisterResponse{Token: token}, nil
}
