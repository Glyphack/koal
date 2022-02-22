package todoapi

import (
	"context"
	"fmt"
	"github.com/glyphack/koal/ent"
	todov1 "github.com/glyphack/koal/gen/proto/go/todo/v1"
	todoitem "github.com/glyphack/koal/internal/module/todo/domain/todo"
	todoinfra "github.com/glyphack/koal/internal/module/todo/infrastructure"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	itemRepository todoinfra.TodoRepository
}

func (s server) GetProjects(ctx context.Context, empty *emptypb.Empty) (*todov1.GetProjectsResponse, error) {
	projects, err := s.itemRepository.GetAllMemberProjects(ctx, fmt.Sprint(ctx.Value("userId")))
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
	if request.GetProjectId() == "" {
		return nil, status.Error(codes.InvalidArgument, "Project must be specified")
	}

	projectId, err := uuid.Parse(request.GetProjectId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "ProjectId is invalid")
	}
	todoItem := todoitem.Item{
		UUId:    uuid.New(),
		Title:   request.Title,
		OwnerId: fmt.Sprint(ctx.Value("userId")),
		Project: &todoitem.Project{UUId: projectId},
	}
	err = s.itemRepository.CreateItem(ctx, &todoItem)
	if err != nil {
		return nil, status.Error(codes.Internal, "Cannot create todo item")
	}
	// TODO return project name
	return &todov1.CreateTodoItemResponse{
		CreatedItem: &todov1.TodoItem{
			Id:     todoItem.UUId.String(),
			Title:  todoItem.Title,
			IsDone: false,
			Project: &todov1.Project{
				Id: todoItem.Project.UUId.String(),
			},
		},
	}, nil
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
	ownerId := fmt.Sprint(ctx.Value("userId"))
	items, err := s.itemRepository.AllUndoneItems(ctx, ownerId)
	if err != nil {
		return nil, status.Error(codes.Internal, "Cannot retrieve items")
	}
	var undoneItems []*todov1.TodoItem
	for _, item := range items {
		undoneItems = append(undoneItems, &todov1.TodoItem{
			Id:     item.UUId.String(),
			Title:  item.Title,
			IsDone: false,
			Project: &todov1.Project{
				Id:   item.Project.UUId.String(),
				Name: item.Project.Name,
			},
		})
	}
	return &todov1.GetUndoneListResponse{Items: undoneItems}, nil
}

func NewServer(dbConnection *ent.Client) *server {
	return &server{itemRepository: todoinfra.ItemDB{ItemClient: dbConnection.TodoItem, ProjectClient: dbConnection.Project}}
}
