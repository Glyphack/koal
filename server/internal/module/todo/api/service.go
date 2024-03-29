package todoapi

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/glyphack/koal/ent"
	todov1 "github.com/glyphack/koal/gen/proto/go/todo/v1"
	tododomain "github.com/glyphack/koal/internal/module/todo/domain/todo"
	todoinfra "github.com/glyphack/koal/internal/module/todo/infrastructure"
	todousecase "github.com/glyphack/koal/internal/module/todo/usecase"
	"github.com/glyphack/koal/pkg/grpcerrors"
	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	log "github.com/sirupsen/logrus"
)

type server struct {
	itemRepository    todoinfra.TodoRepository
	useCaseInteractor todousecase.TodoUseCase
}

func (s server) GetProjects(
	ctx context.Context,
	_ *emptypb.Empty,
) (*todov1.GetProjectsResponse, error) {
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

func (s server) GetProjectDetails(
	ctx context.Context,
	request *todov1.GetProjectDetailsRequest,
) (*todov1.GetProjectDetailsResponse, error) {
	userId := fmt.Sprint(ctx.Value("userId"))
	project, err := s.useCaseInteractor.GetProject(ctx, userId, request.Id)
	if err != nil {
		return nil, TranslateDomainAndInfraError(err)
	}
	projectMsg := &todov1.Project{
		Id:   project.Project.UUId.String(),
		Name: project.Project.Name,
	}
	var itemsMsg []*todov1.TodoItem
	for _, item := range project.Items {
		itemsMsg = append(itemsMsg, convertDomainItemToResponseItem(item))
	}

	return &todov1.GetProjectDetailsResponse{
		Info:  projectMsg,
		Items: itemsMsg,
	}, nil
}

func (s server) CreateProject(
	ctx context.Context,
	request *todov1.CreateProjectRequest,
) (*todov1.CreateProjectResponse, error) {
	userId := fmt.Sprint(ctx.Value("userId"))
	project, err := s.useCaseInteractor.CreateProject(ctx, userId, request.Name)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &todov1.CreateProjectResponse{
		CreatedProject: &todov1.Project{
			Id:   project.UUId.String(),
			Name: project.Name,
		},
	}, nil
}

func (s server) EditProject(
	ctx context.Context,
	request *todov1.EditProjectRequest,
) (*todov1.EditProjectResponse, error) {
	userId := fmt.Sprint(ctx.Value("userId"))
	project, err := s.useCaseInteractor.UpdateProject(
		ctx,
		userId,
		request.Project.Id,
		request.Project.Name,
	)
	if err != nil {
		return nil, TranslateDomainAndInfraError(err)
	}
	return &todov1.EditProjectResponse{
		UpdatedProject: &todov1.Project{
			Id:   project.UUId.String(),
			Name: project.Name,
		},
	}, nil
}

func (s server) DeleteProject(
	ctx context.Context,
	request *todov1.DeleteProjectRequest,
) (*emptypb.Empty, error) {
	userId := fmt.Sprint(ctx.Value("userId"))
	err := s.useCaseInteractor.DeleteProject(ctx, userId, request.Id)
	if err != nil {
		return nil, TranslateDomainAndInfraError(err)
	}
	return &emptypb.Empty{}, nil
}

func (s server) CreateTodoItem(
	ctx context.Context,
	request *todov1.CreateTodoItemRequest,
) (*todov1.CreateTodoItemResponse, error) {
	projectId, err := uuid.Parse(request.GetProjectId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "ProjectId is invalid")
	}
	todoItem := tododomain.TodoItem{
		UUId:        uuid.New(),
		Title:       request.Title,
		OwnerId:     fmt.Sprint(ctx.Value("userId")),
		Project:     &tododomain.Project{UUId: projectId},
		Description: request.Description,
	}
	err = s.itemRepository.CreateItem(ctx, &todoItem)
	if err != nil {
		return nil, status.Error(codes.Internal, "Cannot create todo item")
	}

	return &todov1.CreateTodoItemResponse{
		CreatedItem: convertDomainItemToResponseItem(&todoItem),
	}, nil
}

func (s server) DeleteTodoItem(
	ctx context.Context,
	request *todov1.DeleteTodoItemRequest,
) (*emptypb.Empty, error) {
	userId := fmt.Sprint(ctx.Value("userId"))
	err := s.useCaseInteractor.DeleteItem(ctx, request.Id, userId)
	if err != nil {
		return nil, TranslateDomainAndInfraError(err)
	}
	return &emptypb.Empty{}, nil
}

