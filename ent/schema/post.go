package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Post struct {
	ent.Schema
}

func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").NotEmpty(),
		field.String("image").Optional(),
		field.String("category").NotEmpty(),
		field.String("code").Optional(),
		field.String("user_id").NotEmpty(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("likes", Like.Type),
		edge.To("comments", Comment.Type),
		edge.To("shares", Share.Type),
	}
}
