package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// // Fields of the User.
// func (User) Fields() []ent.Field {
// 	return []ent.Field{
// 		field.String("email").
// 			NotEmpty().
// 			Unique(),
// 		field.String("google_id").
// 			Optional().
// 			Unique(),
// 		field.String("password").
// 			Optional(),
// 		field.String("token").
// 			Optional(),
// 		field.String("full_name").
// 			Optional(),
// 		field.Time("created_at").
// 			Default(time.Now).
// 			Immutable(),
// 		field.Time("updated_at").
// 			Default(time.Now).
// 			UpdateDefault(time.Now),
// 	}
// }

// // Edges of the User.
// func (User) Edges() []ent.Edge {
// 	return nil
// }

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").
			NotEmpty().
			Unique(),
		field.String("google_id").
			Optional().
			Unique(),
		field.String("password").
			Optional(),
		field.String("token").
			Optional(),
		field.String("full_name").
			Optional(),
		field.Bool("is_admin").
			Default(false), // Add this field
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (User) Edges() []ent.Edge {
	return nil
}
