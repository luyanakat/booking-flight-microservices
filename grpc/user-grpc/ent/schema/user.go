package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"strconv"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
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
		field.String("email").MaxLen(200).Unique(),
		field.String("password").MaxLen(255).Sensitive(),
		field.Int64("customer_id").Optional(),
		field.Int64("access_id").Default(2),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
