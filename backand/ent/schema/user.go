package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique(),
		field.Time("age").
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}).
			Default(nil),
		field.String("email").
			Default("unknown"),
		field.String("googlenum").
			Default(""),
		field.String("hobby").
			Default("unknown"),
		field.String("lang").
			Default("unknown"),
		field.String("github").
			Default("unknown"),
		field.String("gitlab").
			Default("unknown"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("writer", Book.Type),
	}
}
