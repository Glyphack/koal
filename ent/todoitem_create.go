// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/glyphack/koal/ent/project"
	"github.com/glyphack/koal/ent/todoitem"
	"github.com/google/uuid"
)

// TodoItemCreate is the builder for creating a TodoItem entity.
type TodoItemCreate struct {
	config
	mutation *TodoItemMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (tic *TodoItemCreate) SetTitle(s string) *TodoItemCreate {
	tic.mutation.SetTitle(s)
	return tic
}

// SetCreatedAt sets the "created_at" field.
func (tic *TodoItemCreate) SetCreatedAt(t time.Time) *TodoItemCreate {
	tic.mutation.SetCreatedAt(t)
	return tic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tic *TodoItemCreate) SetNillableCreatedAt(t *time.Time) *TodoItemCreate {
	if t != nil {
		tic.SetCreatedAt(*t)
	}
	return tic
}

// SetUUID sets the "uuid" field.
func (tic *TodoItemCreate) SetUUID(u uuid.UUID) *TodoItemCreate {
	tic.mutation.SetUUID(u)
	return tic
}

// SetOwnerID sets the "owner_id" field.
func (tic *TodoItemCreate) SetOwnerID(s string) *TodoItemCreate {
	tic.mutation.SetOwnerID(s)
	return tic
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (tic *TodoItemCreate) SetProjectID(id int) *TodoItemCreate {
	tic.mutation.SetProjectID(id)
	return tic
}

// SetNillableProjectID sets the "project" edge to the Project entity by ID if the given value is not nil.
func (tic *TodoItemCreate) SetNillableProjectID(id *int) *TodoItemCreate {
	if id != nil {
		tic = tic.SetProjectID(*id)
	}
	return tic
}

// SetProject sets the "project" edge to the Project entity.
func (tic *TodoItemCreate) SetProject(p *Project) *TodoItemCreate {
	return tic.SetProjectID(p.ID)
}

// Mutation returns the TodoItemMutation object of the builder.
func (tic *TodoItemCreate) Mutation() *TodoItemMutation {
	return tic.mutation
}

// Save creates the TodoItem in the database.
func (tic *TodoItemCreate) Save(ctx context.Context) (*TodoItem, error) {
	var (
		err  error
		node *TodoItem
	)
	tic.defaults()
	if len(tic.hooks) == 0 {
		if err = tic.check(); err != nil {
			return nil, err
		}
		node, err = tic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TodoItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tic.check(); err != nil {
				return nil, err
			}
			tic.mutation = mutation
			if node, err = tic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tic.hooks) - 1; i >= 0; i-- {
			if tic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tic *TodoItemCreate) SaveX(ctx context.Context) *TodoItem {
	v, err := tic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tic *TodoItemCreate) Exec(ctx context.Context) error {
	_, err := tic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tic *TodoItemCreate) ExecX(ctx context.Context) {
	if err := tic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tic *TodoItemCreate) defaults() {
	if _, ok := tic.mutation.CreatedAt(); !ok {
		v := todoitem.DefaultCreatedAt()
		tic.mutation.SetCreatedAt(v)
	}
	if _, ok := tic.mutation.UUID(); !ok {
		v := todoitem.DefaultUUID()
		tic.mutation.SetUUID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tic *TodoItemCreate) check() error {
	if _, ok := tic.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "title"`)}
	}
	if _, ok := tic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := tic.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`ent: missing required field "uuid"`)}
	}
	if _, ok := tic.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner_id", err: errors.New(`ent: missing required field "owner_id"`)}
	}
	return nil
}

func (tic *TodoItemCreate) sqlSave(ctx context.Context) (*TodoItem, error) {
	_node, _spec := tic.createSpec()
	if err := sqlgraph.CreateNode(ctx, tic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tic *TodoItemCreate) createSpec() (*TodoItem, *sqlgraph.CreateSpec) {
	var (
		_node = &TodoItem{config: tic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: todoitem.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: todoitem.FieldID,
			},
		}
	)
	if value, ok := tic.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: todoitem.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := tic.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: todoitem.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := tic.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: todoitem.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := tic.mutation.OwnerID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: todoitem.FieldOwnerID,
		})
		_node.OwnerID = value
	}
	if nodes := tic.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todoitem.ProjectTable,
			Columns: []string{todoitem.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.project_items = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TodoItemCreateBulk is the builder for creating many TodoItem entities in bulk.
type TodoItemCreateBulk struct {
	config
	builders []*TodoItemCreate
}

// Save creates the TodoItem entities in the database.
func (ticb *TodoItemCreateBulk) Save(ctx context.Context) ([]*TodoItem, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ticb.builders))
	nodes := make([]*TodoItem, len(ticb.builders))
	mutators := make([]Mutator, len(ticb.builders))
	for i := range ticb.builders {
		func(i int, root context.Context) {
			builder := ticb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TodoItemMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ticb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ticb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ticb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ticb *TodoItemCreateBulk) SaveX(ctx context.Context) []*TodoItem {
	v, err := ticb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ticb *TodoItemCreateBulk) Exec(ctx context.Context) error {
	_, err := ticb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ticb *TodoItemCreateBulk) ExecX(ctx context.Context) {
	if err := ticb.Exec(ctx); err != nil {
		panic(err)
	}
}
