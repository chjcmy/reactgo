package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Cipher holds the schema definition for the Cipher entity.
type Cipher struct {
	ent.Schema
}

// Fields of the Cipher.
func (Cipher) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("default"),
	}
}

// Edges of the Cipher.
func (Cipher) Edges() []ent.Edge {
	return nil
}
