package todoinfra

import (
	"context"

	"github.com/glyphack/koal/ent"
	"github.com/glyphack/koal/ent/project"
	"github.com/glyphack/koal/ent/todoitem"
	tododomain "github.com/glyphack/koal/internal/module/todo/domain/todo"
	"github.com/google/uuid"
)

type ItemDB struct {
	ProjectClient *ent.ProjectClient
	ItemClient    *ent.TodoItemClient
}

func (i ItemDB) CreateProject(ctx context.Context, project *tododomain.Project) error {
	err := i.ProjectClient.Create().SetName(project.Name).SetOwnerID(project.OwnerId).SetUUID(project.UUId).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (i ItemDB) GetItemById(ctx context.Context, Id string) (*tododomain.TodoItem, error) {
	itemUUID, err := uuid.Parse(Id)
	if err != nil {
		return nil, err
	}
	dbItem, err := i.ItemClient.Query().Where(todoitem.UUID(itemUUID)).WithProject().First(ctx)
	if err != nil {
		return nil, err
	}
	return &tododomain.TodoItem{
		UUId:    dbItem.UUID,
		Title:   dbItem.Title,
		OwnerId: dbItem.OwnerID,
		IsDone:  dbItem.IsDone,
		Project: &tododomain.Project{
			UUId:    dbItem.Edges.Project.UUID,
			Name:    dbItem.Edges.Project.Name,
			OwnerId: dbItem.Edges.Project.OwnerID,
		},
	}, nil

}
func (i ItemDB) UpdateItem(ctx context.Context, Id string, updatedItem *tododomain.TodoItem) error {
	itemUUID, err := uuid.Parse(Id)
	if err != nil {
		return err
	}
	_, err = i.ItemClient.Update().Where(todoitem.UUID(itemUUID)).SetTitle(updatedItem.Title).SetIsDone(updatedItem.IsDone).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (i ItemDB) AllItems(ctx context.Context, OwnerId string) ([]*tododomain.TodoItem, error) {
	dbItems, err := i.ItemClient.Query().Where(todoitem.OwnerID(OwnerId)).WithProject().All(ctx)
	if err != nil {
		return nil, err
	}
	var items []*tododomain.TodoItem
	for _, item := range dbItems {
		items = append(items, &tododomain.TodoItem{
			UUId:    item.UUID,
			Title:   item.Title,
			OwnerId: item.OwnerID,
			Project: &tododomain.Project{
				UUId:    item.Edges.Project.UUID,
				Name:    item.Edges.Project.Name,
				OwnerId: item.Edges.Project.OwnerID,
			},
			IsDone: item.IsDone,
		})
	}
	return items, nil
}

func (i ItemDB) AllUndoneItems(ctx context.Context, ownerId string) ([]*tododomain.TodoItem, error) {
	dbUndoneItems, err := i.ItemClient.Query().Where(
		todoitem.OwnerID(ownerId),
		todoitem.IsDone(false),
	).All(ctx)
	if err != nil {
		return nil, err
	}
	var items []*tododomain.TodoItem
	for _, dbItem := range dbUndoneItems {
		itemProject, err := dbItem.QueryProject().First(ctx)
		if err != nil {
			return nil, err
		}
		items = append(items, &tododomain.TodoItem{
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
func (i ItemDB) GetProject(ctx context.Context, ID string) (*tododomain.ProjectInfo, error) {
	projectUUID, _ := uuid.Parse(ID)
	dbProject, err := i.ProjectClient.Query().Where(
		project.UUID(projectUUID)).WithItems().First(ctx)
	if err != nil {
		return nil, err
	}
	var items []*tododomain.TodoItem
	dbItems := dbProject.Edges.Items
	domainProject := &tododomain.Project{
		UUId:    dbProject.UUID,
		Name:    dbProject.Name,
		OwnerId: dbProject.OwnerID}
	for _, item := range dbItems {
		items = append(items, &tododomain.TodoItem{
			UUId:    item.UUID,
			Title:   item.Title,
			OwnerId: item.OwnerID,
			Project: domainProject,
		})
	}
	return &tododomain.ProjectInfo{
		Project: domainProject,
		Items:   items,
	}, nil
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

func (i ItemDB) CreateItem(ctx context.Context, newItem *tododomain.TodoItem) error {
	projectId, err := i.ProjectClient.Query().Where(project.UUID(newItem.Project.UUId)).FirstID(ctx)
	if err != nil {
		return err
	}
	err = i.ItemClient.Create().SetOwnerID(newItem.OwnerId).SetTitle(newItem.Title).SetUUID(newItem.UUId).SetProjectID(projectId).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (i ItemDB) DeleteItem(ctx context.Context, ID string) error {
	itemUUID, _ := uuid.Parse(ID)
	_, err := i.ItemClient.Delete().Where(todoitem.UUID(itemUUID)).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (i ItemDB) DeleteProject(ctx context.Context, ID string) error {
	projectUUID, _ := uuid.Parse(ID)
	_, err := i.ProjectClient.Delete().Where(project.UUID(projectUUID)).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (i ItemDB) UpdateProjectById(ctx context.Context, ID string, name string) error {
	projectUUID, _ := uuid.Parse(ID)
	err := i.ProjectClient.Update().SetName(name).Where(project.UUID(projectUUID)).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
