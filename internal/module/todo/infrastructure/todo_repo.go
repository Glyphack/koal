package todoinfra

import (
	"context"

	tododomain "github.com/glyphack/koal/internal/module/todo/domain/todo"
	_ "github.com/mattn/go-sqlite3"
)

type TodoRepository interface {
	AllItems(ctx context.Context, OwnerId string) ([]*tododomain.TodoItem, error)
	GetItemById(ctx context.Context, Id string) (*tododomain.TodoItem, error)
	AllUndoneItems(ctx context.Context, ownerId string) ([]*tododomain.TodoItem, error)
	CreateItem(ctx context.Context, newItem *tododomain.TodoItem) error
	UpdateItem(ctx context.Context, Id string, updatedItem *tododomain.TodoItem) error
	DeleteItem(ctx context.Context, ID string) error
	GetAllMemberProjects(ctx context.Context, OwnerId string) ([]*tododomain.Project, error)
	CreateProject(ctx context.Context, project *tododomain.Project) error
	GetProject(ctx context.Context, ID string) (*tododomain.ProjectInfo, error)
	DeleteProject(ctx context.Context, ID string) error
	UpdateProjectById(ctx context.Context, ID string, name string) error
}
