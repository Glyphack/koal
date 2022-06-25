// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ProjectsColumns holds the columns for the "projects" table.
	ProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "uuid", Type: field.TypeUUID},
		{Name: "owner_id", Type: field.TypeString},
	}
	// ProjectsTable holds the schema information for the "projects" table.
	ProjectsTable = &schema.Table{
		Name:       "projects",
		Columns:    ProjectsColumns,
		PrimaryKey: []*schema.Column{ProjectsColumns[0]},
	}
	// TodoItemsColumns holds the columns for the "todo_items" table.
	TodoItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "is_done", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "uuid", Type: field.TypeUUID},
		{Name: "owner_id", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "project_items", Type: field.TypeInt, Nullable: true},
	}
	// TodoItemsTable holds the schema information for the "todo_items" table.
	TodoItemsTable = &schema.Table{
		Name:       "todo_items",
		Columns:    TodoItemsColumns,
		PrimaryKey: []*schema.Column{TodoItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "todo_items_projects_items",
				Columns:    []*schema.Column{TodoItemsColumns[7]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "uuid", Type: field.TypeUUID},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ProjectsTable,
		TodoItemsTable,
		UsersTable,
	}
)

func init() {
	TodoItemsTable.ForeignKeys[0].RefTable = ProjectsTable
}
