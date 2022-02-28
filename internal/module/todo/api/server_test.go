package todoapi_test

import (
	"context"
	"github.com/glyphack/koal/ent/enttest"
	todov1 "github.com/glyphack/koal/gen/proto/go/todo/v1"
	todoapi "github.com/glyphack/koal/internal/module/todo/api"
	"github.com/glyphack/koal/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"
)

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
	response, err := server.CreateTodoItem(ctx, &todov1.CreateTodoItemRequest{
		Project: &todov1.CreateTodoItemRequest_ProjectId{ProjectId: project.UUID.String()},
		Title:   "new task",
	})
	assert.Nil(t, err)
	assert.Equal(t, client.TodoItem.Query().FirstX(ctx).Title, response.GetCreatedItem().GetTitle())
}
