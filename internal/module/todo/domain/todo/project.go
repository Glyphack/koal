package tododomain

import "github.com/google/uuid"

// Project is a value object
type Project struct {
	UUId    uuid.UUID
	Name    string
	OwnerId string
}
