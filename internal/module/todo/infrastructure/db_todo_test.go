package todoinfra_test

import (
	"context"
	"github.com/glyphack/koal/ent/enttest"
	todoitem "github.com/glyphack/koal/internal/module/todo/domain/todo"
	todoinfra "github.com/glyphack/koal/internal/module/todo/infrastructure"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
