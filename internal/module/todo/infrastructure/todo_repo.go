package todoinfra

import (
	"context"
	_ "github.com/mattn/go-sqlite3"

	"github.com/glyphack/koal/internal/module/todo/domain/todo"
)

type TodoRepository interface {
	AllItems(ctx context.Context, OwnerId string) ([]*tododomain.Item, error)
	AllUndoneItems(ctx context.Context, ownerId string) ([]*tododomain.Item, error)
	CreateItem(ctx context.Context, newItem *tododomain.Item) error
	DeleteItem(ctx context.Context, ID string) error
	AssignItemToProject(ctx context.Context, ProjectId string)
	GetAllMemberProjects(ctx context.Context, OwnerId string) ([]*tododomain.Project, error)
	CreateProject(ctx context.Context, name string) error
	GetProject(ctx context.Context, ID string) (*tododomain.Project, error)
	DeleteProject(ctx context.Context, ID string) error
	UpdateProjectById(ctx context.Context, ID string, name string) error
}
