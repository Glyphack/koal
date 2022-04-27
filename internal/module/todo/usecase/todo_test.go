package todousecase_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/glyphack/koal/ent"
	"github.com/glyphack/koal/ent/enttest"
	tododomain "github.com/glyphack/koal/internal/module/todo/domain/todo"
	todoinfra "github.com/glyphack/koal/internal/module/todo/infrastructure"
	todousecase "github.com/glyphack/koal/internal/module/todo/usecase"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TodoUseCaseSuite struct {
	suite.Suite
	TodoRepository todoinfra.TodoRepository
	Client         *ent.Client
}

func (suite *TodoUseCaseSuite) TearDownTest() {
	suite.Client.Close()
}

func (suite *TodoUseCaseSuite) SetupTest() {
	client := enttest.Open(suite.T(), "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	suite.Client = client
	suite.TodoRepository = todoinfra.ItemDB{
		ProjectClient: client.Project,
		ItemClient:    client.TodoItem,
	}
}

func TestTodoUseCase(t *testing.T) {
	suite.Run(t, new(TodoUseCaseSuite))
}

func ProjectFixture() *tododomain.Project {
	return &tododomain.Project{
		UUId:    uuid.New(),
		Name:    "project1",
		OwnerId: "user1",
	}
}
func (suite *TodoUseCaseSuite) TestTodoUseCase_DeleteItem() {
	type args struct {
		ctx    context.Context
		itemId string
		userId string
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		setupData   func(s *TodoUseCaseSuite, args args)
		assertState func(s *TodoUseCaseSuite, args args)
	}{
		{
			name: "test can delete existing item",
			args: args{
				ctx:    context.Background(),
				itemId: uuid.NewString(),
				userId: "user",
			},
			wantErr: false,
			setupData: func(s *TodoUseCaseSuite, args args) {
				project := &tododomain.Project{
					UUId:    uuid.New(),
					Name:    "test",
					OwnerId: args.userId,
				}
				err := s.TodoRepository.CreateProject(args.ctx, project)
				if err != nil {
					suite.T().Fatal(err)
				}
				err = s.TodoRepository.CreateItem(args.ctx, &tododomain.TodoItem{
					UUId:    uuid.MustParse(args.itemId),
					Title:   "test",
					OwnerId: args.userId,
					Project: project,
					IsDone:  false,
				})
				if err != nil {
					suite.T().Fatal(err)
				}
			},
			assertState: func(s *TodoUseCaseSuite, args args) {
				i, err := s.TodoRepository.GetItemById(args.ctx, args.itemId)
				suite.T().Log(i)
				assert.NotNil(suite.T(), err)
			},
		}, {
			name: "return error for non existing item",
			args: args{
				ctx:    context.Background(),
				itemId: uuid.NewString(),
				userId: "user",
			},
			wantErr:     true,
			setupData:   func(s *TodoUseCaseSuite, args args) {},
			assertState: func(s *TodoUseCaseSuite, args args) {},
		}, {
			name: "user without permission cannot delete item",
			args: args{
				ctx:    context.Background(),
				itemId: uuid.NewString(),
				userId: "user",
			},
			wantErr: true,
			setupData: func(s *TodoUseCaseSuite, args args) {
				project := &tododomain.Project{
					UUId:    uuid.New(),
					Name:    "test",
					OwnerId: args.userId,
				}
				err := s.TodoRepository.CreateProject(args.ctx, project)
				if err != nil {
					suite.T().Fatal(err)
				}
				err = s.TodoRepository.CreateItem(args.ctx, &tododomain.TodoItem{
					UUId:    uuid.MustParse(args.itemId),
					Title:   "test",
					OwnerId: "wrongUser",
					Project: project,
					IsDone:  false,
				})
				if err != nil {
					suite.T().Fatal(err)
				}
			},
			assertState: func(s *TodoUseCaseSuite, args args) {
				i, err := s.TodoRepository.GetItemById(args.ctx, args.itemId)
				assert.Nil(suite.T(), err)
				assert.NotNil(suite.T(), i)
			},
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			u := &todousecase.TodoUseCase{
				TodoRepository: suite.TodoRepository,
			}
			tt.setupData(suite, tt.args)
			if err := u.DeleteItem(tt.args.ctx, tt.args.itemId, tt.args.userId); (err != nil) != tt.wantErr {
				suite.T().Errorf("TodoUseCase.DeleteItem() error = %v, wantErr %v", err, tt.wantErr)
			}
			tt.assertState(suite, tt.args)
		})
	}
}

