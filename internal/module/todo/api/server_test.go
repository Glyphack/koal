package todoapi_test

import (
	"context"
	todoapi "github.com/glyphack/koal/internal/module/todo/api"
	"github.com/glyphack/koal/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"
)

func beforeEach(t *testing.T) *testutils.TestWithDBClient {
	return &testutils.TestWithDBClient{Client: testutils.SetupTestEntClient(t)}
}

func TestGetUserProjectsWithOneProject(t *testing.T) {
	dependencies := beforeEach(t)
	server := todoapi.NewServer(dependencies.Client)
	ownerId := "Sh"
	ctx := testutils.GetAuthenticatedContext(context.Background(), ownerId)
	_, err := dependencies.Client.Project.Create().SetOwnerID(ownerId).SetName("testProj").Save(ctx)
	if err != nil {
		t.Fatal(err)
	}
	response, err := server.GetProjects(ctx, &emptypb.Empty{})
	assert.Nil(t, err)
	assert.Equal(t, len(response.Projects), 1)
}

func TestGetUserProjectsWithMultipleProjects(t *testing.T) {
	dependencies := beforeEach(t)
	server := todoapi.NewServer(dependencies.Client)
	ownerId := "Sh"
	ctx := testutils.GetAuthenticatedContext(context.Background(), ownerId)
	_, err := dependencies.Client.Project.Create().SetOwnerID(ownerId).SetName("testProj1").Save(ctx)
	if err != nil {
		t.Fatal(err)
	}
	_, err = dependencies.Client.Project.Create().SetOwnerID(ownerId).SetName("testProj2").Save(ctx)
	if err != nil {
		t.Fatal(err)
	}
	response, err := server.GetProjects(ctx, &emptypb.Empty{})
	assert.Nil(t, err)
	assert.Equal(t, len(response.Projects), 2)
}
