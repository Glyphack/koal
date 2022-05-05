package todoinfra_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/glyphack/koal/ent"
	"github.com/glyphack/koal/ent/enttest"
	tododomain "github.com/glyphack/koal/internal/module/todo/domain/todo"
	todoitem "github.com/glyphack/koal/internal/module/todo/domain/todo"
	todoinfra "github.com/glyphack/koal/internal/module/todo/infrastructure"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	Client *ent.Client
	ItemDB todoinfra.ItemDB
}

func (suite *Suite) TearDownTest() {
	suite.Client.Close()
}

func (suite *Suite) SetupTest() {
	client := enttest.Open(suite.T(), "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	suite.Client = client
	suite.ItemDB = todoinfra.ItemDB{
		ProjectClient: client.Project,
		ItemClient:    client.TodoItem,
	}
}

func TestTodoUseCase(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (suite *Suite) Test_db_todo_GetItemById() {
	type args struct {
		ctx    context.Context
		itemId string
	}
	tests := []struct {
		name      string
		setupItem *tododomain.TodoItem
		args      args
		want      *tododomain.TodoItem
		wantErr   bool
	}{
		{
			name: "query item without project with valid ID",
			args: args{ctx: context.Background(), itemId: "f47ac10b-58cc-0372-8567-0e02b2c3d479"},
			setupItem: &todoitem.TodoItem{
				UUId:    uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
				Title:   "item1",
				OwnerId: "user1",
				Project: nil,
				IsDone:  false,
			},
			want: &todoitem.TodoItem{
				UUId:    uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
				Title:   "item1",
				OwnerId: "user1",
				Project: nil,
				IsDone:  false,
			},
			wantErr: false,
		},
		{
			name: "query item with project with valid ID",
			args: args{ctx: context.Background(), itemId: "f47ac10b-58cc-0372-8567-0e02b2c3d478"},
			setupItem: &todoitem.TodoItem{
				UUId:    uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d478"),
				Title:   "item1",
				OwnerId: "user1",
				Project: &tododomain.Project{
					UUId:    uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d478"),
					Name:    "project1",
					OwnerId: "user1",
				},
				IsDone: false,
			},
			want: &todoitem.TodoItem{
				UUId:    uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d478"),
				Title:   "item1",
				OwnerId: "user1",
				Project: &tododomain.Project{
					UUId:    uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d478"),
					Name:    "project1",
					OwnerId: "user1",
				},
				IsDone: false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			t := suite.T()
			if tt.setupItem.Project != nil {
				err := suite.ItemDB.CreateProject(tt.args.ctx, tt.setupItem.Project)
				if err != nil {
					t.Fatal(err)
				}
			}
			err := suite.ItemDB.CreateItem(tt.args.ctx, tt.setupItem)
			if err != nil {
				t.Fatal(err)
			}
			got, err := suite.ItemDB.GetItemById(tt.args.ctx, tt.args.itemId)

			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("result = %v, want %v", got, tt.want)
			}
		})
	}
}

func (suite *Suite) Test_db_todo_UpdateItem() {
	type args struct {
		ctx         context.Context
		itemId      string
		updatedItem *tododomain.TodoItem
	}
	tests := []struct {
		name      string
		setupItem *tododomain.TodoItem
		args      args
		want      *tododomain.TodoItem
		wantErr   bool
	}{
		{
			name: "can update item is done",
			setupItem: &todoitem.TodoItem{
				UUId:    uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d480"),
				Title:   "title",
				OwnerId: "user1",
				Project: nil,
				IsDone:  false,
			},
			args: args{
				ctx:    context.Background(),
				itemId: "f47ac10b-58cc-0372-8567-0e02b2c3d480",
				updatedItem: &todoitem.TodoItem{
					UUId:    uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d480"),
					Title:   "title",
					OwnerId: "owner1",
					Project: nil,
					IsDone:  true,
				},
			},
			want: &todoitem.TodoItem{
				UUId:    uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d480"),
				Title:   "title",
				OwnerId: "user1",
				Project: nil,
				IsDone:  true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			t := suite.T()
			if tt.setupItem.Project != nil {
				err := suite.ItemDB.CreateProject(tt.args.ctx, tt.setupItem.Project)
				if err != nil {
					t.Fatal(err)
				}
			}
			err := suite.ItemDB.CreateItem(tt.args.ctx, tt.setupItem)
			if err != nil {
				t.Fatal(err)
			}
			err = suite.ItemDB.UpdateItem(tt.args.ctx, tt.args.itemId, tt.args.updatedItem)
			got, err := suite.ItemDB.GetItemById(tt.args.ctx, tt.args.itemId)

			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("result = %v, want %v", got, tt.want)
			}
		})
	}
}

func (suite *Suite) Test_db_todo_CreateProject() {
	type args struct {
		ctx context.Context
		in1 *tododomain.Project
	}
	tests := []struct {
		name    string
		args    args
		want    *tododomain.ProjectInfo
		wantErr bool
	}{
		{
			name: "can create project",
			args: args{
				ctx: context.Background(),
				in1: &todoitem.Project{
					UUId:    uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
					Name:    "New Project",
					OwnerId: "user",
				},
			},
			want: &todoitem.ProjectInfo{
				Project: &todoitem.Project{
					UUId:    uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
					Name:    "New Project",
					OwnerId: "user",
				},
				Items: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			t := suite.T()
			err := suite.ItemDB.CreateProject(tt.args.ctx, tt.args.in1)
			got, err := suite.ItemDB.GetProject(tt.args.ctx, tt.args.in1.UUId.String())

			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("result = %v, want %v", got, tt.want)
			}
		})
	}
}

