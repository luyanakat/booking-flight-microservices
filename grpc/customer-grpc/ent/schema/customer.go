package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"strconv"
	"time"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
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
		field.String("name").MaxLen(100),
		field.String("email").MaxLen(100),
		field.String("address").MaxLen(200),
		field.String("phone_number").MaxLen(20).MinLen(10),
		field.String("identify_number").MaxLen(12),
		field.Time("date_of_birth"),
		field.String("member_code").Optional(),
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return nil
}
