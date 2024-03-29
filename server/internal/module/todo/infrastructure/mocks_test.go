// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package todoinfra

import (
	"context"
	tododomain "github.com/glyphack/koal/internal/module/todo/domain/todo"
	"sync"
)

// Ensure, that TodoRepositoryMock does implement TodoRepository.
// If this is not the case, regenerate this file with moq.
var _ TodoRepository = &TodoRepositoryMock{}

// TodoRepositoryMock is a mock implementation of TodoRepository.
//
// 	func TestSomethingThatUsesTodoRepository(t *testing.T) {
//
// 		// make and configure a mocked TodoRepository
// 		mockedTodoRepository := &TodoRepositoryMock{
// 			AllUndoneItemsFunc: func(ctx context.Context, ownerId string) ([]*tododomain.TodoItem, error) {
// 				panic("mock out the AllUndoneItems method")
// 			},
// 			CreateItemFunc: func(ctx context.Context, newItem *tododomain.TodoItem) error {
// 				panic("mock out the CreateItem method")
// 			},
// 			CreateProjectFunc: func(ctx context.Context, project *tododomain.Project) error {
// 				panic("mock out the CreateProject method")
// 			},
// 			DeleteItemFunc: func(ctx context.Context, ID string) error {
// 				panic("mock out the DeleteItem method")
// 			},
// 			DeleteProjectFunc: func(ctx context.Context, ID string) error {
// 				panic("mock out the DeleteProject method")
// 			},
// 			GetAllMemberProjectsFunc: func(ctx context.Context, OwnerId string) ([]*tododomain.Project, error) {
// 				panic("mock out the GetAllMemberProjects method")
// 			},
// 			GetItemByIdFunc: func(ctx context.Context, Id string) (*tododomain.TodoItem, error) {
// 				panic("mock out the GetItemById method")
// 			},
// 			GetItemsFunc: func(ctx context.Context, itemQuery TodoItemQuery) (*[]tododomain.TodoItem, error) {
// 				panic("mock out the GetItems method")
// 			},
// 			GetProjectFunc: func(ctx context.Context, ID string) (*tododomain.ProjectInfo, error) {
// 				panic("mock out the GetProject method")
// 			},
// 			UpdateItemFunc: func(ctx context.Context, Id string, updatedItem *tododomain.TodoItem) error {
// 				panic("mock out the UpdateItem method")
// 			},
// 			UpdateProjectByIdFunc: func(ctx context.Context, ID string, name string) error {
// 				panic("mock out the UpdateProjectById method")
// 			},
// 		}
//
// 		// use mockedTodoRepository in code that requires TodoRepository
// 		// and then make assertions.
//
// 	}
type TodoRepositoryMock struct {
	// AllUndoneItemsFunc mocks the AllUndoneItems method.
	AllUndoneItemsFunc func(ctx context.Context, ownerId string) ([]*tododomain.TodoItem, error)

	// CreateItemFunc mocks the CreateItem method.
	CreateItemFunc func(ctx context.Context, newItem *tododomain.TodoItem) error

	// CreateProjectFunc mocks the CreateProject method.
	CreateProjectFunc func(ctx context.Context, project *tododomain.Project) error

	// DeleteItemFunc mocks the DeleteItem method.
	DeleteItemFunc func(ctx context.Context, ID string) error

	// DeleteProjectFunc mocks the DeleteProject method.
	DeleteProjectFunc func(ctx context.Context, ID string) error

	// GetAllMemberProjectsFunc mocks the GetAllMemberProjects method.
	GetAllMemberProjectsFunc func(ctx context.Context, OwnerId string) ([]*tododomain.Project, error)

	// GetItemByIdFunc mocks the GetItemById method.
	GetItemByIdFunc func(ctx context.Context, Id string) (*tododomain.TodoItem, error)

	// GetItemsFunc mocks the GetItems method.
	GetItemsFunc func(ctx context.Context, itemQuery TodoItemQuery) (*[]tododomain.TodoItem, error)

	// GetProjectFunc mocks the GetProject method.
	GetProjectFunc func(ctx context.Context, ID string) (*tododomain.ProjectInfo, error)

	// UpdateItemFunc mocks the UpdateItem method.
	UpdateItemFunc func(ctx context.Context, Id string, updatedItem *tododomain.TodoItem) error

	// UpdateProjectByIdFunc mocks the UpdateProjectById method.
	UpdateProjectByIdFunc func(ctx context.Context, ID string, name string) error

	// calls tracks calls to the methods.
	calls struct {
		// AllUndoneItems holds details about calls to the AllUndoneItems method.
		AllUndoneItems []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// OwnerId is the ownerId argument value.
			OwnerId string
		}
		// CreateItem holds details about calls to the CreateItem method.
		CreateItem []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// NewItem is the newItem argument value.
			NewItem *tododomain.TodoItem
		}
		// CreateProject holds details about calls to the CreateProject method.
		CreateProject []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Project is the project argument value.
			Project *tododomain.Project
		}
		// DeleteItem holds details about calls to the DeleteItem method.
		DeleteItem []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the ID argument value.
			ID string
		}
		// DeleteProject holds details about calls to the DeleteProject method.
		DeleteProject []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the ID argument value.
			ID string
		}
		// GetAllMemberProjects holds details about calls to the GetAllMemberProjects method.
		GetAllMemberProjects []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// OwnerId is the OwnerId argument value.
			OwnerId string
		}
		// GetItemById holds details about calls to the GetItemById method.
		GetItemById []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the Id argument value.
			ID string
		}
		// GetItems holds details about calls to the GetItems method.
		GetItems []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ItemQuery is the itemQuery argument value.
			ItemQuery TodoItemQuery
		}
		// GetProject holds details about calls to the GetProject method.
		GetProject []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the ID argument value.
			ID string
		}
		// UpdateItem holds details about calls to the UpdateItem method.
		UpdateItem []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the Id argument value.
			ID string
			// UpdatedItem is the updatedItem argument value.
			UpdatedItem *tododomain.TodoItem
		}
		// UpdateProjectById holds details about calls to the UpdateProjectById method.
		UpdateProjectById []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the ID argument value.
			ID string
			// Name is the name argument value.
			Name string
		}
	}
	lockAllUndoneItems       sync.RWMutex
	lockCreateItem           sync.RWMutex
	lockCreateProject        sync.RWMutex
	lockDeleteItem           sync.RWMutex
	lockDeleteProject        sync.RWMutex
	lockGetAllMemberProjects sync.RWMutex
	lockGetItemById          sync.RWMutex
	lockGetItems             sync.RWMutex
	lockGetProject           sync.RWMutex
	lockUpdateItem           sync.RWMutex
	lockUpdateProjectById    sync.RWMutex
}

