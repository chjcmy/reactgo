// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/backand/ent/book"
	"github.com/backand/ent/predicate"
	"github.com/backand/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where adds a new predicate for the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.predicates = append(uu.mutation.predicates, ps...)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePassword(s *string) *UserUpdate {
	if s != nil {
		uu.SetPassword(*s)
	}
	return uu
}

// SetAge sets the "age" field.
func (uu *UserUpdate) SetAge(t time.Time) *UserUpdate {
	uu.mutation.SetAge(t)
	return uu
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAge(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetAge(*t)
	}
	return uu
}

// SetHobby sets the "hobby" field.
func (uu *UserUpdate) SetHobby(s string) *UserUpdate {
	uu.mutation.SetHobby(s)
	return uu
}

// SetNillableHobby sets the "hobby" field if the given value is not nil.
func (uu *UserUpdate) SetNillableHobby(s *string) *UserUpdate {
	if s != nil {
		uu.SetHobby(*s)
	}
	return uu
}

// SetLang sets the "lang" field.
func (uu *UserUpdate) SetLang(s string) *UserUpdate {
	uu.mutation.SetLang(s)
	return uu
}

// SetNillableLang sets the "lang" field if the given value is not nil.
func (uu *UserUpdate) SetNillableLang(s *string) *UserUpdate {
	if s != nil {
		uu.SetLang(*s)
	}
	return uu
}

// SetGithub sets the "github" field.
func (uu *UserUpdate) SetGithub(s string) *UserUpdate {
	uu.mutation.SetGithub(s)
	return uu
}

// SetNillableGithub sets the "github" field if the given value is not nil.
func (uu *UserUpdate) SetNillableGithub(s *string) *UserUpdate {
	if s != nil {
		uu.SetGithub(*s)
	}
	return uu
}

// SetGitlab sets the "gitlab" field.
func (uu *UserUpdate) SetGitlab(s string) *UserUpdate {
	uu.mutation.SetGitlab(s)
	return uu
}

// SetNillableGitlab sets the "gitlab" field if the given value is not nil.
func (uu *UserUpdate) SetNillableGitlab(s *string) *UserUpdate {
	if s != nil {
		uu.SetGitlab(*s)
	}
	return uu
}

// AddWriterIDs adds the "writer" edge to the Book entity by IDs.
func (uu *UserUpdate) AddWriterIDs(ids ...int) *UserUpdate {
	uu.mutation.AddWriterIDs(ids...)
	return uu
}

// AddWriter adds the "writer" edges to the Book entity.
func (uu *UserUpdate) AddWriter(b ...*Book) *UserUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return uu.AddWriterIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearWriter clears all "writer" edges to the Book entity.
func (uu *UserUpdate) ClearWriter() *UserUpdate {
	uu.mutation.ClearWriter()
	return uu
}

// RemoveWriterIDs removes the "writer" edge to Book entities by IDs.
func (uu *UserUpdate) RemoveWriterIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveWriterIDs(ids...)
	return uu
}

// RemoveWriter removes "writer" edges to Book entities.
func (uu *UserUpdate) RemoveWriter(b ...*Book) *UserUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return uu.RemoveWriterIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
	}
	if value, ok := uu.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldAge,
		})
	}
	if value, ok := uu.mutation.Hobby(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldHobby,
		})
	}
	if value, ok := uu.mutation.Lang(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLang,
		})
	}
	if value, ok := uu.mutation.Github(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldGithub,
		})
	}
	if value, ok := uu.mutation.Gitlab(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldGitlab,
		})
	}
	if uu.mutation.WriterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.WriterTable,
			Columns: []string{user.WriterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: book.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedWriterIDs(); len(nodes) > 0 && !uu.mutation.WriterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.WriterTable,
			Columns: []string{user.WriterColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.WriterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.WriterTable,
			Columns: []string{user.WriterColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePassword(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPassword(*s)
	}
	return uuo
}

// SetAge sets the "age" field.
func (uuo *UserUpdateOne) SetAge(t time.Time) *UserUpdateOne {
	uuo.mutation.SetAge(t)
	return uuo
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAge(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetAge(*t)
	}
	return uuo
}

// SetHobby sets the "hobby" field.
func (uuo *UserUpdateOne) SetHobby(s string) *UserUpdateOne {
	uuo.mutation.SetHobby(s)
	return uuo
}

// SetNillableHobby sets the "hobby" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableHobby(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetHobby(*s)
	}
	return uuo
}

// SetLang sets the "lang" field.
func (uuo *UserUpdateOne) SetLang(s string) *UserUpdateOne {
	uuo.mutation.SetLang(s)
	return uuo
}

// SetNillableLang sets the "lang" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLang(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetLang(*s)
	}
	return uuo
}

// SetGithub sets the "github" field.
func (uuo *UserUpdateOne) SetGithub(s string) *UserUpdateOne {
	uuo.mutation.SetGithub(s)
	return uuo
}

// SetNillableGithub sets the "github" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableGithub(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetGithub(*s)
	}
	return uuo
}

// SetGitlab sets the "gitlab" field.
func (uuo *UserUpdateOne) SetGitlab(s string) *UserUpdateOne {
	uuo.mutation.SetGitlab(s)
	return uuo
}

// SetNillableGitlab sets the "gitlab" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableGitlab(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetGitlab(*s)
	}
	return uuo
}

// AddWriterIDs adds the "writer" edge to the Book entity by IDs.
func (uuo *UserUpdateOne) AddWriterIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddWriterIDs(ids...)
	return uuo
}

// AddWriter adds the "writer" edges to the Book entity.
func (uuo *UserUpdateOne) AddWriter(b ...*Book) *UserUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return uuo.AddWriterIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearWriter clears all "writer" edges to the Book entity.
func (uuo *UserUpdateOne) ClearWriter() *UserUpdateOne {
	uuo.mutation.ClearWriter()
	return uuo
}

// RemoveWriterIDs removes the "writer" edge to Book entities by IDs.
func (uuo *UserUpdateOne) RemoveWriterIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveWriterIDs(ids...)
	return uuo
}

// RemoveWriter removes "writer" edges to Book entities.
func (uuo *UserUpdateOne) RemoveWriter(b ...*Book) *UserUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return uuo.RemoveWriterIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing User.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
	}
	if value, ok := uuo.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldAge,
		})
	}
	if value, ok := uuo.mutation.Hobby(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldHobby,
		})
	}
	if value, ok := uuo.mutation.Lang(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLang,
		})
	}
	if value, ok := uuo.mutation.Github(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldGithub,
		})
	}
	if value, ok := uuo.mutation.Gitlab(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldGitlab,
		})
	}
	if uuo.mutation.WriterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.WriterTable,
			Columns: []string{user.WriterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: book.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedWriterIDs(); len(nodes) > 0 && !uuo.mutation.WriterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.WriterTable,
			Columns: []string{user.WriterColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.WriterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.WriterTable,
			Columns: []string{user.WriterColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}