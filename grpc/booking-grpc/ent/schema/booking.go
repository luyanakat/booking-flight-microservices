package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"strconv"
	"time"
)

// Booking holds the schema definition for the Booking entity.
type Booking struct {
	ent.Schema
}

// Fields of the Booking.
func (Booking) Fields() []ent.Field {
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
		field.Int64("customer_id"),
		field.Int64("flight_id"),
		field.String("code").MaxLen(100).DefaultFunc(func() string {
			id, err := gonanoid.Generate("ABCDEFGHIJKMNOPQRSTUVWXYZ123456789", 6)
			if err != nil {
				panic(err)
			}
			return id
		}),
		field.Enum("status").Values("Scheduled", "Cancel", "Completed"),
		field.Int64("ticket_id"),
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Booking.
func (Booking) Edges() []ent.Edge {
	return nil
}
