// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/backand/ent/unit"
)

// Unit is the model entity for the Unit schema.
type Unit struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// ContentName holds the value of the "content_name" field.
	ContentName string `json:"content_name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UnitQuery when eager-loading is set.
	Edges UnitEdges `json:"edges"`
}

// UnitEdges holds the relations/edges for other nodes in the graph.
type UnitEdges struct {
	// Contents holds the value of the contents edge.
	Contents []*Book `json:"contents,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ContentsOrErr returns the Contents value or an error if the edge
// was not loaded in eager-loading.
func (e UnitEdges) ContentsOrErr() ([]*Book, error) {
	if e.loadedTypes[0] {
		return e.Contents, nil
	}
	return nil, &NotLoadedError{edge: "contents"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Unit) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case unit.FieldID:
			values[i] = new(sql.NullInt64)
		case unit.FieldContent, unit.FieldContentName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Unit", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Unit fields.
func (u *Unit) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case unit.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case unit.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				u.Content = value.String
			}
		case unit.FieldContentName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content_name", values[i])
			} else if value.Valid {
				u.ContentName = value.String
			}
		}
	}
	return nil
}

// QueryContents queries the "contents" edge of the Unit entity.
func (u *Unit) QueryContents() *BookQuery {
	return (&UnitClient{config: u.config}).QueryContents(u)
}

// Update returns a builder for updating this Unit.
// Note that you need to call Unit.Unwrap() before calling this method if this Unit
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *Unit) Update() *UnitUpdateOne {
	return (&UnitClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the Unit entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *Unit) Unwrap() *Unit {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: Unit is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *Unit) String() string {
	var builder strings.Builder
	builder.WriteString("Unit(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", content=")
	builder.WriteString(u.Content)
	builder.WriteString(", content_name=")
	builder.WriteString(u.ContentName)
	builder.WriteByte(')')
	return builder.String()
}

// Units is a parsable slice of Unit.
type Units []*Unit

func (u Units) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
