package todoinfra

import (
	"context"
	_ "github.com/mattn/go-sqlite3"

	todoitem "github.com/glyphack/koal/internal/module/todo/domain/todo"
)

type TodoRepository interface {
	AllItems(ctx context.Context, OwnerId string) ([]*todoitem.Project, error)
	CreateItem(ctx context.Context, newItem *todoitem.Item) error
	DeleteItem(ctx context.Context, ID string) error
	AssignItemToProject(ctx context.Context, ProjectId string)
	AllProjects(ctx context.Context, OwnerId string) ([]*todoitem.Project, error)
	CreateProject(ctx context.Context, name string) error
	GetProject(ctx context.Context, ID string) (*todoitem.Project, error)
	DeleteProject(ctx context.Context, ID string) error
	UpdateProjectById(ctx context.Context, ID string, name string) error
}
