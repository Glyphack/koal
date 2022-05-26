package todoinfra

import (
	"context"
	"fmt"

	"github.com/glyphack/koal/ent"
	"github.com/glyphack/koal/ent/project"
	"github.com/glyphack/koal/ent/todoitem"
	tododomain "github.com/glyphack/koal/internal/module/todo/domain/todo"
	"github.com/glyphack/koal/pkg/entutils"
	"github.com/google/uuid"
)

type ItemDB struct {
	ProjectClient *ent.ProjectClient
	ItemClient    *ent.TodoItemClient
	Client        *ent.Client
}

func (i ItemDB) CreateItem(ctx context.Context, newItem *tododomain.TodoItem) error {
	if newItem.Project == nil {
		err := i.ItemClient.Create().SetOwnerID(newItem.OwnerId).SetTitle(newItem.Title).SetUUID(newItem.UUId).Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	}
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
	if ent.IsNotFound(err) {
		return fmt.Errorf("%w", NotFoundErr)
	}
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
	if ent.IsNotFound(err) {
		return nil, NotFoundErr
	}
	if err != nil {
		return nil, err
	}
	project, projectNotExist := dbItem.Edges.ProjectOrErr()
	if projectNotExist != nil {
		return &tododomain.TodoItem{
			UUId:    dbItem.UUID,
			Title:   dbItem.Title,
			OwnerId: dbItem.OwnerID,
			IsDone:  dbItem.IsDone,
			Project: nil,
		}, nil

	}
	return &tododomain.TodoItem{
		UUId:    dbItem.UUID,
		Title:   dbItem.Title,
		OwnerId: dbItem.OwnerID,
		IsDone:  dbItem.IsDone,
		Project: &tododomain.Project{
			UUId:    project.UUID,
			Name:    project.Name,
			OwnerId: project.OwnerID,
		},
	}, nil

}

func (i ItemDB) UpdateItem(ctx context.Context, Id string, updatedItem *tododomain.TodoItem) error {
	itemUUID, err := uuid.Parse(Id)
	if err != nil {
		return err
	}
	_, err = i.ItemClient.Update().Where(todoitem.UUID(itemUUID)).SetTitle(updatedItem.Title).SetIsDone(updatedItem.IsDone).Save(ctx)
	if ent.IsNotFound(err) {
		return fmt.Errorf("%w", NotFoundErr)
	}
	if err != nil {
		return err
	}
	return nil
}

func (i ItemDB) AllUndoneItems(ctx context.Context, ownerId string) ([]*tododomain.TodoItem, error) {
	dbUndoneItems, err := i.ItemClient.Query().Where(
		todoitem.OwnerID(ownerId),
		todoitem.IsDone(false),
	).Order(ent.Asc(todoitem.FieldIsDone), ent.Desc(todoitem.FieldCreatedAt)).All(ctx)

	if ent.IsNotFound(err) {
		return nil, NotFoundErr
	}
	if err != nil {
		return nil, fmt.Errorf("cannot query undone items: %w", err)
	}
	var items []*tododomain.TodoItem
	for _, dbItem := range dbUndoneItems {
		itemProject, err := dbItem.QueryProject().First(ctx)
		if err != nil {
			return nil, fmt.Errorf("cannot load item project: %w", err) 
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

func (i ItemDB) CreateProject(ctx context.Context, project *tododomain.Project) error {
	err := i.ProjectClient.Create().SetName(project.Name).SetOwnerID(project.OwnerId).SetUUID(project.UUId).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (i ItemDB) GetProject(ctx context.Context, ID string) (*tododomain.ProjectInfo, error) {
	projectUUID, _ := uuid.Parse(ID)
	dbProject, err := i.ProjectClient.Query().Where(
		project.UUID(projectUUID)).WithItems().First(ctx)

	if ent.IsNotFound(err) {
		return nil, NotFoundErr
	}
	if err != nil {
		return nil, err
	}
	var items []*tododomain.TodoItem
	dbItems, err := i.ItemClient.Query().Where(todoitem.HasProjectWith(project.UUID(projectUUID))).Order(
		ent.Asc(todoitem.FieldIsDone), ent.Desc(todoitem.FieldCreatedAt),
	).All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to get project items: %w", err)
	}
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
			IsDone:  item.IsDone,
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

func (i ItemDB) DeleteProject(ctx context.Context, ID string) error {
	projectUUID, _ := uuid.Parse(ID)
	tx, err := i.Client.Tx(ctx)
	if err != nil {
		return err
	}
	txClient := tx.Client()
	_, err = txClient.Project.Delete().Where(project.UUID(projectUUID)).Exec(ctx)
	if ent.IsNotFound(err) {
		return entutils.Rollback(tx, NotFoundErr)
	}
	if err != nil {
		return entutils.Rollback(tx, err)
	}
	_, err = txClient.TodoItem.Delete().Where(todoitem.HasProjectWith(project.UUID(projectUUID))).Exec(ctx)
	if err != nil {
		return entutils.Rollback(tx, err)
	}
	return tx.Commit()
}

func (i ItemDB) UpdateProjectById(ctx context.Context, ID string, name string) error {
	projectUUID, _ := uuid.Parse(ID)
	err := i.ProjectClient.Update().SetName(name).Where(project.UUID(projectUUID)).Exec(ctx)
	if ent.IsNotFound(err) {
		return fmt.Errorf("%w", NotFoundErr)
	}
	if err != nil {
		return err
	}
	return nil
}
