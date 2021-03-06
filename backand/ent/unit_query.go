// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/backand/ent/book"
	"github.com/backand/ent/predicate"
	"github.com/backand/ent/unit"
)

// UnitQuery is the builder for querying Unit entities.
type UnitQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Unit
	// eager-loading edges.
	withContents *BookQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UnitQuery builder.
func (uq *UnitQuery) Where(ps ...predicate.Unit) *UnitQuery {
	uq.predicates = append(uq.predicates, ps...)
	return uq
}

// Limit adds a limit step to the query.
func (uq *UnitQuery) Limit(limit int) *UnitQuery {
	uq.limit = &limit
	return uq
}

// Offset adds an offset step to the query.
func (uq *UnitQuery) Offset(offset int) *UnitQuery {
	uq.offset = &offset
	return uq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uq *UnitQuery) Unique(unique bool) *UnitQuery {
	uq.unique = &unique
	return uq
}

// Order adds an order step to the query.
func (uq *UnitQuery) Order(o ...OrderFunc) *UnitQuery {
	uq.order = append(uq.order, o...)
	return uq
}

// QueryContents chains the current query on the "contents" edge.
func (uq *UnitQuery) QueryContents() *BookQuery {
	query := &BookQuery{config: uq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(unit.Table, unit.FieldID, selector),
			sqlgraph.To(book.Table, book.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, unit.ContentsTable, unit.ContentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Unit entity from the query.
// Returns a *NotFoundError when no Unit was found.
func (uq *UnitQuery) First(ctx context.Context) (*Unit, error) {
	nodes, err := uq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{unit.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uq *UnitQuery) FirstX(ctx context.Context) *Unit {
	node, err := uq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Unit ID from the query.
// Returns a *NotFoundError when no Unit ID was found.
func (uq *UnitQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{unit.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uq *UnitQuery) FirstIDX(ctx context.Context) int {
	id, err := uq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Unit entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Unit entity is not found.
// Returns a *NotFoundError when no Unit entities are found.
func (uq *UnitQuery) Only(ctx context.Context) (*Unit, error) {
	nodes, err := uq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{unit.Label}
	default:
		return nil, &NotSingularError{unit.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uq *UnitQuery) OnlyX(ctx context.Context) *Unit {
	node, err := uq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Unit ID in the query.
// Returns a *NotSingularError when exactly one Unit ID is not found.
// Returns a *NotFoundError when no entities are found.
func (uq *UnitQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{unit.Label}
	default:
		err = &NotSingularError{unit.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uq *UnitQuery) OnlyIDX(ctx context.Context) int {
	id, err := uq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Units.
func (uq *UnitQuery) All(ctx context.Context) ([]*Unit, error) {
	if err := uq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return uq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (uq *UnitQuery) AllX(ctx context.Context) []*Unit {
	nodes, err := uq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Unit IDs.
func (uq *UnitQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := uq.Select(unit.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uq *UnitQuery) IDsX(ctx context.Context) []int {
	ids, err := uq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uq *UnitQuery) Count(ctx context.Context) (int, error) {
	if err := uq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return uq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (uq *UnitQuery) CountX(ctx context.Context) int {
	count, err := uq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uq *UnitQuery) Exist(ctx context.Context) (bool, error) {
	if err := uq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return uq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (uq *UnitQuery) ExistX(ctx context.Context) bool {
	exist, err := uq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UnitQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uq *UnitQuery) Clone() *UnitQuery {
	if uq == nil {
		return nil
	}
	return &UnitQuery{
		config:       uq.config,
		limit:        uq.limit,
		offset:       uq.offset,
		order:        append([]OrderFunc{}, uq.order...),
		predicates:   append([]predicate.Unit{}, uq.predicates...),
		withContents: uq.withContents.Clone(),
		// clone intermediate query.
		sql:  uq.sql.Clone(),
		path: uq.path,
	}
}

// WithContents tells the query-builder to eager-load the nodes that are connected to
// the "contents" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UnitQuery) WithContents(opts ...func(*BookQuery)) *UnitQuery {
	query := &BookQuery{config: uq.config}
	for _, opt := range opts {
		opt(query)
	}
	uq.withContents = query
	return uq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Content string `json:"content,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Unit.Query().
//		GroupBy(unit.FieldContent).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (uq *UnitQuery) GroupBy(field string, fields ...string) *UnitGroupBy {
	group := &UnitGroupBy{config: uq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return uq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Content string `json:"content,omitempty"`
//	}
//
//	client.Unit.Query().
//		Select(unit.FieldContent).
//		Scan(ctx, &v)
//
func (uq *UnitQuery) Select(fields ...string) *UnitSelect {
	uq.fields = append(uq.fields, fields...)
	return &UnitSelect{UnitQuery: uq}
}

func (uq *UnitQuery) prepareQuery(ctx context.Context) error {
	for _, f := range uq.fields {
		if !unit.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uq.path != nil {
		prev, err := uq.path(ctx)
		if err != nil {
			return err
		}
		uq.sql = prev
	}
	return nil
}

func (uq *UnitQuery) sqlAll(ctx context.Context) ([]*Unit, error) {
	var (
		nodes       = []*Unit{}
		_spec       = uq.querySpec()
		loadedTypes = [1]bool{
			uq.withContents != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Unit{config: uq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, uq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := uq.withContents; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Unit)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Contents = []*Book{}
		}
		query.withFKs = true
		query.Where(predicate.Book(func(s *sql.Selector) {
			s.Where(sql.InValues(unit.ContentsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.unit_contents
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "unit_contents" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "unit_contents" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Contents = append(node.Edges.Contents, n)
		}
	}

	return nodes, nil
}

func (uq *UnitQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uq.querySpec()
	return sqlgraph.CountNodes(ctx, uq.driver, _spec)
}

func (uq *UnitQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := uq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (uq *UnitQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   unit.Table,
			Columns: unit.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: unit.FieldID,
			},
		},
		From:   uq.sql,
		Unique: true,
	}
	if unique := uq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := uq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, unit.FieldID)
		for i := range fields {
			if fields[i] != unit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uq *UnitQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uq.driver.Dialect())
	t1 := builder.Table(unit.Table)
	columns := uq.fields
	if len(columns) == 0 {
		columns = unit.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uq.sql != nil {
		selector = uq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range uq.predicates {
		p(selector)
	}
	for _, p := range uq.order {
		p(selector)
	}
	if offset := uq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UnitGroupBy is the group-by builder for Unit entities.
type UnitGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ugb *UnitGroupBy) Aggregate(fns ...AggregateFunc) *UnitGroupBy {
	ugb.fns = append(ugb.fns, fns...)
	return ugb
}

// Scan applies the group-by query and scans the result into the given value.
func (ugb *UnitGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ugb.path(ctx)
	if err != nil {
		return err
	}
	ugb.sql = query
	return ugb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ugb *UnitGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ugb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ugb *UnitGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ugb.fields) > 1 {
		return nil, errors.New("ent: UnitGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ugb *UnitGroupBy) StringsX(ctx context.Context) []string {
	v, err := ugb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ugb *UnitGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ugb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{unit.Label}
	default:
		err = fmt.Errorf("ent: UnitGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ugb *UnitGroupBy) StringX(ctx context.Context) string {
	v, err := ugb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ugb *UnitGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ugb.fields) > 1 {
		return nil, errors.New("ent: UnitGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ugb *UnitGroupBy) IntsX(ctx context.Context) []int {
	v, err := ugb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ugb *UnitGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ugb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{unit.Label}
	default:
		err = fmt.Errorf("ent: UnitGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ugb *UnitGroupBy) IntX(ctx context.Context) int {
	v, err := ugb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ugb *UnitGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ugb.fields) > 1 {
		return nil, errors.New("ent: UnitGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ugb *UnitGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ugb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ugb *UnitGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ugb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{unit.Label}
	default:
		err = fmt.Errorf("ent: UnitGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ugb *UnitGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ugb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ugb *UnitGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ugb.fields) > 1 {
		return nil, errors.New("ent: UnitGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ugb *UnitGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ugb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ugb *UnitGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ugb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{unit.Label}
	default:
		err = fmt.Errorf("ent: UnitGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ugb *UnitGroupBy) BoolX(ctx context.Context) bool {
	v, err := ugb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ugb *UnitGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ugb.fields {
		if !unit.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ugb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ugb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ugb *UnitGroupBy) sqlQuery() *sql.Selector {
	selector := ugb.sql.Select()
	aggregation := make([]string, 0, len(ugb.fns))
	for _, fn := range ugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ugb.fields)+len(ugb.fns))
		for _, f := range ugb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ugb.fields...)...)
}

// UnitSelect is the builder for selecting fields of Unit entities.
type UnitSelect struct {
	*UnitQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (us *UnitSelect) Scan(ctx context.Context, v interface{}) error {
	if err := us.prepareQuery(ctx); err != nil {
		return err
	}
	us.sql = us.UnitQuery.sqlQuery(ctx)
	return us.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (us *UnitSelect) ScanX(ctx context.Context, v interface{}) {
	if err := us.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (us *UnitSelect) Strings(ctx context.Context) ([]string, error) {
	if len(us.fields) > 1 {
		return nil, errors.New("ent: UnitSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := us.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (us *UnitSelect) StringsX(ctx context.Context) []string {
	v, err := us.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (us *UnitSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = us.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{unit.Label}
	default:
		err = fmt.Errorf("ent: UnitSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (us *UnitSelect) StringX(ctx context.Context) string {
	v, err := us.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (us *UnitSelect) Ints(ctx context.Context) ([]int, error) {
	if len(us.fields) > 1 {
		return nil, errors.New("ent: UnitSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := us.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (us *UnitSelect) IntsX(ctx context.Context) []int {
	v, err := us.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (us *UnitSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = us.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{unit.Label}
	default:
		err = fmt.Errorf("ent: UnitSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (us *UnitSelect) IntX(ctx context.Context) int {
	v, err := us.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (us *UnitSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(us.fields) > 1 {
		return nil, errors.New("ent: UnitSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := us.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (us *UnitSelect) Float64sX(ctx context.Context) []float64 {
	v, err := us.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (us *UnitSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = us.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{unit.Label}
	default:
		err = fmt.Errorf("ent: UnitSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (us *UnitSelect) Float64X(ctx context.Context) float64 {
	v, err := us.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (us *UnitSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(us.fields) > 1 {
		return nil, errors.New("ent: UnitSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := us.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (us *UnitSelect) BoolsX(ctx context.Context) []bool {
	v, err := us.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (us *UnitSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = us.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{unit.Label}
	default:
		err = fmt.Errorf("ent: UnitSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (us *UnitSelect) BoolX(ctx context.Context) bool {
	v, err := us.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (us *UnitSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := us.sql.Query()
	if err := us.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
