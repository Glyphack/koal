package todousecase

import (
	"context"
	"errors"
	todoitem "github.com/glyphack/koal/internal/module/todo/domain/todo"
	todoinfra "github.com/glyphack/koal/internal/module/todo/infrastructure"
	"github.com/google/uuid"
)

type UseCase struct {
	ItemRepository todoinfra.TodoRepository
}

var ItemDoesNotExist = errors.New("item does not exist")
var PermissionDenied = errors.New("not Allowed to delete item")
var ProjectDoesNotExist = errors.New("cannot find project")

func (u *UseCase) DeleteItem(ctx context.Context, itemId string, userId string) error {
	item, err := u.ItemRepository.GetItemById(ctx, itemId)
	if err != nil {
		return ItemDoesNotExist
	}
	ok := todoitem.IsUserAllowedToDeleteItem(item, userId)
	if !ok {
		return PermissionDenied
	}
	return nil
}

func (u UseCase) CreateProject(ctx context.Context, userId string, projectName string) (*todoitem.Project, error) {
	project := &todoitem.Project{
		UUId:    uuid.New(),
		Name:    projectName,
		OwnerId: userId,
	}
	err := u.ItemRepository.CreateProject(ctx, project)
	if err != nil {
		return nil, errors.New("cannot create Item try again later")
	}
	return project, nil
}
func (u UseCase) GetProject(ctx context.Context, userId string, projectId string) (*todoitem.ProjectInfo, error) {
	dbProject, err := u.ItemRepository.GetProject(ctx, projectId)
	if err != nil {

		return nil, ProjectDoesNotExist
	}
	if dbProject.Project.OwnerId != userId {
		return nil, PermissionDenied
	}
	return dbProject, nil
}
