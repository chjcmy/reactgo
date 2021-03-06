// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/backand/ent/book"
	"github.com/backand/ent/unit"
)

// UnitCreate is the builder for creating a Unit entity.
type UnitCreate struct {
	config
	mutation *UnitMutation
	hooks    []Hook
}

// SetContent sets the "content" field.
func (uc *UnitCreate) SetContent(s string) *UnitCreate {
	uc.mutation.SetContent(s)
	return uc
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (uc *UnitCreate) SetNillableContent(s *string) *UnitCreate {
	if s != nil {
		uc.SetContent(*s)
	}
	return uc
}

// SetContentName sets the "content_name" field.
func (uc *UnitCreate) SetContentName(s string) *UnitCreate {
	uc.mutation.SetContentName(s)
	return uc
}

// SetNillableContentName sets the "content_name" field if the given value is not nil.
func (uc *UnitCreate) SetNillableContentName(s *string) *UnitCreate {
	if s != nil {
		uc.SetContentName(*s)
	}
	return uc
}

// AddContentIDs adds the "contents" edge to the Book entity by IDs.
func (uc *UnitCreate) AddContentIDs(ids ...int) *UnitCreate {
	uc.mutation.AddContentIDs(ids...)
	return uc
}

// AddContents adds the "contents" edges to the Book entity.
func (uc *UnitCreate) AddContents(b ...*Book) *UnitCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return uc.AddContentIDs(ids...)
}

// Mutation returns the UnitMutation object of the builder.
func (uc *UnitCreate) Mutation() *UnitMutation {
	return uc.mutation
}

// Save creates the Unit in the database.
func (uc *UnitCreate) Save(ctx context.Context) (*Unit, error) {
	var (
		err  error
		node *Unit
	)
	uc.defaults()
	if len(uc.hooks) == 0 {
		if err = uc.check(); err != nil {
			return nil, err
		}
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UnitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uc.check(); err != nil {
				return nil, err
			}
			uc.mutation = mutation
			if node, err = uc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			if uc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UnitCreate) SaveX(ctx context.Context) *Unit {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UnitCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UnitCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UnitCreate) defaults() {
	if _, ok := uc.mutation.Content(); !ok {
		v := unit.DefaultContent
		uc.mutation.SetContent(v)
	}
	if _, ok := uc.mutation.ContentName(); !ok {
		v := unit.DefaultContentName
		uc.mutation.SetContentName(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UnitCreate) check() error {
	if _, ok := uc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "content"`)}
	}
	if _, ok := uc.mutation.ContentName(); !ok {
		return &ValidationError{Name: "content_name", err: errors.New(`ent: missing required field "content_name"`)}
	}
	return nil
}

func (uc *UnitCreate) sqlSave(ctx context.Context) (*Unit, error) {
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (uc *UnitCreate) createSpec() (*Unit, *sqlgraph.CreateSpec) {
	var (
		_node = &Unit{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: unit.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: unit.FieldID,
			},
		}
	)
	if value, ok := uc.mutation.Content(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unit.FieldContent,
		})
		_node.Content = value
	}
	if value, ok := uc.mutation.ContentName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unit.FieldContentName,
		})
		_node.ContentName = value
	}
	if nodes := uc.mutation.ContentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   unit.ContentsTable,
			Columns: []string{unit.ContentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: book.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UnitCreateBulk is the builder for creating many Unit entities in bulk.
type UnitCreateBulk struct {
	config
	builders []*UnitCreate
}

// Save creates the Unit entities in the database.
func (ucb *UnitCreateBulk) Save(ctx context.Context) ([]*Unit, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*Unit, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UnitMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UnitCreateBulk) SaveX(ctx context.Context) []*Unit {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UnitCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UnitCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
