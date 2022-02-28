package tododomain

import "github.com/google/uuid"

type Project struct {
	UUId    uuid.UUID
	Name    string
	OwnerId string
}
