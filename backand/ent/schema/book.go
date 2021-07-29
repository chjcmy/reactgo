package schema

import "C"
import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			Default("unknown"),
		field.Time("create_at").
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}).
			Default(time.Now),
		field.Time("updated_at").
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}).
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Text("subject").
			Default("unknown"),
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("unitid", Unit.Type).
			Ref("contents").
			Unique(),
		edge.From("userid", User.Type).
			Ref("writer").
			Unique(),
	}
}
