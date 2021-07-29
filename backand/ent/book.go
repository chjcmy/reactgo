// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/backand/ent/book"
	"github.com/backand/ent/unit"
	"github.com/backand/ent/user"
)

// Book is the model entity for the Book schema.
type Book struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt time.Time `json:"create_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Subject holds the value of the "subject" field.
	Subject string `json:"subject,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BookQuery when eager-loading is set.
	Edges         BookEdges `json:"edges"`
	unit_contents *int
	user_writer   *int
}

// BookEdges holds the relations/edges for other nodes in the graph.
type BookEdges struct {
	// Unitid holds the value of the unitid edge.
	Unitid *Unit `json:"unitid,omitempty"`
	// Userid holds the value of the userid edge.
	Userid *User `json:"userid,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UnitidOrErr returns the Unitid value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookEdges) UnitidOrErr() (*Unit, error) {
	if e.loadedTypes[0] {
		if e.Unitid == nil {
			// The edge unitid was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: unit.Label}
		}
		return e.Unitid, nil
	}
	return nil, &NotLoadedError{edge: "unitid"}
}

// UseridOrErr returns the Userid value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookEdges) UseridOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.Userid == nil {
			// The edge userid was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Userid, nil
	}
	return nil, &NotLoadedError{edge: "userid"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Book) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case book.FieldID:
			values[i] = new(sql.NullInt64)
		case book.FieldTitle, book.FieldSubject:
			values[i] = new(sql.NullString)
		case book.FieldCreateAt, book.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case book.ForeignKeys[0]: // unit_contents
			values[i] = new(sql.NullInt64)
		case book.ForeignKeys[1]: // user_writer
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Book", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Book fields.
func (b *Book) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case book.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		case book.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				b.Title = value.String
			}
		case book.FieldCreateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				b.CreateAt = value.Time
			}
		case book.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				b.UpdatedAt = value.Time
			}
		case book.FieldSubject:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field subject", values[i])
			} else if value.Valid {
				b.Subject = value.String
			}
		case book.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field unit_contents", value)
			} else if value.Valid {
				b.unit_contents = new(int)
				*b.unit_contents = int(value.Int64)
			}
		case book.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_writer", value)
			} else if value.Valid {
				b.user_writer = new(int)
				*b.user_writer = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUnitid queries the "unitid" edge of the Book entity.
func (b *Book) QueryUnitid() *UnitQuery {
	return (&BookClient{config: b.config}).QueryUnitid(b)
}

// QueryUserid queries the "userid" edge of the Book entity.
func (b *Book) QueryUserid() *UserQuery {
	return (&BookClient{config: b.config}).QueryUserid(b)
}

// Update returns a builder for updating this Book.
// Note that you need to call Book.Unwrap() before calling this method if this Book
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Book) Update() *BookUpdateOne {
	return (&BookClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the Book entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Book) Unwrap() *Book {
	tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Book is not a transactional entity")
	}
	b.config.driver = tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Book) String() string {
	var builder strings.Builder
	builder.WriteString("Book(")
	builder.WriteString(fmt.Sprintf("id=%v", b.ID))
	builder.WriteString(", title=")
	builder.WriteString(b.Title)
	builder.WriteString(", create_at=")
	builder.WriteString(b.CreateAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(b.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", subject=")
	builder.WriteString(b.Subject)
	builder.WriteByte(')')
	return builder.String()
}

// Books is a parsable slice of Book.
type Books []*Book

func (b Books) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}