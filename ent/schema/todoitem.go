package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TodoItem holds the schema definition for the TodoItem entity.
type TodoItem struct {
	ent.Schema
}

// Fields of the TodoItem.
func (TodoItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.Bool("is_done").Default(false),
		field.Time("created_at").
			Default(time.Now),
		field.UUID("uuid", uuid.UUID{}).
			Default(uuid.New),
		field.String("owner_id"),
		field.String("description").Optional(),
	}
}

// Edges of the TodoItem.
func (TodoItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).
			Ref("items").
			Unique(),
	}
}