func (suite *TodoUseCaseSuite) TestTodoUseCase_CreateProject() {
	type args struct {
		ctx         context.Context
		userId      string
		projectName string
	}
	tests := []struct {
		name    string
		args    args
		want    *tododomain.Project
		wantErr bool
	}{
		{
			name: "Can Create valid project",
			args: args{
				ctx:         context.Background(),
				userId:      "user",
				projectName: "test",
			},
			want: &tododomain.Project{
				Name:    "test",
				OwnerId: "user",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			t := suite.T()
			u := todousecase.TodoUseCase{
				TodoRepository: suite.TodoRepository,
			}
			got, err := u.CreateProject(tt.args.ctx, tt.args.userId, tt.args.projectName)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoUseCase.CreateProject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Name != tt.args.projectName || got.OwnerId != tt.args.userId {
				t.Errorf("TodoUseCase.CreateProject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (suite *TodoUseCaseSuite) TestTodoUseCase_GetProject() {
	project := ProjectFixture()
	err := suite.TodoRepository.CreateProject(context.Background(), project)
	if err != nil {
		suite.T().Fatal(err)
	}
	type args struct {
		ctx       context.Context
		userId    string
		projectId string
	}
	tests := []struct {
		name    string
		args    args
		want    *tododomain.ProjectInfo
		wantErr bool
	}{
		{
			name: "can get existing project",
			args: args{
				ctx:       context.Background(),
				userId:    project.OwnerId,
				projectId: project.UUId.String(),
			},
			want: &tododomain.ProjectInfo{
				Project: &tododomain.Project{
					UUId:    project.UUId,
					Name:    project.Name,
					OwnerId: project.OwnerId,
				},
				Items: nil,
			},
			wantErr: false,
		},
		{
			name: "return error for non existing project",
			args: args{
				ctx:       context.Background(),
				userId:    "user",
				projectId: uuid.NewString(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "cannot get project without permission",
			args: args{
				ctx:       context.Background(),
				userId:    "user",
				projectId: project.UUId.String(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			t := suite.T()
			u := todousecase.TodoUseCase{
				TodoRepository: suite.TodoRepository,
			}
			got, err := u.GetProject(tt.args.ctx, tt.args.userId, tt.args.projectId)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoUseCase.GetProject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoUseCase.GetProject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (suite *TodoUseCaseSuite) TestTodoUseCase_UpdateProject() {
	project := ProjectFixture()
	err := suite.TodoRepository.CreateProject(context.Background(), project)
	if err != nil {
		suite.T().Fatal(err)
	}
	type args struct {
		ctx       context.Context
		userId    string
		projectId string
		name      string
	}
	tests := []struct {
		name    string
		args    args
		want    *tododomain.Project
		wantErr bool
	}{
		{
			name: "update will edit project",
			args: args{ctx: context.Background(), userId: project.OwnerId, projectId: project.UUId.String(), name: "newName"},
			want: &tododomain.Project{
				UUId:    project.UUId,
				Name:    "newName",
				OwnerId: project.OwnerId,
			},
			wantErr: false,
		},
		{
			name:    "cannot update without permission",
			args:    args{ctx: context.Background(), userId: "wrongUser", projectId: project.UUId.String(), name: project.Name},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "cannot update non existing project",
			args:    args{ctx: context.Background(), userId: project.OwnerId, projectId: uuid.NewString(), name: project.Name},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			t := suite.T()
			u := todousecase.TodoUseCase{
				TodoRepository: suite.TodoRepository,
			}
			got, err := u.UpdateProject(tt.args.ctx, tt.args.userId, tt.args.projectId, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoUseCase.UpdateProject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoUseCase.UpdateProject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (suite *TodoUseCaseSuite) TestTodoUseCase_DeleteProject() {
	project := ProjectFixture()
	err := suite.TodoRepository.CreateProject(context.Background(), project)
	if err != nil {
		suite.T().Fatal(err)
	}
	type args struct {
		ctx       context.Context
		userId    string
		projectId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "can delete project",
			args: args{
				ctx:       context.Background(),
				userId:    project.OwnerId,
				projectId: project.UUId.String(),
			},
			wantErr: false,
		},
		{
			name: "cannot delete project without permission",
			args: args{
				ctx:       context.Background(),
				userId:    "wrongUser",
				projectId: project.UUId.String(),
			},
			wantErr: true,
		},
		{
			name: "cannot delete non existing project",
			args: args{
				ctx:       context.Background(),
				userId:    project.OwnerId,
				projectId: uuid.NewString(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			t := suite.T()
			u := todousecase.TodoUseCase{
				TodoRepository: suite.TodoRepository,
			}
			if err := u.DeleteProject(tt.args.ctx, tt.args.userId, tt.args.projectId); (err != nil) != tt.wantErr {
				t.Errorf("TodoUseCase.DeleteProject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func (suite *TodoUseCaseSuite) TestTodoUseCase_UpdateItem() {
	testItem := &tododomain.TodoItem{
		UUId:    uuid.New(),
		Title:   "title",
		OwnerId: "owner1",
		Project: &tododomain.Project{
			UUId:    uuid.New(),
			Name:    "testProject",
			OwnerId: "owner1",
		},
		IsDone: false,
	}
	todo_repo_mock := &todoinfra.TodoRepositoryMock{
		GetItemByIdFunc: func(ctx context.Context, Id string) (*tododomain.TodoItem, error) {
			return testItem, nil
		},
		UpdateItemFunc: func(ctx context.Context, Id string, updatedItem *tododomain.TodoItem) error {
			return nil
		},
	}
	type args struct {
		ctx      context.Context
		userId   string
		itemId   string
		newTitle string
		isDone   bool
	}
	tests := []struct {
		name    string
		args    args
		want    *tododomain.TodoItem
		wantErr bool
	}{
		{
			name: "can make item done",
			args: args{
				ctx:      context.Background(),
				userId:   testItem.OwnerId,
				itemId:   testItem.UUId.String(),
				newTitle: testItem.Title,
				isDone:   true,
			},
			want: &tododomain.TodoItem{
				UUId:    testItem.UUId,
				Title:   testItem.Title,
				OwnerId: testItem.OwnerId,
				Project: testItem.Project,
				IsDone:  true,
			},
			wantErr: false,
		},
		{
			name: "can change item title",
			args: args{
				ctx:      context.Background(),
				userId:   testItem.OwnerId,
				itemId:   testItem.UUId.String(),
				newTitle: "new",
				isDone:   true,
			},
			want: &tododomain.TodoItem{
				UUId:    testItem.UUId,
				Title:   "new",
				OwnerId: testItem.OwnerId,
				Project: testItem.Project,
				IsDone:  true,
			},
			wantErr: false,
		},
		{
			name: "cannot change item title without permission",
			args: args{
				ctx:      context.Background(),
				userId:   testItem.OwnerId,
				itemId:   testItem.UUId.String(),
				newTitle: testItem.Title,
				isDone:   true,
			},
			want: &tododomain.TodoItem{
				UUId:    testItem.UUId,
				Title:   testItem.Title,
				OwnerId: testItem.OwnerId,
				Project: testItem.Project,
				IsDone:  true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			t := suite.T()
			u := todousecase.TodoUseCase{
				TodoRepository: todo_repo_mock,
			}
			actual, err := u.UpdateItem(tt.args.ctx, tt.args.itemId, tt.args.newTitle, tt.args.isDone, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoUseCase.DeleteProject() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(actual, tt.want) {
				t.Errorf("TodoUseCase.UpdateProject() = %v, want %v", actual, tt.want)
			}

		})
	}

}
