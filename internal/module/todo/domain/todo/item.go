package tododomain

import (
	"errors"

	"github.com/google/uuid"
)

type TodoItem struct {
	UUId        uuid.UUID
	Title       string
	OwnerId     string
	Project     *Project
	IsDone      bool
	Description string // html formatted description text
}

var PermissionDenied = errors.New("you cannot edit this item")

func (i *TodoItem) UpdateTitle(userId string, newTitle string) error {
	if ok := isUserAllowedToEditItem(i, userId); ok == false {
		return PermissionDenied
	}
	i.Title = newTitle
	return nil
}

func (i *TodoItem) UpdateStatus(userId string, isDone bool) error {
	if ok := isUserAllowedToEditItem(i, userId); ok == false {
		return PermissionDenied
	}
	i.IsDone = isDone
	return nil
}

func (i *TodoItem) UpdateDescription(userId string, newDesc string) error {
	if ok := isUserAllowedToEditItem(i, userId); ok == false {
		return PermissionDenied
	}
	i.Description = newDesc
	return nil
}

func isUserAllowedToEditItem(item *TodoItem, userId string) bool {
	if userId != item.OwnerId {
		return false
	}
	return true
}

func IsUserAllowedToDeleteItem(item *TodoItem, userId string) bool {
	if userId != item.OwnerId {
		return false
	}
	return true
}