func (s server) UpdateTodoItem(
	ctx context.Context,
	request *todov1.UpdateTodoItemRequest,
) (*emptypb.Empty, error) {
	userId := fmt.Sprint(ctx.Value("userId"))
	_, err := s.useCaseInteractor.UpdateItem(
		ctx,
		request.Id,
		request.Title,
		request.IsDone,
		userId,
		request.Description,
	)
	if err != nil {
		return nil, TranslateDomainAndInfraError(err)
	}
	return &emptypb.Empty{}, nil
}

func (s server) GetUndoneList(
	ctx context.Context,
	_ *emptypb.Empty,
) (*todov1.GetUndoneListResponse, error) {
	ownerId := fmt.Sprint(ctx.Value("userId"))
	items, err := s.itemRepository.AllUndoneItems(ctx, ownerId)
	if err != nil {
		st := status.New(codes.Internal, "Cannot retrieve items")
		st, _ = st.WithDetails(&errdetails.DebugInfo{
			StackEntries: []string{err.Error()},
			Detail:       "",
		})
		return nil, st.Err()
	}
	var undoneItems []*todov1.TodoItem
	for _, item := range items {
		undoneItems = append(undoneItems, convertDomainItemToResponseItem(item))
	}
	return &todov1.GetUndoneListResponse{Items: undoneItems}, nil
}

func (s server) GetTodoItems(
	ctx context.Context,
	GetTodoItemsRequest *todov1.GetTodoItemsRequest,
) (*todov1.GetTodoItemsResponse, error) {
	itemQuery := todoinfra.TodoItemQuery{}
	if GetTodoItemsRequest.ItemId != nil {
		itemUUID, err := uuid.Parse(*GetTodoItemsRequest.ItemId)
		if err != nil {
			st, _ := grpcerrors.FieldValidationErorr(
				"item_id",
				"item_id must be in uuid format",
				*GetTodoItemsRequest.ItemId,
			)
			return nil, st.Err()
		}
		itemQuery.UUId = itemUUID
	}
	if len(GetTodoItemsRequest.ProjectIds) > 0 {
		idsToQuery := []uuid.UUID{}
		for _, projectId := range strings.Split(GetTodoItemsRequest.ProjectIds, ",") {
			projectUUID, err := uuid.Parse(projectId)
			if err != nil {
				st, _ := grpcerrors.FieldValidationErorr(
					"project_id",
					"project_ids must be in uuid format",
					string(projectId),
				)
				st, _ = st.WithDetails(&errdetails.DebugInfo{
					StackEntries: []string{
						fmt.Sprintf("project id %v", projectId),
						fmt.Sprintf("error, %v", err),
					},
					Detail: "uuid parsing stack trace",
				})
				return nil, st.Err()
			}
			idsToQuery = append(idsToQuery, projectUUID)
		}
		itemQuery.ProjectIds = idsToQuery
	}
	if GetTodoItemsRequest.IsDone != nil {
		if *GetTodoItemsRequest.IsDone == true {
			itemQuery.IsDone = todoinfra.Done
		} else {
			itemQuery.IsDone = todoinfra.NotDone
		}
	}
	itemQuery.OwnerId = fmt.Sprint(ctx.Value("userId"))
	log.Debugf("Querying todo items with %v", itemQuery)
	todoItems, err := s.itemRepository.GetItems(ctx, itemQuery)
	if err != nil {
		st := status.New(codes.Internal, "Cannot retrieve items")
		st, _ = st.WithDetails(&errdetails.DebugInfo{
			StackEntries: []string{err.Error()},
			Detail:       "",
		})
		return nil, st.Err()
	}
	responseItems := []*todov1.TodoItem{}
	for _, item := range *todoItems {
		responseItems = append(responseItems, convertDomainItemToResponseItem(&item))
	}

	return &todov1.GetTodoItemsResponse{
		TodoItems: responseItems,
	}, nil
}

func NewServer(dbConnection *ent.Client) *server {
	itemRepo := &todoinfra.ItemDB{
		ItemClient:    dbConnection.TodoItem,
		ProjectClient: dbConnection.Project,
		Client:        dbConnection,
	}
	return &server{
		itemRepository:    itemRepo,
		useCaseInteractor: todousecase.TodoUseCase{TodoRepository: itemRepo},
	}
}

func TranslateDomainAndInfraError(err error) error {
	if errors.Is(err, todousecase.PermissionDenied) {
		return status.Error(codes.PermissionDenied, err.Error())
	} else if errors.Is(err, todoinfra.NotFoundErr) {
		return status.Error(codes.NotFound, err.Error())
	}
	return status.Error(codes.Internal, "internal error")
}

func convertDomainItemToResponseItem(item *tododomain.TodoItem) *todov1.TodoItem {
	return &todov1.TodoItem{
		Id:     item.UUId.String(),
		Title:  item.Title,
		IsDone: item.IsDone,
		Project: &todov1.Project{
			Id:   item.Project.UUId.String(),
			Name: item.Project.Name,
		},
		Description: item.Description,
	}
}
