package todoinfra

import (
	"context"
	"github.com/glyphack/koal/ent"
	"github.com/glyphack/koal/ent/project"
	todoitem "github.com/glyphack/koal/internal/module/todo/domain/todo"
)

type ItemDB struct {
	ProjectClient *ent.ProjectClient
	ItemClient    *ent.TodoItemClient
}

func (i ItemDB) AllItems(ctx context.Context, OwnerId string) ([]*todoitem.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (i ItemDB) AssignItemToProject(ctx context.Context, ProjectId string) {
	//TODO implement me
	panic("implement me")
}

func (i ItemDB) GetProject(ctx context.Context, ID string) (*todoitem.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (i ItemDB) GetAllMemberProjects(ctx context.Context, OwnerId string) ([]*todoitem.Project, error) {
	dbProjects, err := i.ProjectClient.Query().Where(project.OwnerID(OwnerId)).All(ctx)
	if err != nil {
		return nil, err
	}
	var projects []*todoitem.Project
	for _, dbProject := range dbProjects {
		projects = append(projects, &todoitem.Project{
			Name:    dbProject.Name,
			OwnerId: dbProject.OwnerID,
			UUId:    dbProject.UUID,
		})
	}
	return projects, nil
}

func (i ItemDB) CreateItem(ctx context.Context, newItem *todoitem.Item) error {
	createItemQuery := i.ItemClient.Create().SetOwnerID(newItem.OwnerId).SetTitle(newItem.Title)
	if newItem.Project != nil {
		projectId, err := i.ProjectClient.Query().Where(project.UUID(newItem.Project.UUId)).FirstID(ctx)
		if err != nil {
			return err
		}
		createItemQuery.SetProjectID(projectId)
	}

	err := createItemQuery.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (i ItemDB) DeleteItem(ctx context.Context, ID string) error {
	//TODO implement me
	panic("implement me")
}

func (i ItemDB) CreateProject(ctx context.Context, name string) error {
	//TODO implement me
	panic("implement me")
}

func (i ItemDB) DeleteProject(ctx context.Context, ID string) error {
	//TODO implement me
	panic("implement me")
}

func (i ItemDB) UpdateProjectById(ctx context.Context, ID string, name string) error {
	//TODO implement me
	panic("implement me")
}
