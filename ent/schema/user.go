package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique(),
		field.String("password"),
		field.Time("created_at").
			Default(time.Now),
		field.UUID("uuid", uuid.UUID{}).
			Default(uuid.New),
	}
}

func (User) Index() []ent.Index {
	return []ent.Index{
		index.Fields("email").Unique(),
	}
}
