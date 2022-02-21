package todoapi

import (
	"context"
	"fmt"
	"github.com/glyphack/koal/ent"
	todov1 "github.com/glyphack/koal/gen/proto/go/todo/v1"
	todoinfra "github.com/glyphack/koal/internal/module/todo/infrastructure"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	itemRepository todoinfra.TodoRepository
}

func (s server) GetProjects(ctx context.Context, empty *emptypb.Empty) (*todov1.GetProjectsResponse, error) {
	projects, err := s.itemRepository.AllProjects(ctx, fmt.Sprint(ctx.Value("userId")))
	if err != nil {
		return nil, status.Error(codes.Internal, "Cannot load user projects")
	}
	var projectPresentations []*todov1.Project

	for _, project := range projects {
		projectPresentations = append(projectPresentations,
			&todov1.Project{
				Id:   project.UUId.String(),
				Name: project.Name,
			})
	}

	return &todov1.GetProjectsResponse{Projects: projectPresentations}, nil

}

func (s server) GetProjectDetails(ctx context.Context, request *todov1.GetProjectDetailsRequest) (*todov1.GetProjectDetailsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) CreateProject(ctx context.Context, request *todov1.CreateProjectRequest) (*todov1.CreateProjectResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) EditProject(ctx context.Context, request *todov1.EditProjectRequest) (*todov1.EditProjectResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) DeleteProject(ctx context.Context, request *todov1.DeleteProjectRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) CreateTodoItem(ctx context.Context, request *todov1.CreateTodoItemRequest) (*todov1.CreateTodoItemResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) DeleteTodoItem(ctx context.Context, request *todov1.DeleteTodoItemRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) UpdateTodoItem(ctx context.Context, request *todov1.UpdateTodoItemRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) GetUndoneList(ctx context.Context, empty *emptypb.Empty) (*todov1.GetUndoneListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewServer(dbConnection *ent.Client) *server {
	return &server{itemRepository: todoinfra.ItemDB{ItemClient: dbConnection.TodoItem, ProjectClient: dbConnection.Project}}
}