func (suite *Suite) Test_db_todo_DeleteProject() {
	ctx := context.Background()
	project := &todoitem.Project{
		UUId:    uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d480"),
		OwnerId: "user1",
		Name:    "project",
	}
	suite.ItemDB.CreateProject(ctx, project)
	type args struct {
		ctx context.Context
		Id  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "can delete item",
			args: args{
				ctx: context.Background(),
				Id:  project.UUId.String(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			t := suite.T()
			err := suite.ItemDB.DeleteProject(tt.args.ctx, tt.args.Id)

			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// after successful delete, the project should not be found

			if _, err = suite.ItemDB.GetProject(tt.args.ctx, tt.args.Id); !errors.Is(err, todoinfra.NotFoundErr) {
				t.Errorf("item is not deleted id = %v, err = %v", tt.args.Id, err)
			}
		})
	}
}

func (suite *Suite) Test_db_todo_UpdateProjectById() {
	ctx := context.Background()
	project := &todoitem.Project{
		UUId:    uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d480"),
		OwnerId: "user1",
		Name:    "project",
	}
	suite.ItemDB.CreateProject(ctx, project)
	type args struct {
		ctx  context.Context
		Id   string
		Name string
	}
	tests := []struct {
		name        string
		args        args
		wantProject *tododomain.ProjectInfo
		wantErr     bool
	}{
		{
			name: "can update project title",
			args: args{
				ctx:  context.Background(),
				Id:   project.UUId.String(),
				Name: "new project",
			},
			wantProject: &todoitem.ProjectInfo{
				Project: &todoitem.Project{UUId: project.UUId, Name: "new project", OwnerId: project.OwnerId},
				Items:   nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			t := suite.T()
			err := suite.ItemDB.UpdateProjectById(tt.args.ctx, tt.args.Id, tt.args.Name)

			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			gotProject, err := suite.ItemDB.GetProject(tt.args.ctx, tt.args.Id)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(gotProject.Project, tt.wantProject.Project) {
				t.Errorf("projects does not match want = %v, got = %v", tt.wantProject.Project, gotProject.Project)
			}
		},
		)
	}
}

func (suite *Suite) Test_db_todo_DeleteItem() {
	ctx := context.Background()
	item := &todoitem.TodoItem{
		UUId:    uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d480"),
		Title:   "title",
		OwnerId: "user1",
		Project: nil,
		IsDone:  false,
	}
	suite.ItemDB.CreateItem(ctx, item)
	type args struct {
		ctx context.Context
		Id  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "can delete item",
			args: args{
				ctx: context.Background(),
				Id:  item.UUId.String(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			t := suite.T()
			err := suite.ItemDB.DeleteItem(tt.args.ctx, tt.args.Id)

			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// after successful delete, the item should not be found

			if _, err = suite.ItemDB.GetItemById(tt.args.ctx, tt.args.Id); !errors.Is(err, todoinfra.NotFoundErr) {
				t.Errorf("item is not deleted id = %v, err = %v", tt.args.Id, err)
			}
		})
	}
}

func (suite *Suite) Test_db_todo_AllUndoneItems() {
	ctx := context.Background()

	project := &todoitem.Project{
		UUId:    [16]byte{},
		Name:    "project1",
		OwnerId: "user1",
	}
	doneItem := &todoitem.TodoItem{
		UUId:    uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d480"),
		Title:   "title",
		OwnerId: "user1",
		Project: project,
		IsDone:  true,
	}
	UndoneItem := &todoitem.TodoItem{
		UUId:    uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d480"),
		Title:   "title",
		OwnerId: "user1",
		Project: project,
		IsDone:  false,
	}
	suite.ItemDB.CreateProject(ctx, project)
	suite.ItemDB.CreateItem(ctx, doneItem)
	suite.ItemDB.CreateItem(ctx, UndoneItem)
	type args struct {
		ctx     context.Context
		OwnerId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "only returns undone items",
			args: args{
				ctx:     context.Background(),
				OwnerId: doneItem.OwnerId,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			t := suite.T()
			items, err := suite.ItemDB.AllUndoneItems(tt.args.ctx, tt.args.OwnerId)
			for _, item := range items {
				if item.IsDone {
					t.Errorf("item is done, want undone, item = %v", item)
				}
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestItemDB_AllProjects(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	ctx := context.Background()
	project, err := client.Project.Create().SetOwnerID("test").SetName("test").Save(ctx)
	if err != nil {
		t.Fatal(err)
	}
	itemDb := todoinfra.ItemDB{
		ProjectClient: client.Project,
	}
	projects, err := itemDb.GetAllMemberProjects(ctx, project.OwnerID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, project.UUID, projects[0].UUId)
}

func TestItemDB_CreateItem(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	ctx := context.Background()

	project := client.Project.Create().SetOwnerID("test").SetName("project").SaveX(ctx)

	todoRepo := todoinfra.ItemDB{ItemClient: client.TodoItem, ProjectClient: client.Project}
	err := todoRepo.CreateItem(ctx, &todoitem.TodoItem{
		UUId:    uuid.UUID{},
		Title:   "new task",
		OwnerId: "test",
		Project: &todoitem.Project{
			UUId:    project.UUID,
			Name:    "",
			OwnerId: "",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	createdItem := client.TodoItem.Query().FirstX(ctx)
	assert.Equal(t, createdItem.Title, "new task")
	assert.Equal(t, createdItem.QueryProject().FirstIDX(ctx), project.ID)
}
