package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Unit holds the schema definition for the Unit entity.
type Unit struct {
	ent.Schema
}

// Fields of the Unit.
func (Unit) Fields() []ent.Field {
	return []ent.Field{
		field.String("content").
			Default("unknown"),
		field.String("content_name").
			Default("unknown"),
	}
}

// Edges of the Unit.
func (Unit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("contents", Book.Type),
	}
}