// AllUndoneItems calls AllUndoneItemsFunc.
func (mock *TodoRepositoryMock) AllUndoneItems(ctx context.Context, ownerId string) ([]*tododomain.TodoItem, error) {
	if mock.AllUndoneItemsFunc == nil {
		panic("TodoRepositoryMock.AllUndoneItemsFunc: method is nil but TodoRepository.AllUndoneItems was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		OwnerId string
	}{
		Ctx:     ctx,
		OwnerId: ownerId,
	}
	mock.lockAllUndoneItems.Lock()
	mock.calls.AllUndoneItems = append(mock.calls.AllUndoneItems, callInfo)
	mock.lockAllUndoneItems.Unlock()
	return mock.AllUndoneItemsFunc(ctx, ownerId)
}

// AllUndoneItemsCalls gets all the calls that were made to AllUndoneItems.
// Check the length with:
//     len(mockedTodoRepository.AllUndoneItemsCalls())
func (mock *TodoRepositoryMock) AllUndoneItemsCalls() []struct {
	Ctx     context.Context
	OwnerId string
} {
	var calls []struct {
		Ctx     context.Context
		OwnerId string
	}
	mock.lockAllUndoneItems.RLock()
	calls = mock.calls.AllUndoneItems
	mock.lockAllUndoneItems.RUnlock()
	return calls
}

// CreateItem calls CreateItemFunc.
func (mock *TodoRepositoryMock) CreateItem(ctx context.Context, newItem *tododomain.TodoItem) error {
	if mock.CreateItemFunc == nil {
		panic("TodoRepositoryMock.CreateItemFunc: method is nil but TodoRepository.CreateItem was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		NewItem *tododomain.TodoItem
	}{
		Ctx:     ctx,
		NewItem: newItem,
	}
	mock.lockCreateItem.Lock()
	mock.calls.CreateItem = append(mock.calls.CreateItem, callInfo)
	mock.lockCreateItem.Unlock()
	return mock.CreateItemFunc(ctx, newItem)
}

// CreateItemCalls gets all the calls that were made to CreateItem.
// Check the length with:
//     len(mockedTodoRepository.CreateItemCalls())
func (mock *TodoRepositoryMock) CreateItemCalls() []struct {
	Ctx     context.Context
	NewItem *tododomain.TodoItem
} {
	var calls []struct {
		Ctx     context.Context
		NewItem *tododomain.TodoItem
	}
	mock.lockCreateItem.RLock()
	calls = mock.calls.CreateItem
	mock.lockCreateItem.RUnlock()
	return calls
}

// CreateProject calls CreateProjectFunc.
func (mock *TodoRepositoryMock) CreateProject(ctx context.Context, project *tododomain.Project) error {
	if mock.CreateProjectFunc == nil {
		panic("TodoRepositoryMock.CreateProjectFunc: method is nil but TodoRepository.CreateProject was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Project *tododomain.Project
	}{
		Ctx:     ctx,
		Project: project,
	}
	mock.lockCreateProject.Lock()
	mock.calls.CreateProject = append(mock.calls.CreateProject, callInfo)
	mock.lockCreateProject.Unlock()
	return mock.CreateProjectFunc(ctx, project)
}

// CreateProjectCalls gets all the calls that were made to CreateProject.
// Check the length with:
//     len(mockedTodoRepository.CreateProjectCalls())
func (mock *TodoRepositoryMock) CreateProjectCalls() []struct {
	Ctx     context.Context
	Project *tododomain.Project
} {
	var calls []struct {
		Ctx     context.Context
		Project *tododomain.Project
	}
	mock.lockCreateProject.RLock()
	calls = mock.calls.CreateProject
	mock.lockCreateProject.RUnlock()
	return calls
}

// DeleteItem calls DeleteItemFunc.
func (mock *TodoRepositoryMock) DeleteItem(ctx context.Context, ID string) error {
	if mock.DeleteItemFunc == nil {
		panic("TodoRepositoryMock.DeleteItemFunc: method is nil but TodoRepository.DeleteItem was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  ID,
	}
	mock.lockDeleteItem.Lock()
	mock.calls.DeleteItem = append(mock.calls.DeleteItem, callInfo)
	mock.lockDeleteItem.Unlock()
	return mock.DeleteItemFunc(ctx, ID)
}

// DeleteItemCalls gets all the calls that were made to DeleteItem.
// Check the length with:
//     len(mockedTodoRepository.DeleteItemCalls())
func (mock *TodoRepositoryMock) DeleteItemCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockDeleteItem.RLock()
	calls = mock.calls.DeleteItem
	mock.lockDeleteItem.RUnlock()
	return calls
}

// DeleteProject calls DeleteProjectFunc.
func (mock *TodoRepositoryMock) DeleteProject(ctx context.Context, ID string) error {
	if mock.DeleteProjectFunc == nil {
		panic("TodoRepositoryMock.DeleteProjectFunc: method is nil but TodoRepository.DeleteProject was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  ID,
	}
	mock.lockDeleteProject.Lock()
	mock.calls.DeleteProject = append(mock.calls.DeleteProject, callInfo)
	mock.lockDeleteProject.Unlock()
	return mock.DeleteProjectFunc(ctx, ID)
}

// DeleteProjectCalls gets all the calls that were made to DeleteProject.
// Check the length with:
//     len(mockedTodoRepository.DeleteProjectCalls())
func (mock *TodoRepositoryMock) DeleteProjectCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockDeleteProject.RLock()
	calls = mock.calls.DeleteProject
	mock.lockDeleteProject.RUnlock()
	return calls
}

// GetAllMemberProjects calls GetAllMemberProjectsFunc.
func (mock *TodoRepositoryMock) GetAllMemberProjects(ctx context.Context, OwnerId string) ([]*tododomain.Project, error) {
	if mock.GetAllMemberProjectsFunc == nil {
		panic("TodoRepositoryMock.GetAllMemberProjectsFunc: method is nil but TodoRepository.GetAllMemberProjects was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		OwnerId string
	}{
		Ctx:     ctx,
		OwnerId: OwnerId,
	}
	mock.lockGetAllMemberProjects.Lock()
	mock.calls.GetAllMemberProjects = append(mock.calls.GetAllMemberProjects, callInfo)
	mock.lockGetAllMemberProjects.Unlock()
	return mock.GetAllMemberProjectsFunc(ctx, OwnerId)
}

// GetAllMemberProjectsCalls gets all the calls that were made to GetAllMemberProjects.
// Check the length with:
//     len(mockedTodoRepository.GetAllMemberProjectsCalls())
func (mock *TodoRepositoryMock) GetAllMemberProjectsCalls() []struct {
	Ctx     context.Context
	OwnerId string
} {
	var calls []struct {
		Ctx     context.Context
		OwnerId string
	}
	mock.lockGetAllMemberProjects.RLock()
	calls = mock.calls.GetAllMemberProjects
	mock.lockGetAllMemberProjects.RUnlock()
	return calls
}

// GetItemById calls GetItemByIdFunc.
func (mock *TodoRepositoryMock) GetItemById(ctx context.Context, Id string) (*tododomain.TodoItem, error) {
	if mock.GetItemByIdFunc == nil {
		panic("TodoRepositoryMock.GetItemByIdFunc: method is nil but TodoRepository.GetItemById was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  Id,
	}
	mock.lockGetItemById.Lock()
	mock.calls.GetItemById = append(mock.calls.GetItemById, callInfo)
	mock.lockGetItemById.Unlock()
	return mock.GetItemByIdFunc(ctx, Id)
}

// GetItemByIdCalls gets all the calls that were made to GetItemById.
// Check the length with:
//     len(mockedTodoRepository.GetItemByIdCalls())
func (mock *TodoRepositoryMock) GetItemByIdCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockGetItemById.RLock()
	calls = mock.calls.GetItemById
	mock.lockGetItemById.RUnlock()
	return calls
}

// GetItems calls GetItemsFunc.
func (mock *TodoRepositoryMock) GetItems(ctx context.Context, itemQuery TodoItemQuery) (*[]tododomain.TodoItem, error) {
	if mock.GetItemsFunc == nil {
		panic("TodoRepositoryMock.GetItemsFunc: method is nil but TodoRepository.GetItems was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		ItemQuery TodoItemQuery
	}{
		Ctx:       ctx,
		ItemQuery: itemQuery,
	}
	mock.lockGetItems.Lock()
	mock.calls.GetItems = append(mock.calls.GetItems, callInfo)
	mock.lockGetItems.Unlock()
	return mock.GetItemsFunc(ctx, itemQuery)
}

// GetItemsCalls gets all the calls that were made to GetItems.
// Check the length with:
//     len(mockedTodoRepository.GetItemsCalls())
func (mock *TodoRepositoryMock) GetItemsCalls() []struct {
	Ctx       context.Context
	ItemQuery TodoItemQuery
} {
	var calls []struct {
		Ctx       context.Context
		ItemQuery TodoItemQuery
	}
	mock.lockGetItems.RLock()
	calls = mock.calls.GetItems
	mock.lockGetItems.RUnlock()
	return calls
}

// GetProject calls GetProjectFunc.
func (mock *TodoRepositoryMock) GetProject(ctx context.Context, ID string) (*tododomain.ProjectInfo, error) {
	if mock.GetProjectFunc == nil {
		panic("TodoRepositoryMock.GetProjectFunc: method is nil but TodoRepository.GetProject was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  ID,
	}
	mock.lockGetProject.Lock()
	mock.calls.GetProject = append(mock.calls.GetProject, callInfo)
	mock.lockGetProject.Unlock()
	return mock.GetProjectFunc(ctx, ID)
}

// GetProjectCalls gets all the calls that were made to GetProject.
// Check the length with:
//     len(mockedTodoRepository.GetProjectCalls())
func (mock *TodoRepositoryMock) GetProjectCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockGetProject.RLock()
	calls = mock.calls.GetProject
	mock.lockGetProject.RUnlock()
	return calls
}

// UpdateItem calls UpdateItemFunc.
func (mock *TodoRepositoryMock) UpdateItem(ctx context.Context, Id string, updatedItem *tododomain.TodoItem) error {
	if mock.UpdateItemFunc == nil {
		panic("TodoRepositoryMock.UpdateItemFunc: method is nil but TodoRepository.UpdateItem was just called")
	}
	callInfo := struct {
		Ctx         context.Context
		ID          string
		UpdatedItem *tododomain.TodoItem
	}{
		Ctx:         ctx,
		ID:          Id,
		UpdatedItem: updatedItem,
	}
	mock.lockUpdateItem.Lock()
	mock.calls.UpdateItem = append(mock.calls.UpdateItem, callInfo)
	mock.lockUpdateItem.Unlock()
	return mock.UpdateItemFunc(ctx, Id, updatedItem)
}

// UpdateItemCalls gets all the calls that were made to UpdateItem.
// Check the length with:
//     len(mockedTodoRepository.UpdateItemCalls())
func (mock *TodoRepositoryMock) UpdateItemCalls() []struct {
	Ctx         context.Context
	ID          string
	UpdatedItem *tododomain.TodoItem
} {
	var calls []struct {
		Ctx         context.Context
		ID          string
		UpdatedItem *tododomain.TodoItem
	}
	mock.lockUpdateItem.RLock()
	calls = mock.calls.UpdateItem
	mock.lockUpdateItem.RUnlock()
	return calls
}

// UpdateProjectById calls UpdateProjectByIdFunc.
func (mock *TodoRepositoryMock) UpdateProjectById(ctx context.Context, ID string, name string) error {
	if mock.UpdateProjectByIdFunc == nil {
		panic("TodoRepositoryMock.UpdateProjectByIdFunc: method is nil but TodoRepository.UpdateProjectById was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		ID   string
		Name string
	}{
		Ctx:  ctx,
		ID:   ID,
		Name: name,
	}
	mock.lockUpdateProjectById.Lock()
	mock.calls.UpdateProjectById = append(mock.calls.UpdateProjectById, callInfo)
	mock.lockUpdateProjectById.Unlock()
	return mock.UpdateProjectByIdFunc(ctx, ID, name)
}

// UpdateProjectByIdCalls gets all the calls that were made to UpdateProjectById.
// Check the length with:
//     len(mockedTodoRepository.UpdateProjectByIdCalls())
func (mock *TodoRepositoryMock) UpdateProjectByIdCalls() []struct {
	Ctx  context.Context
	ID   string
	Name string
} {
	var calls []struct {
		Ctx  context.Context
		ID   string
		Name string
	}
	mock.lockUpdateProjectById.RLock()
	calls = mock.calls.UpdateProjectById
	mock.lockUpdateProjectById.RUnlock()
	return calls
}
