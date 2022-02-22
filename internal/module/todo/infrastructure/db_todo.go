package todoinfra

import (
	"context"
	"github.com/glyphack/koal/ent"
	"github.com/glyphack/koal/ent/project"
	"github.com/glyphack/koal/ent/todoitem"
	"github.com/glyphack/koal/internal/module/todo/domain/todo"
)

type ItemDB struct {
	ProjectClient *ent.ProjectClient
	ItemClient    *ent.TodoItemClient
}

func (i ItemDB) AllItems(ctx context.Context, OwnerId string) ([]*tododomain.Item, error) {
	//TODO implement me
	panic("implement me")
}

func (i ItemDB) AllUndoneItems(ctx context.Context, ownerId string) ([]*tododomain.Item, error) {
	dbUndoneItems, err := i.ItemClient.Query().Where(
		todoitem.OwnerID(ownerId),
		todoitem.IsDone(false),
	).All(ctx)
	if err != nil {
		return nil, err
	}
	var items []*tododomain.Item
	for _, dbItem := range dbUndoneItems {
		itemProject, err := dbItem.QueryProject().First(ctx)
		if err != nil {
			return nil, err
		}
		items = append(items, &tododomain.Item{
			Title:   dbItem.Title,
			OwnerId: dbItem.OwnerID,
			UUId:    dbItem.UUID,
			Project: &tododomain.Project{
				UUId:    itemProject.UUID,
				Name:    itemProject.Name,
				OwnerId: itemProject.OwnerID,
			},
		})
	}
	return items, nil
}

func (i ItemDB) AssignItemToProject(ctx context.Context, ProjectId string) {
	//TODO implement me
	panic("implement me")
}

func (i ItemDB) GetProject(ctx context.Context, ID string) (*tododomain.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (i ItemDB) GetAllMemberProjects(ctx context.Context, OwnerId string) ([]*tododomain.Project, error) {
	dbProjects, err := i.ProjectClient.Query().Where(project.OwnerID(OwnerId)).All(ctx)
	if err != nil {
		return nil, err
	}
	var projects []*tododomain.Project
	for _, dbProject := range dbProjects {
		projects = append(projects, &tododomain.Project{
			Name:    dbProject.Name,
			OwnerId: dbProject.OwnerID,
			UUId:    dbProject.UUID,
		})
	}
	return projects, nil
}

func (i ItemDB) CreateItem(ctx context.Context, newItem *tododomain.Item) error {
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
