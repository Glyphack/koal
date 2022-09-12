package todousecase

import (
	"context"

	tododomain "github.com/glyphack/koal/internal/module/todo/domain/todo"
	todoinfra "github.com/glyphack/koal/internal/module/todo/infrastructure"
	"github.com/google/uuid"
)

type TodoUseCase struct {
	TodoRepository todoinfra.TodoRepository
}

func (u *TodoUseCase) UpdateItem(
	ctx context.Context,
	itemId string,
	newTitle string,
	isDone bool,
	userId string,
	newDesc string,
) (*tododomain.TodoItem, error) {
	item, err := u.TodoRepository.GetItemById(ctx, itemId)
	if err != nil {
		return nil, err
	}

	if newTitle != "" {
		err = item.UpdateTitle(userId, newTitle)
		if err != nil {
			return nil, err
		}
	}
	err = item.UpdateStatus(userId, isDone)
	if err != nil {
		return nil, err
	}

	err = item.UpdateDescription(userId, newDesc)
	if err != nil {
		return nil, err
	}
	err = u.TodoRepository.UpdateItem(ctx, itemId, item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (u *TodoUseCase) DeleteItem(ctx context.Context, itemId string, userId string) error {
	item, err := u.TodoRepository.GetItemById(ctx, itemId)
	if err != nil {
		return err
	}
	ok := tododomain.IsUserAllowedToDeleteItem(item, userId)
	if !ok {
		return PermissionDenied
	}
	err = u.TodoRepository.DeleteItem(ctx, itemId)
	if err != nil {
		return err
	}
	return nil
}

func (u TodoUseCase) CreateProject(
	ctx context.Context,
	userId string,
	projectName string,
) (*tododomain.Project, error) {
	project := &tododomain.Project{
		UUId:    uuid.New(),
		Name:    projectName,
		OwnerId: userId,
	}
	err := u.TodoRepository.CreateProject(ctx, project)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (u TodoUseCase) GetProject(
	ctx context.Context,
	userId string,
	projectId string,
) (*tododomain.ProjectInfo, error) {
	dbProject, err := u.TodoRepository.GetProject(ctx, projectId)
	if err != nil {
		return nil, err
	}
	if dbProject.Project.OwnerId != userId {
		return nil, PermissionDenied
	}
	return dbProject, nil
}

func (u TodoUseCase) UpdateProject(
	ctx context.Context,
	userId string,
	projectId string,
	name string,
) (*tododomain.Project, error) {
	projectInfo, err := u.TodoRepository.GetProject(ctx, projectId)
	if err != nil {
		return nil, err
	}
	if projectInfo.Project.OwnerId != userId {
		return nil, PermissionDenied
	}
	newProject := projectInfo.Project
	newProject.Name = name
	err = u.TodoRepository.UpdateProjectById(ctx, projectId, name)
	if err != nil {
		return nil, err
	}
	return newProject, nil
}

func (u TodoUseCase) DeleteProject(ctx context.Context, userId string, projectId string) error {
	projectInfo, err := u.TodoRepository.GetProject(ctx, projectId)
	if err != nil {
		return err
	}
	if projectInfo.Project.OwnerId != userId {
		return PermissionDenied
	}
	err = u.TodoRepository.DeleteProject(ctx, projectId)
	if err != nil {
		return err
	}
	return nil
}
