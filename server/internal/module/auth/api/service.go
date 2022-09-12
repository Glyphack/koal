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

func (s *server) Register(
	ctx context.Context,
	in *authv1pb.RegisterRequest,
) (*authv1pb.RegisterResponse, error) {
	newUser := &authuser.User{}
	if err := newUser.SetEmailAddress(in.GetEmail()); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := newUser.SetPassword(in.GetPassword()); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	user, _ := s.userRepository.GetUser(ctx, in.GetEmail())
	if user != nil {
		return nil, status.Error(
			codes.AlreadyExists,
			"There is already a user with this email address",
		)
	}
	err := s.userRepository.CreateUser(ctx, newUser)
	if err != nil {
		log.WithError(err).Error("Error while saving registered user")
	}
	log.Info("New user registered")

	token, err := newUser.GenerateToken()

	if err != nil {
		log.WithError(err).Error("Failed generating token")
		return nil, status.Error(codes.Internal, "Cannot generate JWT token")
	}

	return &authv1pb.RegisterResponse{Token: token}, nil
}

func (s *server) Login(
	ctx context.Context,
	in *authv1pb.LoginRequest,
) (*authv1pb.LoginResponse, error) {
	user, err := s.userRepository.GetUser(ctx, in.Email)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, status.Error(codes.NotFound, "Incorrect email or password")
		}
		log.WithError(err).Error("Error getting user from DB")
		return nil, status.Error(codes.Unknown, "Unknown error")
	}

	passwordCorrect := user.CheckPassword(in.Password)
	if passwordCorrect == false {
		return nil, status.Error(codes.Unauthenticated, "Incorrect email or password")
	}

	token, err := user.GenerateToken()
	if err != nil {
		log.WithError(err).Error("Failed generating token")
		return nil, status.Error(codes.Internal, "Cannot generate JWT token")
	}

	return &authv1pb.LoginResponse{Token: token}, nil
}

func (s *server) AuthFuncOverride(
	ctx context.Context,
	fullMethodName string,
) (context.Context, error) {
	return ctx, nil
}
