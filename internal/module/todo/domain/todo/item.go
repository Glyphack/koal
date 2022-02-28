package tododomain

import (
	"errors"
	"github.com/google/uuid"
)

type Item struct {
	UUId    uuid.UUID
	Title   string
	OwnerId string
	Project *Project
	IsDone  bool
}

var PermissionDenied = errors.New("you cannot edit this item")

func (i *Item) UpdateTitle(UserId string, newTitle string) error {
	if ok := isUserAllowedToEditItem(i, UserId); ok == false {
		return PermissionDenied
	}
	i.Title = newTitle
	return nil
}

func (i *Item) UpdateStatus(UserId string, isDone bool) error {
	if ok := isUserAllowedToEditItem(i, UserId); ok == false {
		return PermissionDenied
	}
	i.IsDone = isDone
	return nil
}

func isUserAllowedToEditItem(item *Item, userId string) bool {
	if userId != item.OwnerId {
		return false
	}
	return true
}

func IsUserAllowedToDeleteItem(item *Item, userId string) bool {
	if userId != item.OwnerId {
		return false
	}
	return true
}
