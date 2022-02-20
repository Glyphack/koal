package todoitem

import "github.com/google/uuid"

type Item struct {
	UUId    uuid.UUID
	Title   string
	OwnerId string
	Project *Project
}
