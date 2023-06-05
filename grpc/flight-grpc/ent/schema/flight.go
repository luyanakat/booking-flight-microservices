package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"strconv"
	"time"
)

// Flight holds the schema definition for the Flight entity.
type Flight struct {
	ent.Schema
}

// Fields of the Flight.
func (Flight) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").DefaultFunc(func() int64 {
			id, err := gonanoid.Generate("123456789", 9)
			if err != nil {
				panic(err)
			}
			ids, err := strconv.Atoi(id)
			if err != nil {
				panic(err)
			}
			return int64(ids)
		}),
		field.String("name").Unique(),
		field.String("from"),
		field.String("to"),
		field.Time("departure_date"),
		field.Time("arrival_date"),
		field.Int("available_first_slot").Default(30).Min(0),
		field.Int("available_economy_slot").Default(70).Min(0),
		field.Enum("status").Values("Available", "Arrived", "Cancel"),
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Flight.
func (Flight) Edges() []ent.Edge {
	return nil
}
