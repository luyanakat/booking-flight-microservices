package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// AccessLevel holds the schema definition for the AccessLevel entity.
type AccessLevel struct {
	ent.Schema
}

// Fields of the AccessLevel.
func (AccessLevel) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name").MaxLen(20),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the AccessLevel.
func (AccessLevel) Edges() []ent.Edge {
	return nil
}
