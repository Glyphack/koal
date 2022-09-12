package todoapi_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/glyphack/koal/ent"
	"github.com/glyphack/koal/ent/enttest"
	todov1 "github.com/glyphack/koal/gen/proto/go/todo/v1"
	todoapi "github.com/glyphack/koal/internal/module/todo/api"
	tododomain "github.com/glyphack/koal/internal/module/todo/domain/todo"
	todoinfra "github.com/glyphack/koal/internal/module/todo/infrastructure"
	"github.com/google/uuid"

	todousecase "github.com/glyphack/koal/internal/module/todo/usecase"
	"github.com/glyphack/koal/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Suite struct {
	suite.Suite
	Client            *ent.Client
	TodoRepository    todoinfra.TodoRepository
	useCaseInteractor todousecase.TodoUseCase
}

func (suite *Suite) TearDownTest() {
	suite.Client.Close()
}

func (suite *Suite) SetupTest() {
	client := enttest.Open(suite.T(), "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	suite.Client = client
	suite.TodoRepository = todoinfra.ItemDB{
		ProjectClient: client.Project,
		ItemClient:    client.TodoItem,
		Client:        client,
	}
}

func TestTodoUseCase(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (suite *Suite) Test_server_GetProjectDetails() {
	projectWithoutItem := &tododomain.Project{
		UUId:    uuid.New(),
		Name:    "projectWithoutItem",
		OwnerId: "user1",
	}
	err := suite.TodoRepository.CreateProject(context.Background(), projectWithoutItem)
	if err != nil {
		suite.T().Fatal(err)
	}

	projectWithItem := &tododomain.Project{
		UUId:    uuid.New(),
		Name:    "projectWithItem",
		OwnerId: "userWithItems",
	}
	err = suite.TodoRepository.CreateProject(context.Background(), projectWithItem)
	if err != nil {
		suite.T().Fatal(err)
	}
	item := &tododomain.TodoItem{
		UUId:    uuid.New(),
		Title:   "item1",
		OwnerId: projectWithItem.OwnerId,
		Project: projectWithItem,
		IsDone:  false,
	}
	err = suite.TodoRepository.CreateItem(context.Background(), item)
	if err != nil {
		suite.T().Fatal(err)
	}
	type args struct {
		ctx context.Context
		in1 *todov1.GetProjectDetailsRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *todov1.GetProjectDetailsResponse
		wantErr bool
	}{
		{
			name: "Get project details return project info correctly",
			args: args{
				ctx: testutils.GetAuthenticatedContext(
					context.Background(),
					projectWithoutItem.OwnerId,
				),
				in1: &todov1.GetProjectDetailsRequest{
					Id: projectWithoutItem.UUId.String(),
				},
			},
			want: &todov1.GetProjectDetailsResponse{
				Info: &todov1.Project{
					Id:   projectWithoutItem.UUId.String(),
					Name: projectWithoutItem.Name,
				},
				Items: nil,
			},
			wantErr: false,
		}, {
			name: "Get project details returns project items",
			args: args{
				ctx: testutils.GetAuthenticatedContext(
					context.Background(),
					projectWithItem.OwnerId,
				),
				in1: &todov1.GetProjectDetailsRequest{
					Id: projectWithItem.UUId.String(),
				},
			},
			want: &todov1.GetProjectDetailsResponse{
				Info: &todov1.Project{
					Id:   projectWithItem.UUId.String(),
					Name: projectWithItem.Name,
				},
				Items: []*todov1.TodoItem{{
					Id:     item.UUId.String(),
					Title:  item.Title,
					IsDone: item.IsDone,
					Project: &todov1.Project{
						Id:   projectWithItem.UUId.String(),
						Name: projectWithItem.Name,
					},
				}},
			},
			wantErr: false,
		}, {
			name: "return permission denied for un authorized user",
			args: args{
				ctx: testutils.GetAuthenticatedContext(context.Background(), "NotOwner"),
				in1: &todov1.GetProjectDetailsRequest{
					Id: projectWithItem.UUId.String(),
				},
			},
			want: nil, wantErr: true,
		}, {
			name: "return not found for un authorized user",
			args: args{
				ctx: testutils.GetAuthenticatedContext(
					context.Background(),
					projectWithItem.OwnerId,
				),
				in1: &todov1.GetProjectDetailsRequest{
					Id: uuid.NewString(),
				},
			},
			want: nil, wantErr: true,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			t := suite.T()
			s := todoapi.NewServer(suite.Client)
			got, err := s.GetProjectDetails(tt.args.ctx, tt.args.in1)

			if (err != nil) != tt.wantErr {
				t.Errorf("server.GetProjects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.GetProjects() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (suite *Suite) Test_server_CreateProject() {
	type args struct {
		ctx context.Context
		in1 *todov1.CreateProjectRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *todov1.CreateProjectResponse
		wantErr bool
	}{
		{
			name: "creates new project with given name",
			args: args{
				ctx: testutils.GetAuthenticatedContext(context.Background(), "user"),
				in1: &todov1.CreateProjectRequest{
					Name: "NewProject",
				}},
			want: &todov1.CreateProjectResponse{
				CreatedProject: &todov1.Project{
					Name: "NewProject",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			t := suite.T()
			s := todoapi.NewServer(suite.Client)
			got, err := s.CreateProject(tt.args.ctx, tt.args.in1)

			if err != nil {
				assert.Equal(t, err, tt.wantErr)
				return
			}

			if got.GetCreatedProject().Name != tt.want.CreatedProject.Name {
				t.Errorf("server.GetProjects() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (suite *Suite) Test_server_EditProject() {
	projectWithoutItem := &tododomain.Project{
		UUId:    uuid.New(),
		Name:    "projectWithoutItem",
		OwnerId: "user1",
	}
	err := suite.TodoRepository.CreateProject(context.Background(), projectWithoutItem)
	if err != nil {
		suite.T().Fatal(err)
	}
	type args struct {
		ctx context.Context
		in1 *todov1.EditProjectRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *todov1.EditProjectResponse
		wantErr bool
	}{
		{
			name: "edit project will change name",
			args: args{
				ctx: testutils.GetAuthenticatedContext(
					context.Background(),
					projectWithoutItem.OwnerId,
				),
				in1: &todov1.EditProjectRequest{
					Project: &todov1.Project{
						Id:   projectWithoutItem.UUId.String(),
						Name: "New Name",
					},
				}},
			want: &todov1.EditProjectResponse{
				UpdatedProject: &todov1.Project{
					Id:   projectWithoutItem.UUId.String(),
					Name: "New Name",
				},
			},
			wantErr: false,
		},
		{
			name: "un authorized user cannot edit project",
			args: args{
				ctx: testutils.GetAuthenticatedContext(context.Background(), "WrongUser"),
				in1: &todov1.EditProjectRequest{
					Project: &todov1.Project{
						Id:   projectWithoutItem.UUId.String(),
						Name: "Name",
					},
				}},
			want:    nil,
			wantErr: true,
		}, {
			name: "cannot edit non existing project",
			args: args{
				ctx: testutils.GetAuthenticatedContext(context.Background(), "WrongUser"),
				in1: &todov1.EditProjectRequest{
					Project: &todov1.Project{
						Id:   uuid.NewString(),
						Name: "Name",
					},
				}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			t := suite.T()
			s := todoapi.NewServer(suite.Client)
			got, err := s.EditProject(tt.args.ctx, tt.args.in1)

			if tt.wantErr != false {
				assert.NotNil(t, err)
				return
			}

			if got.GetUpdatedProject().Name != tt.want.UpdatedProject.Name {
				t.Errorf("server.GetProjects() = %v, want %v", got.UpdatedProject, tt.want)
			}

			// Check update is persisted in db
			if tt.wantErr == false {
				updateProject, _ := suite.TodoRepository.GetProject(
					tt.args.ctx,
					tt.args.in1.Project.Id,
				)
				assert.Equal(t, updateProject.Project.Name, tt.args.in1.Project.Name)
			}
		})
	}
}
func (suite *Suite) Test_server_DeleteProject() {
	projectWithoutItem := &tododomain.Project{
		UUId:    uuid.New(),
		Name:    "projectWithoutItem",
		OwnerId: "user1",
	}
	err := suite.TodoRepository.CreateProject(context.Background(), projectWithoutItem)
	if err != nil {
		suite.T().Fatal(err)
	}
	type args struct {
		ctx context.Context
		in1 *todov1.DeleteProjectRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "delete project will delete from database",
			args: args{
				ctx: testutils.GetAuthenticatedContext(
					context.Background(),
					projectWithoutItem.OwnerId,
				),
				in1: &todov1.DeleteProjectRequest{
					Id: projectWithoutItem.UUId.String(),
				}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			t := suite.T()
			s := todoapi.NewServer(suite.Client)
			_, err := s.DeleteProject(tt.args.ctx, tt.args.in1)

			if tt.wantErr != false {
				assert.NotNil(t, err)
				return
			}

			_, err = suite.TodoRepository.GetProject(tt.args.ctx, projectWithoutItem.UUId.String())
			assert.NotNil(t, err)
		})
	}
}

func (suite *Suite) Test_server_DeleteTodoItem() {
	projectWithItem := &tododomain.Project{
		UUId:    uuid.New(),
		Name:    "projectWithItem",
		OwnerId: "user1",
	}
	err := suite.TodoRepository.CreateProject(context.Background(), projectWithItem)
	if err != nil {
		suite.T().Fatal(err)
	}

	item := &tododomain.TodoItem{
		UUId:    uuid.New(),
		Title:   "item1",
		OwnerId: projectWithItem.OwnerId,
		Project: projectWithItem,
		IsDone:  false,
	}
	err = suite.TodoRepository.CreateItem(context.Background(), item)
	if err != nil {
		suite.T().Fatal(err)
	}
	type args struct {
		ctx context.Context
		in1 *todov1.DeleteTodoItemRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "delete project will delete from database",
			args: args{
				ctx: testutils.GetAuthenticatedContext(
					context.Background(),
					projectWithItem.OwnerId,
				),
				in1: &todov1.DeleteTodoItemRequest{
					Id: item.UUId.String(),
				}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			t := suite.T()
			s := todoapi.NewServer(suite.Client)
			_, err := s.DeleteTodoItem(tt.args.ctx, tt.args.in1)

			if tt.wantErr != false {
				assert.NotNil(t, err)
				return
			}

			_, err = suite.TodoRepository.GetItemById(tt.args.ctx, item.UUId.String())
			assert.NotNil(t, err)
		})
	}
}

func (suite *Suite) Test_server_UpdateTodoItem() {
	projectWithItem := &tododomain.Project{
		UUId:    uuid.New(),
		Name:    "projectWithItem",
		OwnerId: "user1",
	}
	err := suite.TodoRepository.CreateProject(context.Background(), projectWithItem)
	if err != nil {
		suite.T().Fatal(err)
	}

	item := &tododomain.TodoItem{
		UUId:    uuid.New(),
		Title:   "item1",
		OwnerId: projectWithItem.OwnerId,
		Project: projectWithItem,
		IsDone:  false,
	}
	err = suite.TodoRepository.CreateItem(context.Background(), item)
	if err != nil {
		suite.T().Fatal(err)
	}
	type args struct {
		ctx context.Context
		in1 *todov1.UpdateTodoItemRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *todov1.UpdateTodoItemRequest
		wantErr bool
	}{
		{
			name: "should update title",
			args: args{
				ctx: testutils.GetAuthenticatedContext(
					context.Background(),
					projectWithItem.OwnerId,
				),
				in1: &todov1.UpdateTodoItemRequest{
					Id:     item.UUId.String(),
					Title:  "NewTitle",
					IsDone: false,
				}},
			want: &todov1.UpdateTodoItemRequest{
				Id:     item.UUId.String(),
				Title:  "NewTitle",
				IsDone: false,
			},
			wantErr: false,
		}, {
			name: "should update is done",
			args: args{
				ctx: testutils.GetAuthenticatedContext(
					context.Background(),
					projectWithItem.OwnerId,
				),
				in1: &todov1.UpdateTodoItemRequest{
					Id:     item.UUId.String(),
					Title:  item.Title,
					IsDone: true,
				}},
			want: &todov1.UpdateTodoItemRequest{
				Id:     item.UUId.String(),
				Title:  item.Title,
				IsDone: true,
			},
			wantErr: false,
		},
		{
			name: "should ignore title if title is empty",
			args: args{
				ctx: testutils.GetAuthenticatedContext(
					context.Background(),
					projectWithItem.OwnerId,
				),
				in1: &todov1.UpdateTodoItemRequest{
					Id:     item.UUId.String(),
					Title:  "",
					IsDone: item.IsDone,
				}},
			want: &todov1.UpdateTodoItemRequest{
				Id:     item.UUId.String(),
				Title:  item.Title,
				IsDone: item.IsDone,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			t := suite.T()
			s := todoapi.NewServer(suite.Client)
			_, err := s.UpdateTodoItem(tt.args.ctx, tt.args.in1)

			if tt.wantErr != false {
				assert.NotNil(t, err)
				return
			}

			updatedItem, err := suite.TodoRepository.GetItemById(tt.args.ctx, item.UUId.String())
			t.Log(updatedItem)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.want.Title, updatedItem.Title)
			assert.Equal(t, tt.want.IsDone, updatedItem.IsDone)

		})
	}
}

func (suite *Suite) Test_server_GetUndoneList() {
	type args struct {
		ctx context.Context
		in1 *emptypb.Empty
	}
	tests := []struct {
		name    string
		args    args
		items   []*tododomain.TodoItem
		want    *todov1.GetUndoneListResponse
		wantErr bool
	}{{
		name: "Return all undone items from a single project",
		args: args{
			ctx: testutils.GetAuthenticatedContext(context.Background(), "user1"),
			in1: &emptypb.Empty{},
		},
		items: []*tododomain.TodoItem{{
			UUId:    uuid.New(),
			Title:   "item1",
			OwnerId: "user1",
			Project: &tododomain.Project{
				UUId:    uuid.New(),
				Name:    "projectWithItem",
				OwnerId: "user1",
			},
			IsDone:      false,
			Description: "test",
		}},
		want: &todov1.GetUndoneListResponse{
			Items: []*todov1.TodoItem{
				{
					Id:     uuid.NewString(),
					Title:  "item1",
					IsDone: false,
					Project: &todov1.Project{
						Id:   uuid.NewString(),
						Name: "projectWithItem",
					},
					Description: "test",
				},
			},
		},
		wantErr: false,
	}, {
		name: "only returns user owned items",
		args: args{
			ctx: testutils.GetAuthenticatedContext(context.Background(), "wrongUser"),
			in1: &emptypb.Empty{},
		},
		items: []*tododomain.TodoItem{{
			UUId:    uuid.New(),
			Title:   "item1",
			OwnerId: "user1",
			Project: &tododomain.Project{
				UUId:    uuid.New(),
				Name:    "projectWithItem",
				OwnerId: "user1",
			},
			IsDone: false,
		}},
		want: &todov1.GetUndoneListResponse{
			Items: nil,
		},
		wantErr: false,
	},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			t := suite.T()
			for _, item := range tt.items {
				err := suite.TodoRepository.CreateProject(tt.args.ctx, item.Project)
				if err != nil {
					t.Fatal(err)
				}
				err = suite.TodoRepository.CreateItem(tt.args.ctx, item)
				if err != nil {
					t.Fatal(err)
				}
			}
			s := todoapi.NewServer(suite.Client)
			response, err := s.GetUndoneList(tt.args.ctx, tt.args.in1)

			if tt.wantErr != false {
				assert.NotNil(t, err)
				return
			}

			assert.Nil(t, err)

			assert.Len(t, response.Items, len(tt.want.Items))
			for i, respItems := range response.Items {
				assert.Equal(t, respItems.Title, tt.want.Items[i].Title)
			}

		})
	}
}

// Test cases without suite
func TestGetUserProjectsWithOneProject(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	server := todoapi.NewServer(client)
	ownerId := "Sh"
	ctx := testutils.GetAuthenticatedContext(context.Background(), ownerId)
	_, err := client.Project.Create().SetOwnerID(ownerId).SetName("testProj").Save(ctx)
	if err != nil {
		t.Fatal(err)
	}
	response, err := server.GetProjects(ctx, &emptypb.Empty{})
	assert.Nil(t, err)
	assert.Equal(t, len(response.Projects), 1)
}

func TestGetUserProjectsWithMultipleProjects(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	server := todoapi.NewServer(client)
	ownerId := "Sh"
	ctx := testutils.GetAuthenticatedContext(context.Background(), ownerId)
	_, err := client.Project.Create().SetOwnerID(ownerId).SetName("testProj1").Save(ctx)
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Project.Create().SetOwnerID(ownerId).SetName("testProj2").Save(ctx)
	if err != nil {
		t.Fatal(err)
	}
	response, err := server.GetProjects(ctx, &emptypb.Empty{})
	assert.Nil(t, err)
	assert.Equal(t, len(response.Projects), 2)
}

func TestServer_CreateTodoItem(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	server := todoapi.NewServer(client)
	ownerId := "Sh"
	ctx := testutils.GetAuthenticatedContext(context.Background(), ownerId)
	project := client.Project.Create().SetOwnerID(ownerId).SetName("testProj1").SaveX(ctx)
	request := &todov1.CreateTodoItemRequest{
		ProjectId:   project.UUID.String(),
		Title:       "new task",
		Description: "desc",
	}
	response, err := server.CreateTodoItem(ctx, request)
	assert.Nil(t, err)
	assert.Equal(t, response.GetCreatedItem().GetTitle(), request.Title)
	assert.Equal(t, response.CreatedItem.GetDescription(), request.Description)
}
