package todoinfra_test

import (
	"context"
	todoitem "github.com/glyphack/koal/internal/module/todo/domain/todo"
	todoinfra "github.com/glyphack/koal/internal/module/todo/infrastructure"
	"github.com/glyphack/koal/pkg/testutils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func beforeEach(t *testing.T) *testutils.TestWithDBClient {
	return &testutils.TestWithDBClient{Client: testutils.SetupTestEntClient(t)}
}

func TestItemDB_AllProjects(t *testing.T) {
	test := beforeEach(t)
	ctx := context.Background()
	project, err := test.Client.Project.Create().SetOwnerID("test").SetName("test").Save(ctx)
	if err != nil {
		t.Fatal(err)
	}
	itemDb := todoinfra.ItemDB{
		ProjectClient: test.Client.Project,
	}
	projects, err := itemDb.GetAllMemberProjects(ctx, project.OwnerID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, project.UUId, projects[0].UUId)
}

func TestItemDB_CreateItem_WithoutProject(t *testing.T) {
	test := beforeEach(t)
	ctx := context.Background()
	todoRepo := todoinfra.ItemDB{ItemClient: test.Client.TodoItem}
	err := todoRepo.CreateItem(ctx, &todoitem.Item{
		UUId:    uuid.UUID{},
		Title:   "new task",
		OwnerId: "test",
		Project: nil,
	})

	if err != nil {
		t.Fatal(err)
	}

	createdItem, err := test.Client.TodoItem.Query().First(ctx)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, createdItem.Title, "new task")
}

func TestItemDB_CreateItem_WithProject(t *testing.T) {
	test := beforeEach(t)
	ctx := context.Background()

	project := test.Client.Project.Create().SetOwnerID("test").SetName("project").SaveX(ctx)

	todoRepo := todoinfra.ItemDB{ItemClient: test.Client.TodoItem, ProjectClient: test.Client.Project}
	err := todoRepo.CreateItem(ctx, &todoitem.Item{
		UUId:    uuid.UUID{},
		Title:   "new task",
		OwnerId: "test",
		Project: &todoitem.Project{
			UUId:    project.UUId,
			Name:    "",
			OwnerId: "",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	createdItem := test.Client.TodoItem.Query().FirstX(ctx)
	assert.Equal(t, createdItem.Title, "new task")
	assert.Equal(t, createdItem.QueryProject().FirstIDX(ctx), project.ID)
}
